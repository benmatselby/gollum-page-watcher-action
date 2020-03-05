package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Page defines the struct of each page that has changed during the action
type Page struct {
	Name   string `json:"page_name"`
	Title  string `json:"title"`
	Action string `json:"action"`
	Sha    string `json:"sha"`
	URL    string `json:"html_url"`
}

// GollumEvent houses all of the Page structs
type GollumEvent struct {
	Pages []Page `json:"pages"`
}

func main() {
	if os.Getenv("GITHUB_EVENT_PATH") == "" {
		fmt.Println("There is no GITHUB_EVENT_PATH defined, cannot carry on")
		os.Exit(1)
	}

	fmt.Println(os.Getenv("GITHUB_EVENT_PATH"))

	file, err := ioutil.ReadFile(os.Getenv("GITHUB_EVENT_PATH"))
	if err != nil {
		fmt.Println("Uanble to read the file defined GITHUB_EVENT_PATH, cannot carry on")
		os.Exit(1)
	}

	var gollum GollumEvent
	if err := json.Unmarshal([]byte(file), &gollum); err != nil {
		fmt.Println("Uanble to understand the JSON defined in GITHUB_EVENT_PATH, cannot carry on")
		os.Exit(1)
	}

	fmt.Println(gollum)
}
