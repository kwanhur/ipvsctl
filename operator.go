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
	"os"

	"github.com/urfave/cli/v2"
)

type Operator struct {
	ctx *cli.Context
}

// NewOperator operator instance, include all commands (basic, service, server, timeout and daemon)
func NewOperator() *Operator {
	return &Operator{}
}

// Commands return supported commands
func (o *Operator) Commands() []*cli.Command {
	var cmds []*cli.Command
	cmds = append(cmds, o.BasicCommands()...)
	cmds = append(cmds, o.ServiceCommands()...)
	cmds = append(cmds, o.ServerCommands()...)
	cmds = append(cmds, o.TimeoutCommands()...)
	return cmds
}

// Fatal print error message into STDERR, then exit
func (o *Operator) Fatal(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(o.ctx.App.ErrWriter, format, a...)
	os.Exit(2)
}

// Print output message into STDOUT
func (o *Operator) Print(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(o.ctx.App.Writer, format, a...)
}
