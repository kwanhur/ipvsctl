// +build linux

package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/urfave/cli/v2"
)

var (
	Version  string
	CommitID string
	Built    string
)

func main() {
	app := cli.NewApp()
	app.Usage = "IP Virtual Server controller"
	app.Version = Version
	app.Description = "ipvs controller communicate with ip_vs kernel module"
	app.Authors = Authors()

	cli.VersionPrinter = func(c *cli.Context) {
		lvs, err := NewIPVS()
		if err != nil {
			fmt.Fprintf(c.App.ErrWriter, "%s\n", err)
			os.Exit(2)
		}
		defer lvs.Close()

		info, err := lvs.Info()
		if err != nil {
			fmt.Fprintf(c.App.ErrWriter, "%s\n", err)
			os.Exit(2)
		}

		fmt.Fprintf(c.App.Writer, "IP Virtual Server version %s (size=%d)", info.Version.String(), info.ConnTableSize)
		fmt.Println()
		fmt.Fprintf(c.App.Writer, "%s %s commit id %s\n", c.App.Name, c.App.Version, CommitID)
		fmt.Fprintf(c.App.Writer, "Built by %s %s/%s compiler %s at %s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH, runtime.Compiler, Built)
	}

	app.RunAndExitOnError()
}
