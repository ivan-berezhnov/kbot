package cmd

import (
	"fmt"
	"log"
	"strings"

	telebot "gopkg.in/telebot.v4"
)

const helpText = `Available commands:
/start - start the conversation
/help - show this help
hello - get a greeting
version - show the bot version

Send any other text and kbot will echo it back.`

// registerHandlers adds message handlers to the bot.
func registerHandlers(kbot *telebot.Bot) {
	kbot.Handle("/start", func(c telebot.Context) error {
		return c.Send(fmt.Sprintf("Hi, %s! I'm Kbot %s.\n\n%s", senderName(c), appVersion, helpText))
	})

	kbot.Handle("/help", func(c telebot.Context) error {
		return c.Send(helpText)
	})

	kbot.Handle(telebot.OnText, func(c telebot.Context) error {
		log.Printf("Received text: %q", c.Text())
		return c.Send(replyToText(c.Text(), appVersion))
	})

	kbot.Handle(telebot.OnPhoto, func(c telebot.Context) error {
		return c.Send("Nice photo! I can only read text for now.")
	})

	kbot.Handle(telebot.OnSticker, func(c telebot.Context) error {
		return c.Send("Cool sticker! Send me \"hello\" to talk.")
	})
}

// replyToText returns the bot answer for a text message.
func replyToText(text, version string) string {
	switch strings.ToLower(strings.TrimSpace(text)) {
	case "hello", "hi", "привіт":
		return fmt.Sprintf("Hello I'm Kbot %s!", version)
	case "version":
		return version
	case "help":
		return helpText
	default:
		return fmt.Sprintf("You said: %s", text)
	}
}

func senderName(c telebot.Context) string {
	if s := c.Sender(); s != nil && s.FirstName != "" {
		return s.FirstName
	}
	return "there"
}
