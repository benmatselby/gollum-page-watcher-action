package notifier

import "github.com/benmatselby/gollum-page-watcher-action/github"

// NotificationStrategy defines the interface for all notifiers
type NotificationStrategy interface {
	Send(event *github.GollumEvent) error
}

// Notifier allows for a strategy to be executed
type Notifier struct {
	Strategy NotificationStrategy
}

// Send communicates the event via the NotificationStrategy
func (n *Notifier) Send(event *github.GollumEvent) error {
	return n.Strategy.Send(event)
}
