package notifier

import (
	"fmt"
	"os"

	"github.com/benmatselby/gollum-page-watcher-action/github"
	"github.com/slack-go/slack"
)

// Slack implements the notifier strategy
type Slack struct{}

// Send communicates a message to the Slack API
func (s *Slack) Send(event *github.GollumEvent) error {
	content := ""
	for _, page := range event.Pages {
		content += fmt.Sprintf("<%s|%s>\n", page.URL, page.Title)
	}

	attachments := []slack.Attachment{slack.Attachment{
		Color:      "#2e5685",
		Text:       content,
		AuthorName: event.Sender.Login,
		AuthorIcon: event.Sender.AvatarURL,
	}}

	msg := &slack.WebhookMessage{
		Text:        "The following pages have changed in the wiki",
		Attachments: attachments,
	}

	if os.Getenv("SLACK_USERNAME") != "" {
		msg.Username = os.Getenv("SLACK_USERNAME")
	}

	if os.Getenv("SLACK_CHANNEL") != "" {
		msg.Channel = os.Getenv("SLACK_CHANNEL")
	}

	if os.Getenv("DEBUG") != "" {
		fmt.Println(msg)
		return nil
	}

	return slack.PostWebhook(os.Getenv("SLACK_WEBHOOK"), msg)
}
