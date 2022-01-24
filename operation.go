package main

import (
	"fmt"
	"net"
	"runtime"
	"time"

	"github.com/kwanhur/ipvs"
	"github.com/urfave/cli/v2"
)

type actionFunc func(lvs *IPVS) error

func (o *Operator) show(callback actionFunc) {
	lvs, err := NewIPVS()
	if err != nil {
		o.Fatal("%s\n", err)
	}
	defer lvs.Close()

	if err := callback(lvs); err != nil {
		o.Fatal("%s\n", err)
	}
}

func (o Operator) doAction(callback actionFunc) error {
	lvs, err := NewIPVS()
	if err != nil {
		return err
	}
	defer lvs.Close()

	return callback(lvs)
}

func (o *Operator) ShowVersion() func(c *cli.Context) {
	return func(c *cli.Context) {
		o.ctx = c
		o.show(func(lvs *IPVS) error {
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

func (o *Operator) AddService() cli.ActionFunc {
	return func(c *cli.Context) error {
		o.ctx = c
		return o.doAction(func(lvs *IPVS) error {
			vip := net.ParseIP(c.String("vip"))
			if vip == nil {
				return fmt.Errorf("invalid vip address %s\n", vip)
			}
			vport := uint16(c.Uint("vport"))
			protocol := Protocol(c.String("protocol"))
			if !protocol.Support() {
				return fmt.Errorf("invalid protocol %s", protocol)
			}

			s := ipvs.Service{}
			s.Address = vip
			s.Port = vport
			s.Protocol = protocol.Code()

			return lvs.AddService(&s)
		})
	}
}

func (o *Operator) FlushService() cli.ActionFunc {
	return func(c *cli.Context) error {
		o.ctx = c
		return o.doAction(func(lvs *IPVS) error {
			return lvs.Flush()
		})
	}
}

func (o *Operator) ShowTimeout() cli.ActionFunc {
	return func(c *cli.Context) error {
		o.ctx = c
		return o.doAction(func(lvs *IPVS) error {
			cfg, err := lvs.Config()
			if err != nil {
				return err
			}

			cfgs := []interface{}{cfg.TimeoutTCP.Seconds(), cfg.TimeoutTCPFin.Seconds(), cfg.TimeoutUDP.Seconds()}
			o.Print("Timeout of tcp tcpfin udp\n")
			o.Print("%.0f    %.0f    %.0f\n", cfgs...)

			return nil
		})
	}
}

func (o *Operator) SetTimeout() cli.ActionFunc {
	return func(c *cli.Context) error {
		o.ctx = c
		return o.doAction(func(lvs *IPVS) error {
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
	}
}
