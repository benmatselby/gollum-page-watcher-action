package main

import (
	"os"
	"testing"
)

func TestValidateConfiguration(t *testing.T) {
	tt := []struct {
		name         string
		eventName    string
		eventPath    string
		slackWebhook string
		expectedBool bool
		expectedMsg  string
	}{
		{name: "all env vars defined", eventName: "gollum", eventPath: "/tmp/payload.json", slackWebhook: "https://api.slack.com", expectedBool: true, expectedMsg: ""},
		{name: "event name is not gollum", eventName: "push", eventPath: "/tmp/payload.json", slackWebhook: "https://api.slack.com", expectedBool: false, expectedMsg: "GITHUB_EVENT_NAME is not a 'gollum' event, so nothing to do."},
		{name: "event name is missing", eventName: "", eventPath: "/tmp/payload.json", slackWebhook: "https://api.slack.com", expectedBool: false, expectedMsg: "GITHUB_EVENT_NAME is not a 'gollum' event, so nothing to do."},
		{name: "event path is missing", eventName: "gollum", eventPath: "", slackWebhook: "https://api.slack.com", expectedBool: false, expectedMsg: "There is no GITHUB_EVENT_PATH defined, cannot carry on."},
		{name: "slack webhook is missing", eventName: "gollum", eventPath: "/tmp/payload.json", slackWebhook: "", expectedBool: false, expectedMsg: "There is no SLACK_WEBHOOK defined, therefore we could not post a message to slack."},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			os.Setenv("GITHUB_EVENT_NAME", tc.eventName)
			os.Setenv("GITHUB_EVENT_PATH", tc.eventPath)
			os.Setenv("SLACK_WEBHOOK", tc.slackWebhook)

			ok, msg := ValidateConfiguration()

			if ok != tc.expectedBool {
				t.Fatalf("expected return boolean to be '%v'; got '%v'", tc.expectedBool, ok)
			}

			if msg != tc.expectedMsg {
				t.Fatalf("expected return message to be '%s'; got '%s'", tc.expectedMsg, msg)
			}
		})
	}
}
