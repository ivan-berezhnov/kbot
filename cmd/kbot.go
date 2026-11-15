package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v4"
)

var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "Start the Telegram bot",
	Long: `Start the Telegram bot and answer user messages.

The bot token is read from the TELE_TOKEN environment variable.`,
	Run: func(cmd *cobra.Command, args []string) {
		token := os.Getenv("TELE_TOKEN")
		if token == "" {
			log.Fatal("TELE_TOKEN environment variable is not set")
		}

		fmt.Printf("kbot %s started\n", appVersion)

		kbot, err := telebot.NewBot(telebot.Settings{
			Token:  token,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})
		if err != nil {
			log.Fatalf("Please check TELE_TOKEN env variable. %s", err)
		}

		registerHandlers(kbot)
		kbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)
}
