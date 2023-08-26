package github

import (
	"reflect"
	"testing"

	"github.com/benmatselby/gollum-page-watcher-action/config"
)

func TestGetGollumEvent(t *testing.T) {
	c := config.Config{
		GitHubEventPath: "../example-payloads/valid-payload.json",
	}

	event, err := GetGollumEvent(c)

	if err != nil {
		t.Fatalf("Did not expect an error, got %s", err)
	}

	if len(event.Pages) != 3 {
		t.Fatalf("Expected to find two pages, got %v", len(event.Pages))
	}
}

func TestInvalidPayloadForGetGollumEvent(t *testing.T) {
	c := config.Config{
		GitHubEventPath: "../example-payloads/badly-formed-payload.json",
	}

	_, err := GetGollumEvent(c)

	if err == nil {
		t.Fatal("Expected error and did not find one")
	}

	if err.Error() != "unable to understand the JSON defined in GITHUB_EVENT_PATH, cannot carry on" {
		t.Fatal("Expected an error")
	}
}

func TestMissingPayloadForGetGollumEvent(t *testing.T) {
	c := config.Config{
		GitHubEventPath: "../example-payloads/unknown-payload.json",
	}

	_, err := GetGollumEvent(c)

	if err == nil {
		t.Fatal("Expected error and did not find one")
	}

	if err.Error() != "unable to read the file defined GITHUB_EVENT_PATH, cannot carry on" {
		t.Fatal("Expected an error")
	}
}

func TestGetGollumEventCanFilterPages(t *testing.T) {
	tt := []struct {
		name          string
		pagesToWatch  string
		expectedPages []string
	}{
		{name: "No filter", pagesToWatch: "", expectedPages: []string{"The homepage", "How to test", "How to test actions"}},
		{name: "Filter on Home", pagesToWatch: "home", expectedPages: []string{"The homepage"}},
		{name: "Filter on Test", pagesToWatch: "test", expectedPages: []string{"How to test", "How to test actions"}},
		{name: "Home or Test", pagesToWatch: "(home)|(test)", expectedPages: []string{"The homepage", "How to test", "How to test actions"}},
		{name: "No match", pagesToWatch: "Nada", expectedPages: []string{}},
		{name: "Exact page name match", pagesToWatch: "^How to test$", expectedPages: []string{"How to test"}},
		{name: "Multiple exact page name matches", pagesToWatch: "^The homepage$|^How to test$", expectedPages: []string{"The homepage", "How to test"}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := config.Config{
				GitHubEventPath: "../example-payloads/valid-payload.json",
				PagesToWatch:    tc.pagesToWatch,
			}

			event, err := GetGollumEvent(c)

			if err != nil {
				t.Fatalf("Did not expect an error, got %s", err)
			}

			pages := []string{}
			for _, page := range event.Pages {
				pages = append(pages, page.Title)
			}

			if !reflect.DeepEqual(pages, tc.expectedPages) {
				t.Fatalf("Expected pages %v, got %v", tc.expectedPages, pages)
			}
		})
	}
}
