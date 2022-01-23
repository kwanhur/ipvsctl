package main

import (
	"runtime"
	"time"

	"github.com/kwanhur/ipvs"
	"github.com/urfave/cli/v2"
)

type actionFunc func(lvs *IPVS) error

func (o *Operator) operate(callback actionFunc) {
	lvs, err := NewIPVS()
	if err != nil {
		o.Fatal("%s\n", err)
	}
	defer lvs.Close()

	if err := callback(lvs); err != nil {
		o.Fatal("%s\n", err)
	}
}

func (o *Operator) ShowVersion() func(c *cli.Context) {
	return func(c *cli.Context) {
		o.ctx = c
		o.operate(func(lvs *IPVS) error {
			info, err := lvs.Info()
			if err != nil {
				return err
			}

			o.Print("IP Virtual Server version %s (size=%d)", info.Version.String(), info.ConnTableSize)
			o.Print("\n")
			o.Print("%s %s commit id %s\n", c.App.Name, c.App.Version, CommitID)
			o.Print("Built by %s %s/%s compiler %s at %s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH, runtime.Compiler, Built)

			return nil
		})
	}
}

func (o *Operator) ShowTimeout() cli.ActionFunc {
	return func(c *cli.Context) error {
		o.ctx = c
		o.operate(func(lvs *IPVS) error {
			cfg, err := lvs.Config()
			if err != nil {
				return err
			}

			cfgs := []interface{}{cfg.TimeoutTCP.Seconds(), cfg.TimeoutTCPFin.Seconds(), cfg.TimeoutUDP.Seconds()}
			o.Print("Timeout of tcp tcpfin udp\n")
			o.Print("%.0f    %.0f    %.0f\n", cfgs...)

			return nil
		})
		return nil
	}
}

func (o *Operator) SetTimeout() cli.ActionFunc {
	return func(c *cli.Context) error {
		o.ctx = c
		o.operate(func(lvs *IPVS) error {
			tcp := c.Int("tcp")
			tcpfin := c.Int("tcpfin")
			udp := c.Int("udp")

			cfg := ipvs.Config{
				TimeoutTCP:    time.Duration(tcp) * time.Second,
				TimeoutTCPFin: time.Duration(tcpfin) * time.Second,
				TimeoutUDP:    time.Duration(udp) * time.Second,
			}
			return lvs.SetConfig(&cfg)
		})
		return nil
	}
}
