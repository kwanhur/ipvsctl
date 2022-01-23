package main

import "github.com/urfave/cli/v2"

// TimeoutCommands return timeout relate operation, like get set
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
					Name:    "set",
					Aliases: []string{"s"},
					Usage:   "Sets timeout of tcp tcpfin udp",
					Action:  o.SetTimeout(),
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
			Flags: nil,
		},
	}

	return cmds
}
