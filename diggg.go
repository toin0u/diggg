package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type DnsType []struct {
	Name string   `json:"name"`
	Ipv4 []string `json:"ipv4"`
}

type Servers struct {
	Dns DnsType `json:"dns"`
}

const appVersion = "0.1.0"
const appName = "diggg"
const listDnsFile = ".diggg"
const outputNormalFormat = "| %-25s | %-15s | %-5s | %-5v | %-5v | %-15s |\n"
const outputErrorFormat = "| %-25s | %-15s | %-39s |\n"

var ttlRegexp = regexp.MustCompile(`\d+`)
var ipv4Regexp = regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}\b`)

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	log.Printf("%s took %s to execute.", appName, elapsed)
}

func diggg(domain string, dns string, ip string, wg *sync.WaitGroup) {
	wg.Add(1)

	command := fmt.Sprintf("dig @%s +noall +answer %s", ip, domain)

	parts := strings.Fields(command)
	head := parts[0]
	parts = parts[1:len(parts)]

	output, err := exec.Command(head, parts...).Output()

	if err != nil {
		fmt.Printf(outputErrorFormat, dns, ip, err)
		wg.Done()
		return
	}

	if len(output) == 0 {
		fmt.Printf(outputErrorFormat, dns, ip, "timeout (consider to remove this ip)")
		wg.Done()
		return
	}

	ttl := ttlRegexp.FindString(string(output))
	ttl2, _ := strconv.ParseFloat(ttl, 64)
	ipv4 := ipv4Regexp.FindString(string(output))

	fmt.Printf(outputNormalFormat, dns, ip, ttl, math.Floor(ttl2/60), math.Floor(ttl2/3600), ipv4)

	wg.Done()
}

func main() {
	fmt.Printf("%s, version %s\n", appName, appVersion)

	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s <domain url>\n", appName)
		return
	}

    filePath := filepath.Join(os.Getenv("HOME"), listDnsFile)
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("Error the config file `%s` is missing.\n", filePath)
		return
	}

	fmt.Printf("Domain to diggg `%s`\n\n", os.Args[1])

	decoder := json.NewDecoder(file)
	servers := &Servers{}
	decoder.Decode(&servers)

	fmt.Printf(outputNormalFormat, "Provider name", "Provider IP", "sec.", "min.", "hour.", "Server IP")
	defer timeTrack(time.Now())

	wg := new(sync.WaitGroup)

	for _, dns := range servers.Dns {
		for _, ipv4 := range dns.Ipv4 {
			go diggg(os.Args[1], dns.Name, ipv4, wg)
		}
	}

	wg.Wait()
}
