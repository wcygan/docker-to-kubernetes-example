## Docker to Kubernetes

This example shows how to deploy Docker images to Kubernetes.

The Server's Dockerfile is uploaded
at [wcygan/docker-to-kubernetes-example-server](https://hub.docker.com/r/wcygan/docker-to-kubernetes-example-server) on
Docker Hub via GitHub Actions defined in [publish-docker-image.yml](.github/workflows/publish-docker-image.yml).

### Prerequisites

Install the following tools:

| Tool      | Installation Instructions                                               |
|-----------|-------------------------------------------------------------------------|
| `go`      | https://golang.org/doc/install                                          |
| `protoc`  | https://grpc.io/docs/protoc-installation/                               |
| `buf`     | https://buf.build/docs/installation                                     |
| `docker`  | https://docs.docker.com/get-docker/                                     |
| `grpcurl` | https://github.com/fullstorydev/grpcurl?tab=readme-ov-file#installation |

#### Assumptions

This project uses two patterns that I've defined in other repositories, namely:

1. [Sharing Proto schemas across projects in a monorepo](https://github.com/wcygan/buf-polyglot-example)
2. [Building Docker images and pushing them to Docker Hub using GitHub Actions](https://github.com/wcygan/github-actions-to-docker-registry.)

If you care about how these work, you can go see the repositories for more information. I won't explain the mechanics of
these patterns in this example.

## Generating the Protobuf and gRPC Stubs

Run the following command:

```bash
$ buf generate proto
```

## Starting the Server

Run the following command:

```bash
$ go run go-server/main.go
```

## Testing with grpcurl

You can use [grpcurl](https://github.com/fullstorydev/grpcurl) to test the server:

```
$ grpcurl -plaintext -proto proto/ping/v1/ping.proto -d '{"message": "Hello"}' localhost:50051 ping.v1.PingService/Ping

{
  "message": "Pong"
}
```

## Build the Docker image for the Server

Run this in the root directory of this project:

```
docker build -t server -f server/Dockerfile .
```