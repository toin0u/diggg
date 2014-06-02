diggg
=====

Proceed a pool of goroutines which execute a `dig` command to a domain against
a list of given DNS servers addresses.

This is an experimental [go](http://golang.org/doc/install) project and it's an
wip as unit tests are missing.


Requirement
-----------

The `dig` command should be present and findable in `$PATH`.


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


Output
------

~~~
diggg, version 0.1.0
Domain to diggg `ucc.dk`

| Provider name             | Provider IP     | sec.  | min.  | hour. | Server IP       |
| TDC                       | 194.239.134.83  | 8698  | 144   | 2     | 46.137.187.34   |
| Google                    | 8.8.8.8         | 3766  | 62    | 1     | 46.137.187.34   |
| Google                    | 8.8.4.4         | 3766  | 62    | 1     | 46.137.187.34   |
| CensurFriDNS              | 89.233.43.71    | 76010 | 1266  | 21    | 46.137.187.34   |
| OpenDNS                   | 208.67.220.220  | 72467 | 1207  | 20    | 46.137.187.34   |
| Telenor                   | 212.242.40.51   | timeout (consider to remove this ip)    |
| Level3                    | 4.2.2.1         | 86400 | 1440  | 24    | 46.137.187.34   |
| NextGenTel                | 217.13.4.21     | timeout (consider to remove this ip)    |
| SmartViper                | 208.76.50.50    | 47342 | 789   | 13    | 46.137.187.34   |
| DNS Reactor               | 204.45.18.18    | 47335 | 788   | 13    | 46.137.187.34   |
| Norton ConnectSafe DNS    | 198.153.192.50  | 47337 | 788   | 13    | 46.137.187.34   |
| Comodo Secure DNS         | 8.26.56.26      | 47337 | 788   | 13    | 46.137.187.34   |
| Telenor                   | 193.213.112.4   | timeout (consider to remove this ip)    |
| Telia                     | 194.255.56.79   | timeout (consider to remove this ip)    |
| Dataguard                 | 213.158.233.130 | timeout (consider to remove this ip)    |
| PowerTech                 | 195.159.0.100   | 47341 | 789   | 13    | 46.137.187.34   |
| OpenDNS                   | 208.67.222.222  | 268   | 4     | 0     | 46.137.187.34   |
| DNSadvantage              | 156.154.70.1    | 76004 | 1266  | 21    | 46.137.187.34   |
| Norton ConnectSafe DNS    | 198.153.192.40  | 82469 | 1374  | 22    | 46.137.187.34   |
| TDC                       | 193.162.153.164 | 40326 | 672   | 11    | 46.137.187.34   |
| Norton ConnectSafe DNS    | 198.153.192.60  | 47332 | 788   | 13    | 46.137.187.34   |
| Comodo Secure DNS         | 8.20.247.20     | 47332 | 788   | 13    | 46.137.187.34   |
| CensurFriDNS              | 89.104.194.142  | 36101 | 601   | 10    | 46.137.187.34   |
| Telenor                   | 212.242.40.3    | timeout (consider to remove this ip)    |
| Dataguard                 | 213.158.233.131 | timeout (consider to remove this ip)    |
| Level3                    | 4.2.2.2         | 86311 | 1438  | 23    | 46.137.187.34   |
| Telenor                   | 148.122.161.3   | timeout (consider to remove this ip)    |
| Stofanet                  | 212.10.10.4     | timeout (consider to remove this ip)    |
| OpenNIC                   | 78.46.76.146    | timeout (consider to remove this ip)    |
| DNSadvantage              | 156.154.71.1    | 76005 | 1266  | 21    | 46.137.187.34   |
| Norton ConnectSafe DNS    | 198.153.194.40  | 76005 | 1266  | 21    | 46.137.187.34   |
| Telia                     | 194.255.56.78   | timeout (consider to remove this ip)    |
| NextGenTel                | 217.13.7.136    | timeout (consider to remove this ip)    |
| Stofanet                  | 212.10.10.5     | timeout (consider to remove this ip)    |
| Norton ConnectSafe DNS    | 198.153.194.60  | 76005 | 1266  | 21    | 46.137.187.34   |
| Norton ConnectSafe DNS    | 198.153.194.50  | 76005 | 1266  | 21    | 46.137.187.34   |
| Telenor                   | 212.88.64.14    | exit status 9                           |
| Telenor                   | 212.88.64.199   | exit status 9                           |
| Telenor                   | 212.88.64.15    | exit status 9                           |
| PowerTech                 | 195.159.0.200   | exit status 9                           |
| SmartViper                | 208.76.51.51    | exit status 9                           |
| DNS Reactor               | 204.45.18.26    | exit status 9                           |
~~~


License
-------

diggg is released under the MIT License. See the bundled
[LICENSE](https://github.com/toin0u/diggg/blob/master/LICENSE) file for details.
