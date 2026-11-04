package cmd

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/vladtara/new-project/kbot/internal/bot"
	tele "gopkg.in/telebot.v4"
)

const version = "0.1.0"

func NewRootCommand() *cobra.Command {
	var pollTimeout time.Duration
	root := &cobra.Command{
		Use:           "kbot",
		Short:         "A Telegram bot with command, text, and media handlers",
		Version:       version,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	start := &cobra.Command{
		Use:   "start",
		Short: "Start long polling (requires TELE_TOKEN)",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			token := strings.TrimSpace(os.Getenv("TELE_TOKEN"))
			if token == "" {
				return errors.New("TELE_TOKEN is required; obtain a token from @BotFather")
			}
			if pollTimeout <= 0 || pollTimeout > time.Minute {
				return errors.New("poll-timeout must be greater than zero and at most 1m")
			}
			b, err := tele.NewBot(tele.Settings{
				Token:  token,
				Poller: &tele.LongPoller{Timeout: pollTimeout},
				Client: &http.Client{Timeout: pollTimeout + 10*time.Second},
				OnError: func(_ error, _ tele.Context) {
					// Raw API errors can contain request details, so keep them out of logs.
					fmt.Fprintln(cmd.ErrOrStderr(), "Telegram request or handler failed")
				},
			})
			if err != nil {
				return errors.New("cannot initialize Telegram bot; check TELE_TOKEN and network connectivity")
			}
			bot.Register(b, version)
			fmt.Fprintln(cmd.OutOrStdout(), "Bot started. Press Ctrl+C to stop.")
			go func() {
				<-cmd.Context().Done()
				b.Stop()
			}()
			b.Start()
			return nil
		},
	}
	start.Flags().DurationVar(&pollTimeout, "poll-timeout", 10*time.Second, "Telegram long-poll timeout (maximum 1m)")
	root.AddCommand(start)
	return root
}
