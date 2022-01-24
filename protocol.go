package main

import (
	"strings"
	"syscall"
)

type IProtocol struct {
	proto string
}

func Protocol(protocol string) IProtocol {
	return IProtocol{proto: protocol}
}

func (p *IProtocol) Code() uint16 {
	switch strings.ToUpper(p.proto) {
	case "TCP":
		return syscall.IPPROTO_TCP
	case "UDP":
		return syscall.IPPROTO_UDP
	case "SCTP":
		return syscall.IPPROTO_SCTP
	default:
		return 0
	}
}

func (p *IProtocol) Support() bool {
	switch p.proto {
	case "TCP", "UDP", "SCTP":
		return true
	default:
		return false
	}
}
