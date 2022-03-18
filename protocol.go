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
	"strings"
	"syscall"

	"github.com/kwanhur/ipvs"
)

// IProtocol protocol wrapper
type IProtocol struct {
	proto string
}

// Protocol convert into an IProtocol wrapper
func Protocol(protocol string) IProtocol {
	return IProtocol{proto: strings.ToUpper(protocol)}
}

// Code fetch Protocol represents in syscall, non-supported is zero
func (p *IProtocol) Code() uint16 {
	switch strings.ToUpper(p.proto) {
	case "TCP", "":
		return syscall.IPPROTO_TCP
	case "UDP":
		return syscall.IPPROTO_UDP
	case "SCTP":
		return syscall.IPPROTO_SCTP
	default:
		return 0
	}
}

// IPProto convert into IPProto represents ipvs
func (p *IProtocol) IPProto() ipvs.IPProto {
	return ipvs.IPProto(p.Code())
}

// Support check it's supported or not
func (p *IProtocol) Support() bool {
	switch p.proto {
	case "", "TCP", "UDP", "SCTP":
		return true
	default:
		return false
	}
}
