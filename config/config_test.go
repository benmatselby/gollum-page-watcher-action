package config

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tt := []struct {
		name         string
		eventName    string
		eventPath    string
		slackToken   string
		slackWebhook string
		expectedBool bool
		expectedMsg  string
	}{
		{name: "all env vars defined", eventName: "gollum", eventPath: "/tmp/payload.json", slackWebhook: "https://api.slack.com", expectedBool: true, expectedMsg: ""},
		{name: "event name is not gollum", eventName: "push", eventPath: "/tmp/payload.json", slackWebhook: "https://api.slack.com", expectedBool: false, expectedMsg: "GITHUB_EVENT_NAME is not a 'gollum' event, so nothing to do."},
		{name: "event name is missing", eventName: "", eventPath: "/tmp/payload.json", slackWebhook: "https://api.slack.com", expectedBool: false, expectedMsg: "GITHUB_EVENT_NAME is not a 'gollum' event, so nothing to do."},
		{name: "event path is missing", eventName: "gollum", eventPath: "", slackWebhook: "https://api.slack.com", expectedBool: false, expectedMsg: "There is no GITHUB_EVENT_PATH defined, cannot carry on."},
		{name: "slack webhook and token are missing", eventName: "gollum", eventPath: "/tmp/payload.json", slackWebhook: "", slackToken: "", expectedBool: false, expectedMsg: "You need to provide either SLACK_WEBHOOK or SLACK_TOKEN, none provided therefore we could not post a message to slack."},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := Config{
				GitHubEventName: tc.eventName,
				GitHubEventPath: tc.eventPath,
				SlackWebhook:    tc.slackWebhook,
			}
			ok, msg := c.IsValid()

			if ok != tc.expectedBool {
				t.Fatalf("expected return boolean to be '%v'; got '%v'", tc.expectedBool, ok)
			}

			if msg != tc.expectedMsg {
				t.Fatalf("expected return message to be '%s'; got '%s'", tc.expectedMsg, msg)
			}
		})
	}
}
