# Multiple Environments

I run two clusters:

1. [minikube](https://minikube.sigs.k8s.io/docs/start/): a single-node cluster that runs on my laptop
2. [k3s](https://k3s.io/): a multi-node cluster that runs in my homelab

This is similar to having a [development](development.md) and [production](production.md) environment.

## Context Switching (or, targeting different clusters with `kubectl`)

[kubectx](https://github.com/ahmetb/kubectx) is a tool that makes it easy to switch between clusters.

After setting up my clusters, I've added their kubeconfig files to the `~/.kube/configs` directory:

```
ls ~/.kube/configs
k3s.yaml  minikube.yaml
```

### Switching between clusters

Switch to the `k3s` cluster:

```
kubectx k3s
Switched to context "k3s".
```

Switch to the `minikube` cluster:

```
kubectx minikube
Switched to context "minikube".
```

### Deployment

After specifying the correct cluster, I can manage the deployment using various [scripts](../scripts):

| Script                                        | Description                                  |
|-----------------------------------------------|----------------------------------------------|
| [deploy.sh](../scripts/deploy.sh)             | Deploys the application to the cluster       |
| [undeploy.sh](../scripts/undeploy.sh)         | Deletes the application from the cluster     |
| [port-forward.sh](../scripts/port-forward.sh) | Forwards the application's port to localhost |
| [ping.sh](../scripts/ping.sh)                 | Pings the application                        |

### A note on `kubectx`

I'm not entirely sure how it's working yet because I haven't changed my `~/.kube/config` file, I've only added new files
to the `~/.kube/configs` directory. Yet, `kubectx` is able to switch between them; there has to be something I'm
missing...