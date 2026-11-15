package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// appVersion is set at build time:
// go build -ldflags "-X=github.com/ivan-berezhnov/kbot/cmd.appVersion=v1.0.0"
var appVersion = "dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the kbot version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(appVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
