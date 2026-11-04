package bot

import (
	"strings"

	tele "gopkg.in/telebot.v4"
)

const help = "Commands:\n/start - introduction\n/help - show commands\n/version - bot version\n\nSend hello for a greeting, or any other text for an echo. Photos, documents, and stickers are acknowledged."

func Register(b *tele.Bot, version string) {
	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hello! I am kbot. Send /help to see what I can do.")
	})
	b.Handle("/help", func(c tele.Context) error { return c.Send(help) })
	b.Handle("/version", func(c tele.Context) error { return c.Send("kbot " + version) })
	b.Handle(tele.OnText, handleText)
	b.Handle(tele.OnPhoto, func(c tele.Context) error { return c.Send("Photo received.") })
	b.Handle(tele.OnDocument, func(c tele.Context) error { return c.Send("Document received.") })
	b.Handle(tele.OnSticker, func(c tele.Context) error { return c.Send("Sticker received.") })
	for _, endpoint := range []string{tele.OnVideo, tele.OnAudio, tele.OnVoice, tele.OnVideoNote, tele.OnAnimation, tele.OnLocation, tele.OnContact} {
		b.Handle(endpoint, func(c tele.Context) error {
			return c.Send("This message type is not supported. Send text, a photo, a document, or a sticker.")
		})
	}
}

func handleText(c tele.Context) error {
	text := strings.TrimSpace(c.Text())
	switch {
	case strings.HasPrefix(text, "/"):
		return c.Send("Unknown command. Send /help for available commands.")
	case strings.EqualFold(text, "hello"):
		return c.Send("Hello! How can I help you?")
	case text == "":
		return c.Send("Please send a non-empty text message.")
	default:
		// Echo without a prefix to stay within Telegram's message length limit.
		return c.Send(c.Text())
	}
}
