package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/slack-go/slack"
)

// Page defines the struct of each page that has changed during the action. For more detail see https://developer.github.com/v3/activity/events/types/#gollumevent
type Page struct {
	Name   string `json:"page_name"`
	Title  string `json:"title"`
	Action string `json:"action"`
	Sha    string `json:"sha"`
	URL    string `json:"html_url"`
}

// GollumEvent houses all of the Page structs. For more detail see https://developer.github.com/v3/activity/events/types/#gollumevent
type GollumEvent struct {
	Pages []Page `json:"pages"`
}

// main is the handler for the action
// Environment variables used are defined here: https://help.github.com/en/actions/configuring-and-managing-workflows/using-environment-variables#default-environment-variables
func main() {
	ok, msg := ValidateConfiguration()
	if !ok {
		fmt.Println(msg)
		os.Exit(1)
	}

	pages, err := GetPagesChanged()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = SendSlackMessage(pages)
	if err != nil {
		fmt.Println(msg)
		os.Exit(1)
	}
}

// ValidateConfiguration will validate all of the environment variables required to run the action
func ValidateConfiguration() (bool, string) {
	if os.Getenv("GITHUB_EVENT_NAME") != "gollum" {
		return false, "GITHUB_EVENT_NAME is not a 'gollum' event, so nothing to do."
	}

	if os.Getenv("GITHUB_EVENT_PATH") == "" {
		return false, "There is no GITHUB_EVENT_PATH defined, cannot carry on."
	}

	if os.Getenv("SLACK_WEBHOOK") == "" {
		return false, "There is no SLACK_WEBHOOK defined, therefore we could not post a message to slack."
	}

	return true, ""
}

// GetPagesChanged will use the configuration to determine if any of the "watched" pages
// have been edited or created
func GetPagesChanged() ([]Page, error) {
	file, err := ioutil.ReadFile(os.Getenv("GITHUB_EVENT_PATH"))
	if err != nil {
		return nil, fmt.Errorf("Unable to read the file defined GITHUB_EVENT_PATH, cannot carry on")
	}

	var gollum GollumEvent
	if err := json.Unmarshal([]byte(file), &gollum); err != nil {
		return nil, fmt.Errorf("Unable to understand the JSON defined in GITHUB_EVENT_PATH, cannot carry on")
	}

	return gollum.Pages, nil
}

// SendSlackMessage will send the required message to Slack.
func SendSlackMessage(pages []Page) error {
	attachments := []slack.Attachment{}
	for _, page := range pages {
		attachments = append(attachments, slack.Attachment{
			Color: "#2e5685",
			Text:  fmt.Sprintf("<%s|%s>", page.URL, page.Title),
		})
	}

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
