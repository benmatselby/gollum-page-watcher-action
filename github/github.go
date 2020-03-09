package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/benmatselby/gollum-page-watcher-action/config"
)

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

// Repo represents the repo the wiki is linked to
type Repo struct {
	FullName string `json:"full_name"`
	URL      string `json:"html_url"`
}

// GollumEvent houses all of the Page structs. For more detail see https://developer.github.com/v3/activity/events/types/#gollumevent
type GollumEvent struct {
	Pages  []Page `json:"pages"`
	Sender Sender `json:"sender"`
	Repo   Repo   `json:"repository"`
}

// GetGollumEvent will unmarshal the JSON we receive from GitHub
func GetGollumEvent(config config.Config) (*GollumEvent, error) {
	file, err := ioutil.ReadFile(config.GitHubEventPath)
	if err != nil {
		return nil, fmt.Errorf("Unable to read the file defined GITHUB_EVENT_PATH, cannot carry on")
	}

	var gollum GollumEvent
	if err := json.Unmarshal([]byte(file), &gollum); err != nil {
		return nil, fmt.Errorf("Unable to understand the JSON defined in GITHUB_EVENT_PATH, cannot carry on")
	}

	var watching = regexp.MustCompile(config.PagesToWatch)
	var pages []Page
	for _, page := range gollum.Pages {
		if watching.MatchString(page.Title) {
			pages = append(pages, page)
		}
	}

	gollum.Pages = pages

	return &gollum, nil
}
