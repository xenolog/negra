package cli

import (
	"fmt"

	"github.com/urfave/cli"
	botConfig "github.com/xenolog/negra/pkg/config"
)

// ------------------------------------------------------------------------

// var showCmd = &cli.Command{
// 	Name:  "show",
// 	Usage: "show something...",
// 	PersistentPreRun: func(cmd *cli.Command, args []string) {
// 		fmt.Println("Loading config file.")
// 		if err := botConfig.BotConfig.Parse(botConfig.ConfigFilePath); err != nil {
// 			fmt.Println(err)
// 		}
// 	},
// }

var CmdShowConfig = &cli.Command{
	Name:  "config",
	Usage: "show config",
	Action: func(_ *cli.Context) error {
		fmt.Println("BOT config is:")
		fmt.Println(botConfig.BotConfig.String())
		return nil
	},
}
