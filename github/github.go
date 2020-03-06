package github

// Page defines the struct of each page that has changed during the action. For more detail see https://developer.github.com/v3/activity/events/types/#gollumevent
type Page struct {
	Name   string `json:"page_name"`
	Title  string `json:"title"`
	Action string `json:"action"`
	Sha    string `json:"sha"`
	URL    string `json:"html_url"`
}

// Sender represents the author of the commit
type Sender struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
}

// GollumEvent houses all of the Page structs. For more detail see https://developer.github.com/v3/activity/events/types/#gollumevent
type GollumEvent struct {
	Pages  []Page `json:"pages"`
	Sender Sender `json:"sender"`
}
