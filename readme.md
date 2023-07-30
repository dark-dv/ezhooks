# ezhooks

## What is ezhooks?

**ezhooks** is a webhook package that uses both telegram and discord webhooks to run. You can use the full potential of both webhook objects for both platforms with this package. Also, it is extremley easy to interpet your own platforms into this package also.

## Example

```go
package main

import (
	"ezhooks"
)

func main() {
	e := ezhooks.NewEzHook()
	
	d := ezhooks.NewDiscordHook(00000000000000, "key-here")
	d.AddWebhook(&ezhooks.Webhook{
		Content: "Welcome to the __**ezhooks**__ example!",
		Embeds: []*ezhooks.Embed{
			{
				Title:       "This is a webhook.",
				Description: "You can have text content, or embedded content as shown in this preview!",
				Colour:      1146986,
				Fields: []*ezhooks.Field{
					{
						Name:   "This is an embed field.",
						Value:  "I am the value of said embed field.",
						Inline: true,
					},
					{
						Name:   "This is another embed field.",
						Value:  "I am the value of said other embed field.",
						Inline: true,
					},
					{
						Name:   "Let's go golfing!",
						Value:  "Have you ever played rugby?",
						Inline: true,
					},
				},
				Author: &ezhooks.Author{
					Name: "Dark - The author",
				},
				Footer: &ezhooks.Footer{
					Text:    "This is the footer of the embed.",
					IconURL: "https://creator-hub-prod.s3.us-east-2.amazonaws.com/constructs_pfp_1687920400601.jpeg",
				},
				Thumbnail: &ezhooks.Thumbnail{
					URL: "https://forum.dark-gaming.com/uploads/default/original/2X/a/ae85d87e3563811f54cb24b19790a0dba5ba169c.png",
				},
			},
		},
	})
	e.AddDiscordHook(d)

	if err := e.SendDiscord(); err != nil {
		panic(err)
	}
}
```

Output:

![Output Showcase Image]([output.png](https://raw.githubusercontent.com/dark-dv/ezhooks/main/documents/output.png)https://raw.githubusercontent.com/dark-dv/ezhooks/main/documents/output.png)
