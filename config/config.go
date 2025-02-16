package config

// Config looks everything coming into the action
// See https://help.github.com/en/actions/configuring-and-managing-workflows/using-environment-variables#default-environment-variables
type Config struct {
	// GitHubEventName is the name of the event causing the action to fire
	GitHubEventName string

	// GitHubEventPath is the file path to where the payload JSON file is
	GitHubEventPath string

	// SlackToken allows us to bypass the need for a webhook
	SlackToken string

	// SlackWebhook is the destination of the Slack API call
	SlackWebhook string

	// SlackUsername is the username the message will come from
	SlackUsername string

	// SlackChannel is the destination of the Slack message
	SlackChannel string

	// PagesToWatch defines which page titles to watch. If not specified, all
	// pages will be watched, and notified on.
	PagesToWatch string

	// Debug defines if we are running in debug mode
	Debug string
}

// IsValid will provide insight unto whether we are in an acceptable state
func (c *Config) IsValid() (bool, string) {
	if c.GitHubEventName != "gollum" {
		return false, "GITHUB_EVENT_NAME is not a 'gollum' event, so nothing to do."
	}

	if c.GitHubEventPath == "" {
		return false, "There is no GITHUB_EVENT_PATH defined, cannot carry on."
	}

	if c.SlackToken == "" && c.SlackWebhook == "" {
		return false, "You need to provide either SLACK_WEBHOOK or SLACK_TOKEN, none provided therefore we could not post a message to slack."
	}

	return true, ""
}
