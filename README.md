# Introduction

Pingme is lightweight ping probe command line tool, supporting ICMP, TCP and HTTP protocols.

It can also query IP information from third-party api provider (currently we use [https://ip-api.com](https://ip-api.com)).

# Features

- Support ICMP/TCP/HTTP protocols
- Query basic IP information

# Installation

1. Download latest [release](https://github.com/noobly314/pingme/releases/latest) (recommend)

2. Use go get

```
go get github.com/noobly314/pingme
```

3. Build on your own

```
git clone https://github.com/noobly314/pingme.git
cd pingme
go build
```

# Usage

```
  -h string
        HTTP Ping
  -i string
        ICMP Ping
  -m string
        MTR Trace
  -q string
        Query ip information
  -t string
        TCP Ping
  -v    Version
```

# Examples

```
// Pingme will query ip information and do httping by default.
$ pingme https://www.google.com
IP     :    74.125.24.105
City   :    Ashburn
Country:    United States
ISP    :    Google LLC
AS     :    AS15169 Google LLC

Proxy     :    false
Scheme    :    https
Host      :    www.google.com
DNS Lookup:    0.86 ms
TCP       :    55.87 ms
TLS       :    54.13 ms
Process   :    31.72 ms
Transfer  :    0.26 ms
Total     :    88.71 ms
```

```
// You can specify any ports in tcping mode.
$ pingme -t google.com:12345
TCP     ERROR     google.com:12345
```

# Note

Root permission is required when running ICMP ping, since it needs to open raw socket.

You can either use sudo command, or set setuid bit for pingme.

```
// Use sudo for one-time ping
$ sudo pingme -i google.com

// Set setuid bit
$ sudo chown root:root pingme
$ sudo chmod u+s pingme

```

# License

See the [LICENSE](https://github.com/noobly314/pingme/blob/master/LICENSE.md) file for license rights and limitations (MIT).

# Acknowledgements

[https://ip-api.com](https://ip-api.com)

[lmas/icmp_ping.go](https://gist.github.com/lmas/c13d1c9de3b2224f9c26435eb56e6ef3)

[sparrc/go-ping](https://github.com/sparrc/go-ping)

[davecheney/httpstat](https://github.com/davecheney/httpstat)
