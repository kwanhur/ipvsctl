package main

import (
	"fmt"
	"github.com/kwanhur/ipvs"
)

type Forward struct {
	flag    uint32
	forward string
}

func NewForward(flag uint32) *Forward {
	return &Forward{flag: flag}
}

func NewForward2(fwd string) *Forward {
	if fwd == "" {
		fwd = "dr"
	}
	return &Forward{forward: fwd}
}

func (f *Forward) String() string {
	switch f.flag {
	case ipvs.ConnFwdMask:
		f.forward = "mask"
	case ipvs.ConnFwdLocalNode:
		f.forward = "local"
	case ipvs.ConnFwdMasq:
		f.forward = "masq"
	case ipvs.ConnFwdBypass:
		f.forward = "bypass"
	case ipvs.ConnFwdDirectRoute:
		f.forward = "dr"
	case ipvs.ConnFwdTunnel:
		f.forward = "tun"
	case ipvs.ConnFwdFullNat:
		f.forward = "fnat"
	default:
		f.forward = fmt.Sprintf("unknown(%d)", f.flag)
	}

	return f.forward
}

func (f *Forward) Flag() uint32 {
	switch f.forward {
	case "mask":
		f.flag = ipvs.ConnFwdMask
	case "masq":
		f.flag = ipvs.ConnFwdMasq
	case "local":
		f.flag = ipvs.ConnFwdLocalNode
	case "bypass":
		f.flag = ipvs.ConnFwdBypass
	case "dr":
		f.flag = ipvs.ConnFwdDirectRoute
	case "tun":
		f.flag = ipvs.ConnFwdTunnel
	case "fnat":
		f.flag = ipvs.ConnFwdFullNat
	}

	return f.flag
}
