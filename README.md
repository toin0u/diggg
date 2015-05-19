diggg
=====

Proceed a pool of goroutines which execute a `dig` command to a domain against
a list of given DNS servers addresses.

This is an experimental [go](http://golang.org/doc/install) project and it's an
wip as unit tests are missing.


Installation
------------

    $ go get github.com/toin0u/diggg
    $ go install github.com/toin0u/diggg


Configuration
-------------

Customize and copy the `.diggg.dist` file in `~/` without the `.dist` extension.

Pick some DNS servers [her](http://www.chaz6.com/files/resolv.conf) :)


Usage
-----

    $ diggg google.dk


Troubleshooting
---------------

- [ ] ensure to get an output for each `diggg` request (go is too fast :)
- [ ] make nameserver's port customizable (it's not always 53)


Output
------

~~~
diggg, version 0.2.0
Domain to diggg `sbin.dk`

| Provider name             | Provider IP     | sec.  | min.  | hour. | Server IP       |
| CensurFriDNS              | 89.233.43.71    | 780   | 13    | 0     | 185.14.186.176  |
| CensurFriDNS              | 91.239.100.100  | 780   | 13    | 0     | 185.14.186.176  |
| OpenDNS                   | 208.67.220.220  | 779   | 12    | 0     | 185.14.186.176  |
| OpenDNS                   | 208.67.222.222  | 780   | 13    | 0     | 185.14.186.176  |
| DNSadvantage              | 156.154.71.1    | 780   | 13    | 0     | 185.14.186.176  |
| DNSadvantage              | 156.154.70.1    | 780   | 13    | 0     | 185.14.186.176  |
| Norton ConnectSafe DNS    | 198.153.194.60  | 780   | 13    | 0     | 185.14.186.176  |
| Norton ConnectSafe DNS    | 198.153.192.50  | 780   | 13    | 0     | 185.14.186.176  |
| Norton ConnectSafe DNS    | 198.153.194.40  | 780   | 13    | 0     | 185.14.186.176  |
| Norton ConnectSafe DNS    | 198.153.192.60  | 780   | 13    | 0     | 185.14.186.176  |
| Norton ConnectSafe DNS    | 198.153.192.40  | 780   | 13    | 0     | 185.14.186.176  |
| Norton ConnectSafe DNS    | 198.153.194.50  | 780   | 13    | 0     | 185.14.186.176  |
| Comodo Secure DNS         | 8.26.56.26      | 780   | 13    | 0     | 185.14.186.176  |
| Comodo Secure DNS         | 8.20.247.20     | 2871  | 47    | 0     | 185.14.186.176  |
| Google                    | 8.8.4.4         | 1846  | 30    | 0     | 185.14.186.176  |
| Level3                    | 4.2.2.2         | 2858  | 47    | 0     | 185.14.186.176  |
| Google                    | 8.8.8.8         | 2834  | 47    | 0     | 185.14.186.176  |
| Level3                    | 4.2.2.1         | 3600  | 60    | 1     | 185.14.186.176  |
| SmartViper                | 208.76.50.50    | 780   | 13    | 0     | 185.14.186.176  |
| SmartViper                | 208.76.51.51    | 780   | 13    | 0     | 185.14.186.176  |
| PowerTech                 | 195.159.0.100   | read udp 195.159.0.100:53: i/o timeout  |
~~~


License
-------

diggg is released under the MIT License. See the bundled
[LICENSE](https://github.com/toin0u/diggg/blob/master/LICENSE) file for details.
