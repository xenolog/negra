package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	botConfig "github.com/xenolog/negra/pkg/config"
)

// ------------------------------------------------------------------------

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show something...",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Loading config file.")
		if err := botConfig.BotConfig.Parse(botConfig.ConfigFilePath); err != nil {
			fmt.Println(err)
		}
	},
}

var cmdShowConfig = &cobra.Command{
	Use:   "config",
	Short: "show config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("BOT config is:")
		fmt.Println(botConfig.BotConfig.String())
	},
}
