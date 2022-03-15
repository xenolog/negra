package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli"
	urfaveCli "github.com/urfave/cli"
	negraCli "github.com/xenolog/negra/pkg/cli"
	"k8s.io/klog"
)

var logLevel uint

var displayFlagSet = &cli.Command{
	Name:    "displayFlagSet",
	Aliases: []string{"dfs"},
	Usage:   "display FlagSet",
	Action: func(_ *urfaveCli.Context) error {
		flag.CommandLine.VisitAll(func(flg *flag.Flag) {
			fmt.Printf("[%s]=[%s]\n", flg.Name, flg.Value)
		})
		return nil
	},
}

func main() {
	var exitCode int
	defer func() { os.Exit(exitCode) }() // should be first and in this form
	defer klog.Flush()

	ctx := context.TODO()
	tmp := strings.Split(os.Args[0], "/")
	binaryName := tmp[len(tmp)-1]
	app := urfaveCli.App{
		Name:                   binaryName,
		Usage:                  "Network grapher for netplan",
		UseShortOptionHandling: true,
		Commands:               []*urfaveCli.Command{negraCli.CmdVersion, negraCli.CmdShowConfig, displayFlagSet},
		Before: func(ctx *urfaveCli.Context) error {
			klog.InitFlags(flag.CommandLine)
		exLoop:
			for i := range ctx.App.Flags {
				for _, name := range ctx.App.Flags[i].Names() {
					if flag.CommandLine.Lookup(name) != nil {
						continue exLoop
					}
				}
				ctx.App.Flags[i].Apply(flag.CommandLine)
			}
			flag.Parse()
			flag.CommandLine.Set("v", strconv.Itoa(ctx.Int("log-level"))) // klog specific setup
			return nil
		},
		Flags: []urfaveCli.Flag{ // global flags
			&urfaveCli.UintFlag{
				Name:    "log-level",
				EnvVars: []string{"LOG_LEVEL"},
			},
			&urfaveCli.BoolFlag{
				Name:    "logtostderr",
				EnvVars: []string{"LOG_TO_STDERR"},
				Value:   true,
			},
			&urfaveCli.BoolFlag{
				Name:    "alsologtostderr",
				EnvVars: []string{"ALSO_LOG_TO_STDERR"},
			},
		},
	}

	app.RunContext(ctx, os.Args)
	fmt.Println()
}
