package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	ngConfig "github.com/xenolog/negra/pkg/config"
	"k8s.io/klog"
)

func versionInfo() []string {
	rv := []string{
		fmt.Sprintf("Version: %s", ngConfig.GetVersion()),
		fmt.Sprintf("Go Version: %s", runtime.Version()),
		fmt.Sprintf("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH),
	}
	return rv
}

// func logVersion() {
// 	for _, s := range versionInfo() {
// 		klog.Info(s)
// 	}
// }

var RootCmd = &cobra.Command{
	// SuggestionsMinimumDistance: 2,
	Use: func() string {
		a := strings.Split(os.Args[0], "/")
		return a[len(a)-1]
	}(),
	Version: ngConfig.GetVersion(),
}

var cmdVersion = &cobra.Command{
	Use:     "version",
	Aliases: []string{"ver"},
	Short:   "show version",
	Run: func(cmd *cobra.Command, args []string) {
		for _, s := range versionInfo() {
			fmt.Fprintln(os.Stderr, s)
		}
	},
}

func init() {
	// for _, e := range os.Environ() {
	// 	pair := strings.SplitN(e, "=", 2)
	// 	switch pair[0] {
	// 	case netplanConfigPath:
	// 		if pair[1] != "" {
	// 			kiConfig.NetconfigNetplanPath = pair[1]
	// 		}
	// case reconcileTimeoutEnv:
	// 	if v, err := strconv.Atoi(pair[1]); err != nil || v < 0 {
	// 		klog.Fatalf("The %s env contains wrong value: %s", reconcileTimeoutEnv, err)
	// 	} else {
	// 		reconcileTimeout = time.Duration(v) * time.Second
	// 	}
	// }
	// }

	// Connect standart log, flag and klog together. Cobra will inhrits global CLI flags
	klog.InitFlags(flag.CommandLine)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	// ------------------------------------------------------------------------

	// RootCmd.Flags().StringVar(&ngConfig.ConfigFilePath, "config-", ngConfig.DefaultConfigFile, "Configuration file path")
	// RootCmd.Flags().DurationVar(&botConfig.PollTimeout, "poll-timeout", pollTimeout, fmt.Sprintf("timeout to reconcile objects (seconds, ENV variable %s)", pollTimeoutEnv))
	// showCmd.AddCommand(cmdShowConfig)
	// RootCmd.AddCommand(showCmd)
	// RootCmd.AddCommand(listCmd)
	// RootCmd.AddCommand(dbCmd)
	RootCmd.AddCommand(cmdVersion)
}
