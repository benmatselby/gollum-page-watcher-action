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
...
```
