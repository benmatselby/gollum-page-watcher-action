package notifier

import (
	"github.com/benmatselby/gollum-page-watcher-action/config"
	"github.com/benmatselby/gollum-page-watcher-action/github"
)

// NotificationStrategy defines the interface for all notifiers
type NotificationStrategy interface {
	Send(config.Config, *github.GollumEvent) error
}

// Notifier allows for a strategy to be executed
type Notifier struct {
	Strategy NotificationStrategy
}

// Send communicates the event via the NotificationStrategy
func (n *Notifier) Send(config config.Config, event *github.GollumEvent) error {
	return n.Strategy.Send(config, event)
}
