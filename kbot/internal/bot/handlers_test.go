package bot

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	tele "gopkg.in/telebot.v4"
)

func TestRegisteredHandlers(t *testing.T) {
	tests := []struct {
		name    string
		message tele.Message
		want    string
	}{
		{"start", tele.Message{Text: "/start"}, "Hello! I am kbot. Send /help to see what I can do."},
		{"help", tele.Message{Text: "/help"}, help},
		{"version", tele.Message{Text: "/version"}, "kbot test"},
		{"greeting", tele.Message{Text: "  HeLLo  "}, "Hello! How can I help you?"},
		{"echo", tele.Message{Text: "<b>plain text</b>"}, "<b>plain text</b>"},
		{"unknown", tele.Message{Text: "/missing"}, "Unknown command. Send /help for available commands."},
		{"blank", tele.Message{Text: "   "}, "Please send a non-empty text message."},
		{"photo", tele.Message{Photo: &tele.Photo{}}, "Photo received."},
		{"document", tele.Message{Document: &tele.Document{}}, "Document received."},
		{"sticker", tele.Message{Sticker: &tele.Sticker{}}, "Sticker received."},
		{"voice", tele.Message{Voice: &tele.Voice{}}, "This message type is not supported. Send text, a photo, a document, or a sticker."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/bottest/sendMessage" {
					t.Errorf("unexpected API request: %s", r.URL.Path)
				}
				var payload struct {
					Text string `json:"text"`
				}
				if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
					t.Error(err)
				}
				got = payload.Text
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"ok":true,"result":{"message_id":1,"chat":{"id":1,"type":"private"}}}`))
			}))
			defer server.Close()
			b, err := tele.NewBot(tele.Settings{Token: "test", URL: server.URL, Offline: true, Synchronous: true, OnError: func(err error, _ tele.Context) { t.Error(err) }})
			if err != nil {
				t.Fatal(err)
			}
			Register(b, "test")
			tt.message.Chat = &tele.Chat{ID: 1, Type: tele.ChatPrivate}
			tt.message.Sender = &tele.User{ID: 1}
			b.ProcessUpdate(tele.Update{Message: &tt.message})
			if got != tt.want {
				t.Errorf("reply = %q, want %q", got, tt.want)
			}
		})
	}
}
