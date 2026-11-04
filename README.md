# kbot

A Go Telegram bot using [Cobra](https://github.com/spf13/cobra) for its CLI and [Telebot v4](https://pkg.go.dev/gopkg.in/telebot.v4) for Telegram messages. Tasks use [Rune](https://github.com/rune-task-runner/rune), not Make.

The bot is an independent Go module in `kbot/`. The existing `gohttps/` directory is unrelated and is not built by these tasks.

## Requirements

- Go 1.26.8 or newer, matching `kbot/go.mod`.
- Rune. Install with `go install github.com/rune-task-runner/rune/cmd/rune@latest` and ensure your Go binary directory is on `PATH`.
- A Telegram bot token from [BotFather](https://t.me/BotFather).

Telebot v4 is currently pinned to `v4.0.0-beta.10` in `kbot/go.mod`. The assignment's main requirement specifies v4, so this project does not use v3 from its older installation notes.

## Create and configure the bot

1. Open BotFather in Telegram and send `/newbot`.
2. Follow the prompts to choose a display name and a unique username ending in `_bot`.
3. Store the resulting token in the environment. In Bash, this avoids recording it in shell history:

   ```bash
   read -rsp 'Telegram token: ' TELE_TOKEN; echo
   export TELE_TOKEN
   ```

Do not commit or share the token. Revoke it through BotFather if exposed. The application reads the environment directly; it does not load `.env` files.

Bot link template: `https://t.me/<your_chosen_name>_bot`.
Replace this template with the actual BotFather username before submitting the assignment. No bot username or live token has been supplied for this repository.

Optionally send `/setcommands` to BotFather and configure:

```text
start - Introduction
help - Available commands
version - Bot version
```

## Install and run

```sh
git clone https://github.com/vladtara/new-project.git
cd new-project
rune deps
rune check
rune start
```

`rune start` builds `bin/kbot` and starts long polling. Stop it with Ctrl+C or SIGTERM. Run only one polling instance per token. An existing Telegram webhook must be removed before using long polling.

The binary also works directly:

```sh
./bin/kbot --help
./bin/kbot --version
./bin/kbot start --poll-timeout=15s
```

`TELE_TOKEN` is required only for `start`. The default polling timeout is 10 seconds; accepted values are greater than zero and at most one minute. The HTTP timeout is the polling timeout plus 10 seconds.

Without Rune:

```sh
cd kbot
go test -race ./...
go vet ./...
go build -o ../bin/kbot .
../bin/kbot start
```

## Telegram examples

- `/start`: introduction.
- `/help`: available commands and supported messages.
- `/version`: `kbot 0.1.0`.
- `hello`, ignoring case and surrounding spaces: `Hello! How can I help you?`
- Other text: echoes the original message as plain text.
- Unknown slash command: suggests `/help`.
- Photo, document, or sticker: acknowledges the received type without downloading it.
- Video, audio, voice, video note, animation, location, or contact: explains that the type is unsupported.

In groups, BotFather privacy settings control which messages Telegram delivers. Use a private chat for the initial smoke test.

## Development

```sh
rune --list
rune fmt
rune test
rune vet
rune build
rune check
```

Tests use an offline Telebot instance with a local HTTP test server. They exercise registered handlers and CLI validation without a Telegram token. A live end-to-end check still requires a real bot: start it, send `/start`, `hello`, a photo, and `/missing`, then check the replies.

Handler errors are logged without raw request details to avoid disclosing tokens or message content. The bot does not intentionally persist messages or media.

## Publish

After completing the live smoke test and replacing the bot link template, commit the intended files and push your branch:

```sh
git add README.md Runefile .gitignore kbot
git commit -m "Add Telegram bot with Cobra, Telebot v4 and Rune tasks"
git push -u origin HEAD
```

Repository: https://github.com/vladtara/new-project
