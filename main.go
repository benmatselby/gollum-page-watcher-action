package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	if os.Getenv("GITHUB_EVENT_NAME") != "gollum" {
		fmt.Println("GITHUB_EVENT_NAME is not a 'gollum' event, so nothing to do")
		return
	}
	if os.Getenv("GITHUB_EVENT_PATH") == "" {
		fmt.Println("There is no GITHUB_EVENT_PATH defined, cannot carry on")
		os.Exit(1)
	}

	file, err := ioutil.ReadFile(os.Getenv("GITHUB_EVENT_PATH"))
	if err != nil {
		fmt.Println("Unable to read the file defined GITHUB_EVENT_PATH, cannot carry on")
		os.Exit(1)
	}

	var gollum GollumEvent
	if err := json.Unmarshal([]byte(file), &gollum); err != nil {
		fmt.Println("Unable to understand the JSON defined in GITHUB_EVENT_PATH, cannot carry on")
		os.Exit(1)
	}

	fmt.Println(gollum)
}
