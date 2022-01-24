package main

import "github.com/urfave/cli/v2"

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
		Required:    true,
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

// ServiceCommands return service relate operations, like get set flush import export
func (o *Operator) ServiceCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "service",
			Aliases: []string{"s", "svc", "vs"},
			Usage:   "Operates virtual service[vip:vport protocol] (TCP UDP STCP)/(IPv4 IPv6)",
			Subcommands: []*cli.Command{
				{
					Name:        "add",
					Aliases:     []string{"a", "new", "n", "set", "s"},
					Usage:       "Add ipvs virtual service",
					Description: `Add a virtual service. A service address is uniquely defined by a triplet: IP address, port number,  and  protocol.  Alternatively,  a virtual service may be defined by a firewall-mark.`,
					Action:      o.AddService(),
					Flags:       append(vsFlags, vsOptionFlags...),
				},
				{
					Name:    "del",
					Aliases: []string{"d", "delete"},
					Usage:   "Del ipvs virtual service",
					Action:  o.DelService(),
					Flags:   vsFlags,
				},
				{
					Name:    "flush",
					Aliases: []string{"f"},
					Usage:   "Flush ipvs, all the rules will be clear",
					Action:  o.FlushService(),
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
