# CHANGELOG

## next

## 1.3.0

- [Resolves #16](https://github.com/benmatselby/gollum-page-watcher-action/issues/16) - Updated Slack output, to explain which repo the wiki has changed in.

## 1.2.0

- Package refactor to make it easier to build new "notifiers"
- Provide new variable `PAGES_TO_WATCH` that allows users to define a regex of page titles to watch and be notified when changed.

## 1.1.0

- Show which GitHub user has made the change to the wiki.

## 1.0.0

- Listens to the `gollum` event and posts messages to Slack.
