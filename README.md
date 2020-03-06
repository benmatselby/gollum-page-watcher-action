# Gollum Page Watcher GitHub Action

<a href="https://github.com/benmatselby/gollum-page-watcher-action/actions"><img alt="status" src="https://github.com/benmatselby/gollum-page-watcher-action/workflows/Go/badge.svg"></a>

This GitHub action will watch for certain pages to change in the wiki, and then notify in a Slack channel.

## Secrets

- `SLACK_WEBHOOK`: The webhook to use to send the Slack notification.

## Environment Variables

- `WATCH_PAGES`: The pages we should "watch" for changes to.

## Example

```shell
...
- name: Wiki Watcher
  uses: benmatselby/gollum-page-watcher-action@master
  env:
    WATCH_PAGES:
    SLACK_WEBHOOK: https://hooks.slack.com/services/etc/etc
    SLACK_CHANNEL: #random
    SLACK_USERNAME: Gollum
...
```

## Testing

To test this, you can run it from your command line with the following setup

```shell
GITHUB_EVENT_PATH=gollum-event-payload.json \
GITHUB_EVENT_NAME=gollum \
SLACK_WEBHOOK=[your-slack-webook-url] \
SLACK_CHANNEL=[your-slack-channel] \
DEBUG=true \
go run main.go
```

If `DEBUG` is defined, it will not post to Slack, but rather output the webhook message in your terminal.

![Gollum](./img/gollum.jpg)
