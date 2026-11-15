package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v4"
)

var (
	// TeleToken is the Telegram bot token from the TELE_TOKEN environment variable.
	TeleToken = os.Getenv("TELE_TOKEN")
)

// kbotCmd represents the kbot command.
var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "Start the Telegram bot",
	Long: `Start the Telegram bot and answer user messages.

The bot token is read from the TELE_TOKEN environment variable.`,
	Run: func(cmd *cobra.Command, args []string) {
		if TeleToken == "" {
			log.Fatal("TELE_TOKEN environment variable is not set")
		}

		fmt.Printf("kbot %s started\n", appVersion)

		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})
		if err != nil {
			log.Fatalf("Please check TELE_TOKEN env variable. %s", err)
			return
		}

		kbot.Handle("/start", func(m telebot.Context) error {
			return m.Send(fmt.Sprintf("Hi, %s! I'm Kbot %s.\n\n%s", senderName(m), appVersion, helpText))
		})

		kbot.Handle("/help", func(m telebot.Context) error {
			return m.Send(helpText)
		})

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			log.Printf("Received message: %s", m.Text())
			return m.Send(replyToText(m.Text(), appVersion))
		})

		kbot.Handle(telebot.OnPhoto, func(m telebot.Context) error {
			return m.Send("Nice photo! I can only read text for now.")
		})

		kbot.Handle(telebot.OnSticker, func(m telebot.Context) error {
			return m.Send("Cool sticker! Send me \"hello\" to talk.")
		})

		kbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)
}
