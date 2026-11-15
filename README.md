# kbot

Kbot is a simple Telegram bot written in Go.
It uses [cobra](https://github.com/spf13/cobra) for the command-line interface
and [telebot.v4](https://gopkg.in/telebot.v4) for the Telegram Bot API.

Bot link: [t.me/ivokbot](https://t.me/ivokbot)

## Features

- `kbot start` connects the bot to Telegram (long polling).
- `kbot version` prints the build version.
- Message handlers by type and content:
  - commands `/start` and `/help`;
  - text: `hello` (also `hi`, `привіт`), `version`, `help`; any other text is echoed back;
  - photos and stickers get a short reply.

## Requirements

- Go 1.24 or newer
- A Telegram bot token from [@BotFather](https://t.me/BotFather)

## Installation

1. Create a bot: open [@BotFather](https://t.me/BotFather), send `/newbot`, choose a name
   and a username that ends with `_bot`. Copy the token.
2. Clone and build:

   ```bash
   git clone https://github.com/ivan-berezhnov/kbot.git
   cd kbot
   go build -ldflags "-X=github.com/ivan-berezhnov/kbot/cmd.appVersion=v1.0.0"
   ```

3. Set the token without saving it in the shell history, then start the bot:

   ```bash
   read -s TELE_TOKEN    # paste the token and press Enter
   export TELE_TOKEN
   ./kbot start
   ```

Do not commit the token to the repository.

## Usage

CLI:

```bash
./kbot --help      # list commands
./kbot version     # v1.0.0
./kbot start       # start the bot (alias of "kbot kbot")
```

In Telegram:

| You send       | Kbot answers                    |
|----------------|---------------------------------|
| `/start`       | Greeting and the list of commands |
| `/help`        | The list of commands            |
| `hello`        | `Hello I'm Kbot v1.0.0!`        |
| `version`      | `v1.0.0`                        |
| `how are you?` | `You said: how are you?`        |
| a photo        | `Nice photo! I can only read text for now.` |

## Development

```bash
go vet ./...
go test ./...
```

## Project structure

```
main.go              entry point
cmd/root.go          root cobra command
cmd/version.go       "version" command
cmd/kbot.go          "kbot" / "start" command, creates the telebot.Bot
cmd/handlers.go      message handlers
cmd/handlers_test.go tests for the text replies
```
