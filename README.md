# Gollum Page Watcher GitHub Action

![GitHub Badge](https://github.com/benmatselby/gollum-page-watcher-action/workflows/Go/badge.svg)

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
    WATCH_PAGES: 0.57.2
    TARGET_REPO: benmatselby/benmatselby.github.io
    TOKEN: ${{ secrets.TOKEN }}
    HUGO_ARGS: '-t academic'
    CNAME: benmatselby.github.io
...
```
