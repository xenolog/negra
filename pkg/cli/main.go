package cli

import (
	"fmt"
	"os"
	"runtime"

	"github.com/urfave/cli"
	ngConfig "github.com/xenolog/negra/pkg/config"
)

func versionInfo() []string {
	rv := []string{
		fmt.Sprintf("Version: %s", ngConfig.GetVersion()),
		fmt.Sprintf("Go Version: %s", runtime.Version()),
		fmt.Sprintf("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH),
	}
	return rv
}

// var RootCmd = &cli.Command{
// 	// SuggestionsMinimumDistance: 2,
// 	Use: func() string {
// 		a := strings.Split(os.Args[0], "/")
// 		return a[len(a)-1]
// 	}(),
// 	Version: ngConfig.GetVersion(),
// }

var CmdVersion = &cli.Command{
	Name:    "version",
	Aliases: []string{"ver"},
	Usage:   "show version",
	Action: func(_ *cli.Context) error {
		for _, s := range versionInfo() {
			fmt.Fprintln(os.Stderr, s)
		}
		return nil
	},
}
