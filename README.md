# pingme

# Intro

Pingme is lightweight ping probe tool, supporting ICMP and TCP protocols.

# Features

- Support ICMP/TCP protocols
- Display basic IP information
- Monitor real-time ping statistics

# Usage

```
  -c string
        Config file
  -d    Daemon mode
  -i string
        ICMP destination
  -m string
        MTR destination
  -q string
        Query address
  -s    Serve mode
  -t string
        TCP destination
  -v    Version
```

# Examples

```
// If only address is provided, pingme will probe some common ports.
$ pingme google.com
INFO     ICMP    OPEN      2404:6800:4003:c01::8a    2.4 ms
WARN     TCP     ERROR     [2404:6800:4003:c01::8a]:22
INFO     TCP     OPEN      [2404:6800:4003:c01::8a]:80
INFO     TCP     OPEN      [2404:6800:4003:c01::8a]:443

// Url can be passed to pingme in query mode.
// This feature is designed for convenient copy & paste.
$ pingme -q https://www.google.com/
IP     :    2404:6800:4003:c03::93
City   :    Singapore
Country:    Singapore
ISP    :    Google LLC
AS     :    AS15169 Google LLC

// You can specify any port in tcping mode.
$ pingme -t google.com:12345
WARN     TCP     ERROR     google.com:12345

```

# Note

Root permission is required when running ICMP ping, since it needs open raw socket.

You can either use sudo command, or set setuid bit for pingme.

```
// Use sudo for one-time ping
$ sudo pingme -i google.com

// Set setuid bit
$ sudo chown root:root pingme
$ sudo chmod u+s pingme

```

# License
See the LICENSE file for license rights and limitations (MIT).

# Acknowledgements
[httpstat](https://github.com/davecheney/httpstat)
