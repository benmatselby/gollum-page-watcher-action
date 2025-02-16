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
	content := "The following pages have changed in the wiki:\n"
	for _, page := range event.Pages {
		content += fmt.Sprintf("\t- <%s|%s>\n", page.URL, page.Title)
	}

	attachments := []slack.Attachment{{
		Color:      "#2e5685",
		Text:       content,
		AuthorName: event.Sender.Login,
		AuthorIcon: event.Sender.AvatarURL,
		Footer:     fmt.Sprintf("<%s|%s>", event.Repo.URL, event.Repo.FullName),
	}}

	msg := &slack.WebhookMessage{
		Attachments: attachments,
	}

	if config.SlackUsername != "" {
		msg.Username = config.SlackUsername
	}

	if config.SlackChannel != "" {
		msg.Channel = config.SlackChannel
	}

	if config.Debug != "" {
		fmt.Println("Slack via webhook")
		fmt.Println(config)
		fmt.Println(msg)
		return nil
	}

	return slack.PostWebhook(os.Getenv("SLACK_WEBHOOK"), msg)
}

// Slack implements the notifier strategy
type SlackViaToken struct{}

// Send communicates with a token to the Slack API
func (s *SlackViaToken) Send(config config.Config, event *github.GollumEvent) error {
	content := "The following pages have changed in the wiki:\n"
	for _, page := range event.Pages {
		content += fmt.Sprintf("\t- <%s|%s>\n", page.URL, page.Title)
	}

	attachments := []slack.Attachment{{
		Color:      "#2e5685",
		Text:       content,
		AuthorName: event.Sender.Login,
		AuthorIcon: event.Sender.AvatarURL,
		Footer:     fmt.Sprintf("<%s|%s>", event.Repo.URL, event.Repo.FullName),
	}}

	if config.Debug != "" {
		fmt.Println("Slack via token")
		fmt.Println(config)
		fmt.Println(attachments)
		return nil
	}

	api := slack.New(config.SlackToken)
	channel, _, _, err := api.SendMessage(
		config.SlackChannel, slack.MsgOptionAttachments(attachments...),
		slack.MsgOptionAsUser(true),
		slack.MsgOptionUser(config.SlackUsername),
	)

	fmt.Println(channel)

	return err
}
