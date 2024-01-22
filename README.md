## Docker to Kubernetes

This example shows how to deploy Docker images to Kubernetes.

## Prerequisites

### Kubernetes Cluster

You'll need a [Kubernetes](https://kubernetes.io/) cluster to run your applications on.
See [Choosing a Local Cluster](https://docs.tilt.dev/choosing_clusters) if you're just getting started, or use something
like [k3s](https://k3s.io/) if you have your own hardware.

For local development, I'm using [Minikube](https://minikube.sigs.k8s.io/docs/) (to provide a cluster)
and [Tilt](https://docs.tilt.dev/) (to automate the development workflow).

### Tooling

Install the following tools:

| Tool       | Installation Instructions                                               | Purpose                       |
|------------|-------------------------------------------------------------------------|-------------------------------|
| `go`       | https://golang.org/doc/install                                          | Compile Go code               |
| `buf`      | https://buf.build/docs/installation                                     | Compile Protobuf files        |
| `grpcurl`  | https://github.com/fullstorydev/grpcurl?tab=readme-ov-file#installation | Test gRPC services            |
| `tilt`     | https://docs.tilt.dev/install.html                                      | Automate development workflow |
| `minikube` | https://minikube.sigs.k8s.io/docs/start/                                | Run Kubernetes locally        |

## Quick Start

With the [prerequisites](#prerequisites) installed, you can run the application:

```bash
minikube start && tilt up
```

In a separate terminal, you can use [grpcurl](https://github.com/fullstorydev/grpcurl) to test the application:

```bash
grpcurl -plaintext -proto proto/ping/v1/ping.proto localhost:8080 ping.v1.PingService/Ping
```

You should see the following output:

```
Response: Pong
```

## Documentation

- [Development](documentation/development.md)
- [Production](documentation/production.md)
- [Multiple Environments](documentation/multi-environment.md)
- [Testing](documentation/testing.md)
- [Proto](documentation/proto.md)
- [Continuous Integration](documentation/ci.md)