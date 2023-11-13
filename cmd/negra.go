package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"
	"strings"

	urfaveCli "github.com/urfave/cli/v2"
	negraCli "github.com/xenolog/negra/pkg/cli"
)

var displayFlagSet = &urfaveCli.Command{
	Name:  "debug-displayFlagSet",
	Usage: "display FlagSet",
	Action: func(_ *urfaveCli.Context) error {
		flag.CommandLine.VisitAll(func(flg *flag.Flag) {
			fmt.Printf("[%s]=[%s]\n", flg.Name, flg.Value)
		})
		return nil
	},
}

var testLogOutput = &urfaveCli.Command{
	Name:  "test",
	Usage: "test log output",
	Action: func(_ *urfaveCli.Context) error {
		slog.Debug("log-debug")
		slog.Info("log-info")
		slog.Warn("log-warning")
		slog.Error("log-error")
		return nil
	},
}

func main() {
	var exitCode int
	defer func() { os.Exit(exitCode) }() // should be first and in this form
	// defer slog.Flush()

	ctx := context.TODO()
	tmp := strings.Split(os.Args[0], "/")
	binaryName := tmp[len(tmp)-1]
	app := urfaveCli.App{
		Name:                   binaryName,
		Usage:                  "Network grapher for netplan",
		UseShortOptionHandling: true,
		// Commands:               []*urfaveCli.Command{negraCli.CmdVersion, negraCli.CmdShowConfig, displayFlagSet},
		Commands: []*urfaveCli.Command{negraCli.CmdVersion, displayFlagSet, testLogOutput},
		Before: func(ctx *urfaveCli.Context) error {
			// setup default logging
			logFlags := log.Flags()
			logFlags |= log.LUTC | log.Lmicroseconds
			if ctx.Bool("log-source") {
				logFlags |= log.Lshortfile
			}
			if strings.EqualFold("JSON", ctx.String("log-format")) {
				slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
					AddSource: ctx.Bool("log-source"),
					Level:     slog.Level(ctx.Int("log-level")),
				})))
			} else {
				log.SetFlags(logFlags)
			}
			// pass flags to native golang `flag` package
		exLoop:
			for i := range ctx.App.Flags {
				for _, name := range ctx.App.Flags[i].Names() {
					if flag.CommandLine.Lookup(name) != nil {
						continue exLoop
					}
				}
				ctx.App.Flags[i].Apply(flag.CommandLine) //nolint:errcheck
			}
			flag.Parse()
			flag.CommandLine.Set("v", strconv.Itoa(ctx.Int("log-level"))) //nolint:errcheck // klog specific setup
			return nil
		},
		Flags: []urfaveCli.Flag{ // global flags
			&urfaveCli.IntFlag{
				Name:    "log-level",
				EnvVars: []string{"LOG_LEVEL"},
				Usage:   "see constants at https://pkg.go.dev/log/slog#Level",
			},
			&urfaveCli.StringFlag{
				Name:    "log-format",
				EnvVars: []string{"LOG_FORMAT"},
				Usage:   "you able to switch to use JSON format",
			},
			&urfaveCli.BoolFlag{
				Name:    "log-source",
				EnvVars: []string{"LOG_SOURCE"},
				Usage:   "log source file and line number",
			},
		},
	}

	if err := app.RunContext(ctx, os.Args); err != nil {
		slog.ErrorContext(ctx, "Unable to run application", "err", err)
	}
	fmt.Println()
}
