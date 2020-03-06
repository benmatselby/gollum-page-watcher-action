package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/benmatselby/gollum-page-watcher-action/github"
	"github.com/benmatselby/gollum-page-watcher-action/notifier"
)

// main is the handler for the action
// Environment variables used are defined here: https://help.github.com/en/actions/configuring-and-managing-workflows/using-environment-variables#default-environment-variables
func main() {
	ok, msg := ValidateConfiguration()
	if !ok {
		fmt.Println(msg)
		os.Exit(1)
	}

	event, err := GetGollumEvent()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	commsStrategy := notifier.Notifier{Strategy: &notifier.Slack{}}
	err = commsStrategy.Strategy.Send(event)
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

// GetGollumEvent will unmarshal the JSON we receive from GitHub
func GetGollumEvent() (*github.GollumEvent, error) {
	file, err := ioutil.ReadFile(os.Getenv("GITHUB_EVENT_PATH"))
	if err != nil {
		return nil, fmt.Errorf("Unable to read the file defined GITHUB_EVENT_PATH, cannot carry on")
	}

	var gollum github.GollumEvent
	if err := json.Unmarshal([]byte(file), &gollum); err != nil {
		return nil, fmt.Errorf("Unable to understand the JSON defined in GITHUB_EVENT_PATH, cannot carry on")
	}

	return &gollum, nil
}
