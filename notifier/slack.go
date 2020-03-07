package notifier

import (
	"fmt"
	"os"

	"github.com/benmatselby/gollum-page-watcher-action/config"
	"github.com/benmatselby/gollum-page-watcher-action/github"
	"github.com/slack-go/slack"
)

// Slack implements the notifier strategy
type Slack struct{}

// Send communicates a message to the Slack API
func (s *Slack) Send(config config.Config, event *github.GollumEvent) error {
	content := ""
	for _, page := range event.Pages {
		content += fmt.Sprintf("<%s|%s>\n", page.URL, page.Title)
	}

	attachments := []slack.Attachment{{
		Color:      "#2e5685",
		Text:       content,
		AuthorName: event.Sender.Login,
		AuthorIcon: event.Sender.AvatarURL,
	}}

	msg := &slack.WebhookMessage{
		Text:        "The following pages have changed in the wiki",
		Attachments: attachments,
	}

	if config.SlackUsername != "" {
		msg.Username = config.SlackUsername
	}

	if config.SlackChannel != "" {
		msg.Channel = config.SlackChannel
	}

	if config.Debug != "" {
		fmt.Println(msg)
		return nil
	}

	return slack.PostWebhook(os.Getenv("SLACK_WEBHOOK"), msg)
}
