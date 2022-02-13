package client

import (
	"github.com/aiomonitors/godiscord"
)

// DiscordWebhookClient implements DiscordWebhookClientInterface
// It provides methods to access Discord services via RESTFul APIs
type DiscordWebhookClient struct {
	WebhookURL string
}

// PostStatsToDiscord sends stats to Discord
func (d *DiscordWebhookClient) PostStatsToDiscord() {
	embed := godiscord.NewEmbed(
		"TNB Analysis summary <DATE>",
		"Daily dose of insights about thenewboston digital crypto currency network analysis like total coins distributed, richest account, wealth distribution, etc.",
		"https://tnbexplorer.com/tnb/stats",
	)

	embed.AddField("Total coins distributed", "<Total>", false)
	// embed.SetColor("#F1B379")
	embed.SetAuthor("TNB Explorer", "https://tnbexplorer.com/tnb", "https://itsnikhil.github.io/tnb-analysis/web/assets/maskable_icon.png")
	embed.SetFooter("visit website for further information like charts and rich list", "")

	embed.SendToWebhook(d.WebhookURL)
}
