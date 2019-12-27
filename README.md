# webhook-fanout

webhook-fanout is designed to receive webhooks, and fan them out to multiple receivers.

## Why
Designed to add a webhook component to tools such as
[git-sync](https://github.com/kubernetes/git-sync), it supports a
plugin system to identify endpoints/receivers to forward hooks to.
The original use-case is for multiple k8s pods utilizing git-sync
sidecars to receive hooks from a git server indicating a commit event
has occurred, to either eliminate the polling behavior, or drastically
increase the amount of time between polls, lowering the load on
external git servers.

## TODO
Actually write the code to get k8s pod endpoints by labels.