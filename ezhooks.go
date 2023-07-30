package ezhooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Ezhook contains the information for the webhooks that we have created.
type Ezhook struct {
	discord  *DiscordHook
	telegram *TelegramHook
}

// NewEzHook will create a new Ezhook instance.
func NewEzHook() *Ezhook {
	return new(Ezhook)
}

// AddDiscordHook will add a discord hook to the ezhook instance.
func (e *Ezhook) AddDiscordHook(d *DiscordHook) {
	e.discord = d
}

// AddTelegramHook will add a Telegram hook to the ezhook instance.
func (e *Ezhook) AddTelegramHook(t *TelegramHook) {
	e.telegram = t
}

// SendDiscord will send a discord webhook.
func (e *Ezhook) SendDiscord() error {
	if e.discord.webhook == nil {
		return fmt.Errorf("no webhook to send")
	}

	// Compile the webhook itself into a byte array ready for the web request.
	data, err := json.Marshal(e.discord.webhook)
	if err != nil {
		return err
	}

	// Create a new HTTP client.
	c := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Create a new http post request.
	request, err := http.NewRequest("POST", fmt.Sprintf("https://discord.com/api/webhooks/%d/%s", e.discord.WebhookID, e.discord.WebhookKey), bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	// Set the new request header, to match the discord webhook specification.
	request.Header.Set("Content-Type", "application/json")

	res, err := c.Do(request)
	if err != nil {
		return err
	}

	// Check if the status code equates to a successfull attempt.
	if res.StatusCode >= 205 {
		return fmt.Errorf("bad status code presented to the request handler")
	}

	return nil
}

// SendTelegram will send a telegram webhook.
func (e *Ezhook) SendTelegram() error {
	if e.telegram == nil {
		return fmt.Errorf("no telegram webhook to send")
	}

	// Create a new http client.
	c := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Create a new http get request.
	request, err := http.NewRequest("GET", fmt.Sprintf("https://api.telegram.org/bot/%s/sendMessage?chat_id=%d&text=%s", e.telegram.WebhookKey, e.telegram.WebhookID, e.telegram.Content), nil)
	if err != nil {
		return err
	}

	res, err := c.Do(request)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("bad status code, maybe sent")
	}

	return nil
}
