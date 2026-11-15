package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is the base command when kbot is called without subcommands.
var rootCmd = &cobra.Command{
	Use:   "kbot",
	Short: "Kbot is a simple Telegram bot written in Go",
	Long: `Kbot is a Telegram bot built with cobra and telebot.v4.

Run "kbot start" with the TELE_TOKEN environment variable set
to connect the bot to Telegram and answer user messages.`,
}

// Execute runs the root command. It is called once from main.main().
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
