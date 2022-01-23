package main

import "github.com/urfave/cli/v2"

// ServiceCommands return service relate operations, like get set flush import export
func (o *Operator) ServiceCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "service",
			Aliases: []string{"s", "svc", "vs"},
			Usage:   "Operates virtual service[vip:vport protocol] (TCP UDP STCP)/(IPv4 IPv6)",
			Subcommands: []*cli.Command{
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
