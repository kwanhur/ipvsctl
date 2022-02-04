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
			info, err := lvs.Handler.GetInfo()
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
		s.AddressFamily = syscall.AF_INET6
	} else {
		s.AddressFamily = syscall.AF_INET
	}
	s.Port = vport
	s.Protocol = protocol.IPProto()
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

func (o *Operator) Zero() cli.ActionFunc {
	return func(c *cli.Context) error {
		o.ctx = c
		return o.doAction(func(lvs *IPVS) error {
			return lvs.Handler.Zero()
		})
	}
}

func (o *Operator) StringService() cli.ActionFunc {
	return func(c *cli.Context) error {
		o.ctx = c
		return o.doAction(func(lvs *IPVS) error {
			if s, err := o.service(); err != nil {
				return err
			} else {
				o.Print("vs:%s\n", s.String())

				return nil
			}
		})
	}
}

func (o *Operator) ListService() cli.ActionFunc {
	return func(c *cli.Context) error {
		o.ctx = c
		return o.doAction(func(lvs *IPVS) error {
			services, err := lvs.Handler.GetServices()
			if err != nil {
				return err
			} else {
				o.Print("Protocol Vip:Vport (Scheduler)\n")
				for _, s := range services {
					o.Print("%s\n", s.String())
				}
				return nil
			}
		})
	}
}

func (o *Operator) ExistService() cli.ActionFunc {
	return func(c *cli.Context) error {
		o.ctx = c
		return o.doAction(func(lvs *IPVS) error {
			if s, err := o.service(); err != nil {
				return err
			} else {
				if ok := lvs.Handler.IsServicePresent(s); ok {
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
				return lvs.Handler.NewService(s)
			}
		})
	}
}

func (o *Operator) UpdateService() cli.ActionFunc {
	return func(c *cli.Context) error {
		o.ctx = c
		return o.doAction(func(lvs *IPVS) error {
			if s, err := o.service(); err != nil {
				return err
			} else {
				return lvs.Handler.UpdateService(s)
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
				return lvs.Handler.DelService(s)
			}
		})
	}
}

func (o *Operator) ZeroService() cli.ActionFunc {
	return func(c *cli.Context) error {
		o.ctx = c
		return o.doAction(func(lvs *IPVS) error {
			if s, err := o.service(); err != nil {
				return err
			} else {
				return lvs.Handler.ZeroService(s)
			}
		})
	}
}

func (o *Operator) FlushService() cli.ActionFunc {
	return func(c *cli.Context) error {
		o.ctx = c
		return o.doAction(func(lvs *IPVS) error {
			return lvs.Handler.Flush()
		})
	}
}

func (o *Operator) ShowTimeout() cli.ActionFunc {
	return func(c *cli.Context) error {
		o.ctx = c
		return o.doAction(func(lvs *IPVS) error {
			cfg, err := lvs.Handler.GetConfig()
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
			return lvs.Handler.SetConfig(&cfg)
		})
	}
}
