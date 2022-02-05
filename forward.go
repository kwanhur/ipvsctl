package main

import (
	"fmt"
	"github.com/kwanhur/ipvs"
)

func Forward(flag uint32) string {
	var fwd string
	switch flag {
	case ipvs.ConnFwdMask:
		fwd = "mask"
	case ipvs.ConnFwdLocalNode:
		fwd = "local"
	case ipvs.ConnFwdMasq:
		fwd = "masq"
	case ipvs.ConnFwdBypass:
		fwd = "bypass"
	case ipvs.ConnFwdDirectRoute:
		fwd = "dr"
	case ipvs.ConnFwdTunnel:
		fwd = "tun"
	case ipvs.ConnFwdFullNat:
		fwd = "fnat"
	default:
		fwd = fmt.Sprintf("unknown(%d)", flag)
	}

	return fwd
}
