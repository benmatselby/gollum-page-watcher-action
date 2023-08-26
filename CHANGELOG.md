# CHANGELOG

## next

- Bumped Go version to Go 1.21.

## 1.8.0

- Bumped docker image to Go 1.20 runtime.
- Moved the default branch to `main`, and left `master` running the v1.7.0 version.

## 1.7.0

- Bumped docker image to Go 1.18 runtime.

## 1.6.0

- Bumped docker image to Go 1.17 runtime.

## 1.5.0

- Bumped docker image to Go 1.16 runtime.
- Bump the build environment to test on 1.16.
- Bumped `github.com/slack-go/slack` from 0.6.2 to 0.8.1.

## 1.4.0

- [Feature #19](https://github.com/benmatselby/gollum-page-watcher-action/pull/19) - Update the slack message to be similar to other application output in Slack. Thanks to [@markgaze](https://github.com/markgaze) for the contribution.

## 1.3.0

- [Resolves #16](https://github.com/benmatselby/gollum-page-watcher-action/issues/16) - Updated Slack output, to explain which repo the wiki has changed in.

## 1.2.0

- Package refactor to make it easier to build new "notifiers"
- Provide new variable `PAGES_TO_WATCH` that allows users to define a regex of page titles to watch and be notified when changed.

## 1.1.0

- Show which GitHub user has made the change to the wiki.

## 1.0.0

- Listens to the `gollum` event and posts messages to Slack.
