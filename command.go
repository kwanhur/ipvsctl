// Copyright 2022 kwanhur
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package main

import "github.com/urfave/cli/v2"

var flags = []cli.Flag{
	&cli.BoolFlag{
		Name:     "yes",
		Aliases:  []string{"force", "f", "y"},
		Usage:    "Are you agree to do it?[yes/no]",
		Required: true,
	},
}

var statFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "stats",
		Aliases: []string{"s", "stat"},
		Usage:   "Show statistics information",
	},
}

var vsFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     "vip",
		Usage:    "Specify vs IP address",
		Required: true,
	},
	&cli.UintFlag{
		Name:     "vport",
		Usage:    "Specify vs port number, range [0-65535]",
		Required: true,
	},
	&cli.StringFlag{
		Name:        "protocol",
		Aliases:     []string{"proto"},
		Usage:       "Specify vs protocol, option [TCP UDP SCTP]",
		DefaultText: "TCP",
	},
}

var vsFlags2 = []cli.Flag{
	&cli.StringFlag{
		Name:  "vip",
		Usage: "Specify vs IP address",
	},
	&cli.UintFlag{
		Name:  "vport",
		Usage: "Specify vs port number, range [0-65535]",
	},
}

var vsOptionFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    "scheduler",
		Aliases: []string{"s", "sched"},
		Usage:   "Specify vs scheduling method, option [rr wrr lc wlc lblc lblcr dh sh sed nq]",
	},
	&cli.StringFlag{
		Name:    "persistent",
		Aliases: []string{"pe", "per", "persis"},
		Usage:   "Specify vs persistent name",
	},
	&cli.UintFlag{
		Name:        "timeout",
		Usage:       "Specify persistent session timeout in seconds",
		DefaultText: "300",
	},
	&cli.UintFlag{
		Name:    "netmask",
		Aliases: []string{"M", "mask"},
		Usage:   "Specify which clients are grouped for persistent virtual service, default IPv4/32 IPv6/128",
	},
}

var rsFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     "rip",
		Usage:    "Specify rs IP address",
		Required: true,
	},
	&cli.UintFlag{
		Name:     "rport",
		Usage:    "Specify rs port number, range [0-65535]",
		Required: true,
	},
}

var rsFlags2 = append(rsFlags, []cli.Flag{
	&cli.StringFlag{
		Name:        "forward",
		Aliases:     []string{"fwd"},
		Usage:       "Specify rs connection flag, option [local mask masq bypass dr tun fnat]",
		DefaultText: "dr",
		Required:    true,
	},
	&cli.IntFlag{
		Name:        "weight",
		Aliases:     []string{"w"},
		Usage:       "Specify rs weight",
		DefaultText: "0",
		Required:    true,
	},
}...)

// BasicCommands return basic operations
func (o *Operator) BasicCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "zero",
			Aliases: []string{"z"},
			Usage:   "Zero ipvs all the virtual service stats(byte packet and rate counters)",
			Action:  o.Zero(),
			Flags:   flags,
		},
		{
			Name:    "flush",
			Aliases: []string{"f", "clear"},
			Usage:   "Flush out the virtual server table",
			Action:  o.FlushService(),
			Flags:   flags,
		},
	}
}

// ServiceCommands return service relate operations, like get set flush import export
func (o *Operator) ServiceCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "service",
			Aliases: []string{"s", "svc", "vs"},
			Usage:   "Operates virtual service[vip:vport protocol] (TCP UDP STCP)/(IPv4 IPv6)",
			Subcommands: []*cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l", "ls"},
					Usage:   "List ipvs virtual service",
					Action:  o.ListService(),
					Flags:   statFlags,
				},
				{
					Name:    "string",
					Aliases: []string{"s", "str"},
					Usage:   "Present ipvs virtual service string",
					Action:  o.StringService(),
					Flags:   vsFlags,
				},
				{
					Name:    "check",
					Aliases: []string{"c", "chk", "exist", "find"},
					Usage:   "Check ipvs virtual service exist or not",
					Action:  o.ExistService(),
					Flags:   vsFlags,
				},
				{
					Name:    "get",
					Aliases: []string{"g", "fetch", "one"},
					Usage:   "Get ipvs virtual service",
					Action:  o.GetService(),
					Flags:   append(vsFlags, statFlags...),
				},
				{
					Name:        "add",
					Aliases:     []string{"a", "new", "n", "set"},
					Usage:       "Add ipvs virtual service",
					Description: `Add a virtual service. A service address is uniquely defined by a triplet: IP address, port number,  and  protocol.  Alternatively,  a virtual service may be defined by a firewall-mark.`,
					Action:      o.AddService(),
					Flags:       append(vsFlags, vsOptionFlags...),
				},
				{
					Name:    "update",
					Aliases: []string{"u", "up"},
					Usage:   "Update ipvs virtual service",
					Action:  o.UpdateService(),
					Flags:   append(vsFlags, vsOptionFlags...),
				},
				{
					Name:    "del",
					Aliases: []string{"d", "delete"},
					Usage:   "Del ipvs virtual service",
					Action:  o.DelService(),
					Flags:   vsFlags,
				},
				{
					Name:    "zero",
					Aliases: []string{"z"},
					Usage:   "Zero ipvs virtual server stats(byte packet and rate counters)",
					Action:  o.ZeroService(),
					Flags:   vsFlags,
				},
				{
					Name:    "flush",
					Aliases: []string{"f", "clear"},
					Usage:   "Flush out the virtual server table",
					Action:  o.FlushService(),
					Flags:   flags,
				},
			},
		},
	}
}

// ServerCommands  return server relate operations, like add update del flush
func (o *Operator) ServerCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "server",
			Aliases: []string{"ser", "svr", "d", "dst", "dest", "destination", "rs", "realserver"},
			Usage:   "Operates real server[rip:rport] (IPv4/IPv6)",
			Subcommands: []*cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l", "ls"},
					Usage:   "List ipvs real server",
					Action:  o.ListServer(),
					Flags:   append(vsFlags, statFlags...),
				},
				{
					Name:    "add",
					Aliases: []string{"a", "new", "n", "set"},
					Usage:   "Add ipvs real server",
					Action:  o.AddServer(),
					Flags:   append(vsFlags, rsFlags2...),
				},
				{
					Name:    "update",
					Aliases: []string{"u", "up"},
					Usage:   "Update ipvs real server",
					Action:  o.UpdateServer(),
					Flags:   append(vsFlags, rsFlags2...),
				},
				{
					Name:    "del",
					Aliases: []string{"d", "delete"},
					Usage:   "Del ipvs real server",
					Action:  o.DelServer(),
					Flags:   append(vsFlags, rsFlags...),
				},
				{
					Name:    "flush",
					Aliases: []string{"f", "clear", "purge"},
					Usage:   "Flush rs, all the real servers will be clear",
					Action:  o.FlushServer(),
					Flags:   vsFlags,
				},
			},
		},
	}
}

// AddressCommands  return local address relate operations, like add del flush
func (o *Operator) AddressCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "address",
			Aliases: []string{"a", "addr", "local-address", "la", "laddr"},
			Usage:   "Operates local address (IPv4/IPv6)",
			Subcommands: []*cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l", "ls"},
					Usage:   "List ipvs local address",
					Action:  o.ListAddress(),
					Flags:   vsFlags,
				},
				{
					Name:    "add",
					Aliases: []string{"a", "new", "n", "set"},
					Usage:   "Add ipvs local address",
					Action:  o.AddAddress(),
					Flags: append(vsFlags, &cli.StringSliceFlag{
						Name:     "lip",
						Usage:    "Specify vs local address, multiple with separator comma",
						Required: true,
					}),
				},
				{
					Name:    "del",
					Aliases: []string{"d", "delete"},
					Usage:   "Del ipvs local address",
					Action:  o.DelAddress(),
					Flags: append(vsFlags, &cli.StringSliceFlag{
						Name:     "lip",
						Usage:    "Specify vs local address, multiple with separator comma",
						Required: true,
					}),
				},
				{
					Name:    "flush",
					Aliases: []string{"f", "clear", "purge"},
					Usage:   "Flush all the local address",
					Action:  o.FlushAddress(),
					Flags:   vsFlags,
				},
			},
		},
	}
}

// DaemonCommands return daemon relate operations, like show add del
func (o *Operator) DaemonCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "daemon",
			Aliases: []string{"dm"},
			Usage:   "Operates synchronization daemon",
			Subcommands: []*cli.Command{
				{
					Name:    "show",
					Aliases: []string{"ls", "get"},
					Usage:   "Shows daemons of state, sync-id and mcast-interface",
					Action:  o.ShowDaemon(),
				},
				{
					Name:    "add",
					Aliases: []string{"a", "new", "n", "set"},
					Usage:   "Add daemon, currently only supports IPv4 connections",
					Action:  o.AddDaemon(),
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:     "state",
							Usage:    "Specify daemon state, option [master backup]",
							Required: true,
						},
						&cli.UintFlag{
							Name:     "sync-id",
							Usage:    "Specify daemon syncId",
							Required: true,
						},
						&cli.StringFlag{
							Name:     "mcast-ifn",
							Usage:    "Specify daemon mcast-interface",
							Required: true,
						},
					},
				},
				{
					Name:    "del",
					Aliases: []string{"d", "del"},
					Usage:   "Del daemon",
					Action:  o.DelDaemon(),
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:     "state",
							Usage:    "Specify daemon state, option [master backup]",
							Required: true,
						},
					},
				},
			},
		},
	}
}

// TimeoutCommands return timeout relate operations, like get set
func (o *Operator) TimeoutCommands() []*cli.Command {
	var cmds = []*cli.Command{
		{
			Name:    "timeout",
			Aliases: []string{"t", "to", "out"},
			Usage:   "Operates timeout (tcp tcpfin udp)",
			Subcommands: []*cli.Command{
				{
					Name:    "show",
					Aliases: []string{"ls", "get"},
					Usage:   "Shows timeout of tcp tcpfin udp",
					Action:  o.ShowTimeout(),
				},
				{
					Name:        "set",
					Aliases:     []string{"s"},
					Usage:       "Sets timeout of tcp tcpfin udp",
					Description: `Change the timeout values used for IPVS connections. This command support 3 options, representing  the  timeout   values  (in seconds)  for TCP sessions, TCP sessions after receiving a  FIN packet, and  UDP  packets, respectively. A timeout value 0 means that the current timeout value of the  corresponding  entry  is preserved.`,
					Action:      o.SetTimeout(),
					Flags: []cli.Flag{
						&cli.IntFlag{
							Name:        "tcp",
							Aliases:     []string{"t"},
							Usage:       "set tcp timeout(unit second), 0 means no change",
							DefaultText: "900",
						},
						&cli.IntFlag{
							Name:        "tcpfin",
							Aliases:     []string{"f", "fin"},
							Usage:       "set tcpfin timeout(unit second), 0 means no change",
							DefaultText: "120",
						},
						&cli.IntFlag{
							Name:        "udp",
							Aliases:     []string{"u"},
							Usage:       "set udp timeout(unit second), 0 means no change",
							DefaultText: "300",
						},
					},
				},
			},
		},
	}

	return cmds
}

// ConnectionCommands return connection relate operations, like show
func (o *Operator) ConnectionCommands() []*cli.Command {
	var cmds = []*cli.Command{
		{
			Name:    "connection",
			Aliases: []string{"c", "conn"},
			Usage:   "Operates connection",
			Subcommands: []*cli.Command{
				{
					Name:    "show",
					Aliases: []string{"ls", "get"},
					Usage:   "Shows ip_vs current connection",
					Action:  o.ShowConnection(),
					Flags: append(vsFlags2, &cli.PathFlag{
						Name:        "path",
						Aliases:     []string{"p"},
						Usage:       "specify ip_vs connection file path",
						DefaultText: "/proc/net/ip_vs_conn",
					}),
				},
			},
		},
	}

	return cmds
}
