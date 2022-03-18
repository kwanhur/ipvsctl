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
