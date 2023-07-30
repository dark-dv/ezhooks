package ezhooks

import "net/url"

// TelegramHook contains the information for the telegram webhook itself.
type TelegramHook struct {
	WebhookID  int
	WebhookKey string
	Content    string
}

// NewTelegramHook will create a new telegram webhook.
func NewTelegramHook(i int, k, c string) *TelegramHook {
	return &TelegramHook{
		WebhookID:  i,
		WebhookKey: k,
		Content:    url.PathEscape(c),
	}
}
