package ezhooks

// The models required for a discord webhook instance itself.
type (
	// Webhook is the actual discord webhook instance itself.
	Webhook struct {
		Content string   `json:"content"`
		Embeds  []*Embed `json:"embeds"`
	}

	// Embed model contains the information for an embed field.
	Embed struct {
		Title       string     `json:"title"`
		Description string     `json:"description"`
		Colour      int        `json:"color"`
		Fields      []*Field   `json:"fields"`
		Author      *Author    `json:"author"`
		Footer      *Footer    `json:"footer"`
		Thumbnail   *Thumbnail `json:"thumbnail"`
	}

	// Author contains the information for a discord embed author.
	Author struct {
		Name string `json:"name"`
	}

	// Footer contains the information for the embed footer.
	Footer struct {
		Text    string `json:"text"`
		IconURL string `json:"icon_url"`
	}

	// Thumbnail contains the information for the discord embed thumbnail.
	Thumbnail struct {
		URL string `json:"url"`
	}

	// Field contains the information for a discord embed field.
	Field struct {
		Name   string `json:"name"`
		Value  string `json:"value"`
		Inline bool   `json:"inline"`
	}
)

// DiscordHook contains the information for the discord webhook itself.
type DiscordHook struct {
	WebhookID  int
	WebhookKey string
	webhook    *Webhook
}

// NewDiscordHook will create a new discord webhook.
func NewDiscordHook(i int, k string) *DiscordHook {
	return &DiscordHook{
		WebhookID:  i,
		WebhookKey: k,
	}
}

// AddWebhook will add a discord webhook to the discordhook instance.
func (d *DiscordHook) AddWebhook(w *Webhook) {
	d.webhook = w
}
