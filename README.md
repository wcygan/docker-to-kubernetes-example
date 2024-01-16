## Docker to Kubernetes

This example shows how to deploy Docker images to Kubernetes.

### Prerequisites

Install the following tools:

| Tool     | Installation Instructions                 |
|----------|-------------------------------------------|
| `go`     | https://golang.org/doc/install            |
| `protoc` | https://grpc.io/docs/protoc-installation/ |
| `buf`    | https://buf.build/docs/installation       |
| `docker` | https://docs.docker.com/get-docker/       |


## Generating the Protobuf and gRPC Stubs

Run the following command:

```bash
$ buf generate proto
```