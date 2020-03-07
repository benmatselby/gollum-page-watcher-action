package github

import (
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

	if len(event.Pages) != 2 {
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

	if err.Error() != "Unable to understand the JSON defined in GITHUB_EVENT_PATH, cannot carry on" {
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

	if err.Error() != "Unable to read the file defined GITHUB_EVENT_PATH, cannot carry on" {
		t.Fatal("Expected an error")
	}
}
