package main

import (
	"fmt"
	"os"

	"github.com/benmatselby/gollum-page-watcher-action/config"
	"github.com/benmatselby/gollum-page-watcher-action/github"
	"github.com/benmatselby/gollum-page-watcher-action/notifier"
)

// main is the handler for the action
// Environment variables used are defined here: https://help.github.com/en/actions/configuring-and-managing-workflows/using-environment-variables#default-environment-variables
func main() {
	config := config.Config{
		GitHubEventName: os.Getenv("GITHUB_EVENT_NAME"),
		GitHubEventPath: os.Getenv("GITHUB_EVENT_PATH"),
		SlackWebhook:    os.Getenv("SLACK_WEBHOOK"),
		SlackUsername:   os.Getenv("SLACK_USERNAME"),
		SlackChannel:    os.Getenv("SLACK_CHANNEL"),
		PagesToWatch:    os.Getenv("PAGES_TO_WATCH"),
		Debug:           os.Getenv("DEBUG"),
	}

	ok, msg := config.IsValid()
	if !ok {
		fmt.Println(msg)
		os.Exit(1)
	}

	event, err := github.GetGollumEvent(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(event.Pages) == 0 {
		fmt.Println("No pages being watched have been changed.")
		os.Exit(0)
	}

	commsStrategy := notifier.Notifier{Strategy: &notifier.Slack{}}
	err = commsStrategy.Strategy.Send(config, event)
	if err != nil {
		fmt.Println(msg)
		os.Exit(1)
	}
}
