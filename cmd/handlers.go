package cmd

import (
	"fmt"
	"strings"

	telebot "gopkg.in/telebot.v4"
)

const helpText = `Available commands:
/start - start the conversation
/help - show this help
hello - get a greeting
version - show the bot version

Send any other text and kbot will echo it back.`

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

func senderName(m telebot.Context) string {
	if s := m.Sender(); s != nil && s.FirstName != "" {
		return s.FirstName
	}
	return "there"
}
