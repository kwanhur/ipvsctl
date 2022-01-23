package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

type Operator struct {
	ctx *cli.Context
}

func NewOperator() *Operator {
	return &Operator{}
}

func (o *Operator) Commands() []*cli.Command {
	cmds := o.TimeoutCommands()
	return cmds
}

func (o *Operator) Fatal(format string, a ...interface{}) {
	fmt.Fprintf(o.ctx.App.ErrWriter, format, a...)
	os.Exit(2)
}

func (o *Operator) Print(format string, a ...interface{}) {
	fmt.Fprintf(o.ctx.App.Writer, format, a...)
}
