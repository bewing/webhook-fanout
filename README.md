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

## Usage
Currently this container fans out webhooks to Pods via selector labels.  The following arguments are supported:
* `-namespace=` - Namespace to watch pods in.  Defaults to all namespaces.
* `-selector=` - Label selector string to identify pods.  Defaults to all pods.
* `-targetPort=` - Port to use when fanning hooks out to target pods.  Defaults to 80.

The reference `Dockerfile` is published as `bewing/webhook-fanout:latest`.  Usage of this image is guaranteed to break at some point in the future as argument parsing and CI/CD improves.  If you're using in production, I'd recommend pinning the hash.

## TODO
* Improve argument parsing
* Write an operator and CRDs to allow the cluster to handle fanouts cluster-wide.
* CI/CD
