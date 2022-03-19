<!--
  ~ Copyright 2022 kwanhur
  ~
  ~ Licensed under the Apache License, Version 2.0 (the "License");
  ~ you may not use this file except in compliance with the License.
  ~ You may obtain a copy of the License at
  ~
  ~ http://www.apache.org/licenses/LICENSE-2.0
  ~
  ~ Unless required by applicable law or agreed to in writing, software
  ~ distributed under the License is distributed on an "AS IS" BASIS,
  ~ WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  ~ See the License for the specific language governing permissions and
  ~ limitations under the License.
  ~
-->

## Name

A modern Linux Virtual Server controller.

## Badges

[![Release](https://img.shields.io/github/release/kwanhur/ipvsctl.svg?style=for-the-badge)](https://github.com/kwanhur/ipvsctl/releases/latest)
[![Software License](https://img.shields.io/badge/license-Apache2.0-brightgreen.svg?style=for-the-badge)](LICENSE)
[![Build status](https://img.shields.io/github/workflow/status/kwanhur/ipvsctl/build?style=for-the-badge)](https://github.com/kwanhur/ipvsctl/actions?workflow=build)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=for-the-badge)](http://godoc.org/github.com/kwanhur/ipvsctl)
[![Powered By: KWANHUR](https://img.shields.io/badge/powered%20by-kwanhur-green.svg?style=for-the-badge)](https://github.com/kwanhur)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg?style=for-the-badge)](https://conventionalcommits.org)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit)

## Description

`ipvsctl` is similar to `ipvsadm`, `ipvsctl` support multiple modern operations, include sub-commands `service`, `server`, `timeout`, `zero`.

## Commands

### Service

`ipvsctl` can be used to set up, maintain or retrieve the virtual server table in the Linux kernel.
Supported sub-commands `list`, `string`, `check`, `get`, `add`, `update`, `delete`, `zero`, `flush`.

```shell
./ipvsctl service -h
NAME:
   ipvsctl service - Operates virtual service[vip:vport protocol] (TCP UDP STCP)/(IPv4 IPv6)

USAGE:
   ipvsctl service command [command options] [arguments...]

COMMANDS:
   list, l, ls                 List ipvs virtual service
   string, s, str              Present ipvs virtual service string
   check, c, chk, exist, find  Check ipvs virtual service exist or not
   get, g, fetch, one          Get ipvs virtual service
   add, a, new, n, set         Add ipvs virtual service
   update, u, up               Update ipvs virtual service
   del, d, delete              Del ipvs virtual service
   zero, z                     Zero ipvs virtual server stats(byte packet and rate counters)
   flush, f, clear             Flush ipvs, all the rules will be clear
   help, h                     Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help (default: false)
```

#### List Service

Retrieve all the virtual service. Supported statistics information with flag `--stats` or `-s`.
Statistics includes `Conn`, `PktsIn`, `PktsOut`, `BytesIn`, `BytesOut`, `CPS`, `BPSIn`, `BPSOut`, `PPSIn`, `PPSOut`.

```shell
./ipvsctl service list -h
NAME:
   ipvsctl service list - List ipvs virtual service

USAGE:
   ipvsctl service list [command options] [arguments...]

OPTIONS:
   --stats, -s, --stat  Show statistics information (default: false)
   --help, -h           show help (default: false)
```

#### Check Service

Check specified service existed in Linux virtual server table or not.

```shell
./ipvsctl service check -h
NAME:
   ipvsctl service check - Check ipvs virtual service exist or not

USAGE:
   ipvsctl service check [command options] [arguments...]

OPTIONS:
   --vip value                      Specify vs IP address
   --vport value                    Specify vs port number, range [0-65535] (default: 0)
   --protocol value, --proto value  Specify vs protocol, option [TCP UDP SCTP] (default: TCP)
   --help, -h                       show help (default: false)
```

#### Get Service

Retrieve one virtual service. Similar to [List Service](#List-Service)

```shell
./ipvsctl service get -h
NAME:
   ipvsctl service get - Get ipvs virtual service

USAGE:
   ipvsctl service get [command options] [arguments...]

OPTIONS:
   --vip value                      Specify vs IP address
   --vport value                    Specify vs port number, range [0-65535] (default: 0)
   --protocol value, --proto value  Specify vs protocol, option [TCP UDP SCTP] (default: TCP)
   --stats, -s, --stat              Show statistics information (default: false)
   --help, -h                       show help (default: false)
```

#### Add Service

Add one virtual service into Linux virtual server table. Supported two address family (IPv4, IPv6), three protocols (TCP, UDP, STCP), and ten load balancing algorithms (round-robin, weighted round-robin, least-connection,  weighted  least-connection,
locality-based least-connection, locality-based least-connection with replication, destination-hashing, source-hashing, shortest expected delay and never queue).

```shell
./ipvsctl service add -h
NAME:
   ipvsctl service add - Add ipvs virtual service

USAGE:
   ipvsctl service add [command options] [arguments...]

DESCRIPTION:
   Add a virtual service. A service address is uniquely defined by a triplet: IP address, port number,  and  protocol.  Alternatively,  a virtual service may be defined by a firewall-mark.

OPTIONS:
   --vip value                                                  Specify vs IP address
   --vport value                                                Specify vs port number, range [0-65535] (default: 0)
   --protocol value, --proto value                              Specify vs protocol, option [TCP UDP SCTP] (default: TCP)
   --scheduler value, -s value, --sched value                   Specify vs scheduling method, option [rr wrr lc wlc lblc lblcr dh sh sed nq]
   --persistent value, --pe value, --per value, --persis value  Specify vs persistent name
   --timeout value                                              Specify persistent session timeout in seconds (default: 300)
   --netmask value, -M value, --mask value                      Specify which clients are grouped for persistent virtual service, default IPv4/32 IPv6/128 (default: 0)
   --help, -h                                                   show help (default: false)
```

#### Update Service

Update specified virtual service. Attributes include `scheduler`, `persistent`, `timeout`, `netmask` could be updated.

```shell
./ipvsctl service update -h
NAME:
   ipvsctl service update - Update ipvs virtual service

USAGE:
   ipvsctl service update [command options] [arguments...]

OPTIONS:
   --vip value                                                  Specify vs IP address
   --vport value                                                Specify vs port number, range [0-65535] (default: 0)
   --protocol value, --proto value                              Specify vs protocol, option [TCP UDP SCTP] (default: TCP)
   --scheduler value, -s value, --sched value                   Specify vs scheduling method, option [rr wrr lc wlc lblc lblcr dh sh sed nq]
   --persistent value, --pe value, --per value, --persis value  Specify vs persistent name
   --timeout value                                              Specify persistent session timeout in seconds (default: 300)
   --netmask value, -M value, --mask value                      Specify which clients are grouped for persistent virtual service, default IPv4/32 IPv6/128 (default: 0)
   --help, -h                                                   show help (default: false)
```

#### Delete Service

Delete specified virtual service.

```shell
./ipvsctl service delete -h
NAME:
   ipvsctl service del - Del ipvs virtual service

USAGE:
   ipvsctl service del [command options] [arguments...]

OPTIONS:
   --vip value                      Specify vs IP address
   --vport value                    Specify vs port number, range [0-65535] (default: 0)
   --protocol value, --proto value  Specify vs protocol, option [TCP UDP SCTP] (default: TCP)
   --help, -h                       show help (default: false)
```

#### Zero Service

Zero specified virtual service statistics.

```shell
./ipvsctl service zero -h
NAME:
   ipvsctl service zero - Zero ipvs virtual server stats(byte packet and rate counters)

USAGE:
   ipvsctl service zero [command options] [arguments...]

OPTIONS:
   --vip value                      Specify vs IP address
   --vport value                    Specify vs port number, range [0-65535] (default: 0)
   --protocol value, --proto value  Specify vs protocol, option [TCP UDP SCTP] (default: TCP)
   --help, -h                       show help (default: false)
```

#### Flush Service

Flush all the virtual services.

```shell
./ipvsctl service flush -h
NAME:
   ipvsctl service flush - Flush ipvs, all the rules will be clear

USAGE:
   ipvsctl service flush [command options] [arguments...]

OPTIONS:
   --help, -h  show help (default: false)
```

### Server

`ipvsctl` can be used to set up, maintain or retrieve the real server table with the specified virtual server.

The Linux Virtual Server can be used to build scalable network services based on a cluster of two or more nodes. The active node of the cluster redirects service requests to a collection of  server
hosts  that  will actually perform the services.

Supported sub-commands `list`, `add`, `delete`, `flush`.

```shell
./ipvsctl server -h
NAME:
   ipvsctl server - Operates real server[rip:rport] (IPv4/IPv6)

USAGE:
   ipvsctl server command [command options] [arguments...]

COMMANDS:
   list, l, ls             List ipvs real server
   add, a, new, n, set     Add ipvs real server
   del, d, delete          Del ipvs real server
   flush, f, clear, purge  Flush rs, all the real servers will be clear
   help, h                 Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help (default: false)
```

#### List Server

Retrieve all the servers with specified virtual server. Supported statistics information with flag `--stats` or `-s`.
Statistics includes `ActConn`, `InactConn`, `PersConn`, `Conn`, `PktsIn`, `PktsOut`, `BytesIn`, `BytesOut`, `CPS`, `BPSIn`, `BPSOut`, `PPSIn`, `PPSOut`.

```shell
./ipvsctl server list -h
NAME:
   ipvsctl server list - List ipvs real server

USAGE:
   ipvsctl server list [command options] [arguments...]

OPTIONS:
   --vip value                      Specify vs IP address
   --vport value                    Specify vs port number, range [0-65535] (default: 0)
   --protocol value, --proto value  Specify vs protocol, option [TCP UDP SCTP] (default: TCP)
   --stats, -s, --stat              Show statistics information (default: false)
   --help, -h                       show help (default: false)
```

#### Add Server

Add one real server into specified virtual service. Supported  two address family (IPv4, IPv6), and seven forward modes (local, mask, masq, bypass, direct route, tunnel, full-nat).

```shell
./ipvsctl server add -h
NAME:
   ipvsctl server add - Add ipvs real server

USAGE:
   ipvsctl server add [command options] [arguments...]

OPTIONS:
   --vip value                      Specify vs IP address
   --vport value                    Specify vs port number, range [0-65535] (default: 0)
   --protocol value, --proto value  Specify vs protocol, option [TCP UDP SCTP] (default: TCP)
   --rip value                      Specify rs IP address
   --rport value                    Specify rs port number, range [0-65535] (default: 0)
   --forward value, --fwd value     Specify rs connection flag, option [local mask masq bypass dr tun fnat] (default: dr)
   --weight value, -w value         Specify rs weight (default: 0)
   --help, -h                       show help (default: false)
```

#### Update Server

Update real server attributes `forward`, `weight` with specified virtual service.

```shell
./ipvsctl server update -h
NAME:
   ipvsctl server update - Update ipvs real server

USAGE:
   ipvsctl server update [command options] [arguments...]

OPTIONS:
   --vip value                      Specify vs IP address
   --vport value                    Specify vs port number, range [0-65535] (default: 0)
   --protocol value, --proto value  Specify vs protocol, option [TCP UDP SCTP] (default: TCP)
   --rip value                      Specify rs IP address
   --rport value                    Specify rs port number, range [0-65535] (default: 0)
   --forward value, --fwd value     Specify rs connection flag, option [local mask masq bypass dr tun fnat] (default: dr)
   --weight value, -w value         Specify rs weight (default: 0)
   --help, -h                       show help (default: false)
```

#### Delete Server

Delete one real server with specified virtual service.

```shell
./ipvsctl server delete -h
NAME:
   ipvsctl server del - Del ipvs real server

USAGE:
   ipvsctl server del [command options] [arguments...]

OPTIONS:
   --vip value                      Specify vs IP address
   --vport value                    Specify vs port number, range [0-65535] (default: 0)
   --protocol value, --proto value  Specify vs protocol, option [TCP UDP SCTP] (default: TCP)
   --rip value                      Specify rs IP address
   --rport value                    Specify rs port number, range [0-65535] (default: 0)
   --help, -h                       show help (default: false)
```

#### Flush Server

Flush all the real servers with specified virtual service.

```shell
./ipvsctl server flush -h
NAME:
   ipvsctl server flush - Flush rs, all the real servers will be clear

USAGE:
   ipvsctl server flush [command options] [arguments...]

OPTIONS:
   --vip value                      Specify vs IP address
   --vport value                    Specify vs port number, range [0-65535] (default: 0)
   --protocol value, --proto value  Specify vs protocol, option [TCP UDP SCTP] (default: TCP)
   --help, -h                       show help (default: false)
```

### Timeout

It's about virtual service connections' timeout values (in seconds) for TCP sessions, TCP sessions after receiving a FIN packet, and UDP packets.

Supported sub-commands `show`, `set`.

```shell
./ipvsctl timeout -h
NAME:
   ipvsctl timeout - Operates timeout (tcp tcpfin udp)

USAGE:
   ipvsctl timeout command [command options] [arguments...]

COMMANDS:
   show, ls, get  Shows timeout of tcp tcpfin udp
   set, s         Sets timeout of tcp tcpfin udp
   help, h        Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help (default: false)
```

#### Show Timeout

Show the current timeout values.

```shell
./ipvsctl timeout show -h
NAME:
   ipvsctl timeout show - Shows timeout of tcp tcpfin udp

USAGE:
   ipvsctl timeout show [command options] [arguments...]

OPTIONS:
   --help, -h  show help (default: false)
```

#### Set Timeout

Change timeout values.

```shell
./ipvsctl timeout set -h
NAME:
   ipvsctl timeout set - Sets timeout of tcp tcpfin udp

USAGE:
   ipvsctl timeout set [command options] [arguments...]

DESCRIPTION:
   Change the timeout values used for IPVS connections. This command support 3 options, representing  the  timeout   values  (in seconds)  for TCP sessions, TCP sessions after receiving a  FIN packet, and  UDP  packets, respectively. A timeout value 0 means that the current timeout value of the  corresponding  entry  is preserved.

OPTIONS:
   --tcp value, -t value                  set tcp timeout(unit second), 0 means no change (default: 900)
   --tcpfin value, -f value, --fin value  set tcpfin timeout(unit second), 0 means no change (default: 120)
   --udp value, -u value                  set udp timeout(unit second), 0 means no change (default: 300)
   --help, -h                             show help (default: false)
```

### Zero

It's used to clear out all the virtual services' statistics.

```shell
./ipvsctl zero -h
NAME:
   ipvsctl zero - Zero ipvs all the virtual service stats(byte packet and rate counters)

USAGE:
   ipvsctl zero [command options] [arguments...]

OPTIONS:
   --yes, --force, -f, -y  Are you agree to do it?[yes/no] (default: false)
   --help, -h              show help (default: false)
```
