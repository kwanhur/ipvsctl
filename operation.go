package main

import (
	"fmt"
	"net"
	"runtime"
	"syscall"
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

func (o *Operator) service() (*ipvs.Service, error) {
	vip := net.ParseIP(o.ctx.String("vip"))
	if vip == nil {
		return nil, fmt.Errorf("invalid vip address %s\n", vip)
	}
	vport := uint16(o.ctx.Uint("vport"))
	protocol := Protocol(o.ctx.String("protocol"))
	if !protocol.Support() {
		return nil, fmt.Errorf("invalid protocol %s", protocol)
	}
	sched := o.ctx.String("scheduler")
	pe := o.ctx.String("persistent")
	timeout := uint32(o.ctx.Uint("timeout"))
	netmask := uint32(o.ctx.Uint("netmask"))

	s := ipvs.Service{}
	s.Address = vip
	addrIPv6 := vip.To4() == nil
	if addrIPv6 {
		s.AddressFamily = syscall.IPPROTO_IPV6
	} else {
		s.AddressFamily = syscall.IPPROTO_IP
	}
	s.Port = vport
	s.Protocol = ipvs.IPProto(protocol.Code())
	s.SchedName = sched
	s.PEName = pe
	s.Timeout = timeout
	if addrIPv6 && (netmask == 0 || netmask > 128) {
		netmask = 128
	} else if !addrIPv6 && (netmask == 0 || netmask > 32) {
		netmask = 32
	}
	s.Netmask = netmask

	return &s, nil
}

func (o *Operator) ExistService() cli.ActionFunc {
	return func(c *cli.Context) error {
		o.ctx = c
		return o.doAction(func(lvs *IPVS) error {
			if s, err := o.service(); err != nil {
				return err
			} else {
				if ok := lvs.ExistService(s); ok {
					o.Print("vs:%s found\n", s.String())
				} else {
					o.Print("vs:%s not found\n", s.String())
				}

				return nil
			}
		})
	}
}

func (o *Operator) AddService() cli.ActionFunc {
	return func(c *cli.Context) error {
		o.ctx = c
		return o.doAction(func(lvs *IPVS) error {
			if s, err := o.service(); err != nil {
				return err
			} else {
				return lvs.AddService(s)
			}
		})
	}
}

func (o *Operator) DelService() cli.ActionFunc {
	return func(c *cli.Context) error {
		o.ctx = c
		return o.doAction(func(lvs *IPVS) error {
			if s, err := o.service(); err != nil {
				return err
			} else {
				return lvs.DelService(s)
			}
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
