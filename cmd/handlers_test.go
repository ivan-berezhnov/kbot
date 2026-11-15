package cmd

import "testing"

func TestReplyToText(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{"hello", "hello", "Hello I'm Kbot v1.0.0!"},
		{"hello with spaces and caps", "  HeLLo ", "Hello I'm Kbot v1.0.0!"},
		{"ukrainian greeting", "Привіт", "Hello I'm Kbot v1.0.0!"},
		{"version", "version", "v1.0.0"},
		{"help", "help", helpText},
		{"echo", "how are you?", "You said: how are you?"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replyToText(tt.text, "v1.0.0"); got != tt.want {
				t.Errorf("replyToText(%q) = %q, want %q", tt.text, got, tt.want)
			}
		})
	}
}
