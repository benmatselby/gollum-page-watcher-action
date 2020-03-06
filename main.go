package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		Debug:           os.Getenv("DEBUG"),
	}

	ok, msg := config.IsValid()
	if !ok {
		fmt.Println(msg)
		os.Exit(1)
	}

	event, err := GetGollumEvent(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	commsStrategy := notifier.Notifier{Strategy: &notifier.Slack{}}
	err = commsStrategy.Strategy.Send(config, event)
	if err != nil {
		fmt.Println(msg)
		os.Exit(1)
	}
}

// GetGollumEvent will unmarshal the JSON we receive from GitHub
func GetGollumEvent(config config.Config) (*github.GollumEvent, error) {
	file, err := ioutil.ReadFile(config.GitHubEventPath)
	if err != nil {
		return nil, fmt.Errorf("Unable to read the file defined GITHUB_EVENT_PATH, cannot carry on")
	}

	var gollum github.GollumEvent
	if err := json.Unmarshal([]byte(file), &gollum); err != nil {
		return nil, fmt.Errorf("Unable to understand the JSON defined in GITHUB_EVENT_PATH, cannot carry on")
	}

	return &gollum, nil
}
