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
2. [Building Docker images and pushing them to Docker Hub using GitHub Actions](https://github.com/wcygan/github-actions-to-docker-registry)

If you care about how these work, you can go see the repositories for more information. I won't explain the mechanics of
these patterns in this example.

## Generating the Protobuf and gRPC Stubs

Run the following command:

```bash
$ buf generate proto
```

## Starting the Server using Go

Run the following command:

```bash
$ go run go-server/main.go
```

## Starting the Server using Docker

First, pull the Docker image:

```
$ docker pull wcygan/docker-to-kubernetes-example-server
```

Then, run the Docker image:

```
$ docker run -p 50051:50051 wcygan/docker-to-kubernetes-example-server
```

## Starting the Client

You can use the Go client to test the server:

```
$ go run client/main.go localhost:50051
```

## Using grpcurl to test the Server

Alternatively, you can use [grpcurl](https://github.com/fullstorydev/grpcurl) to test the server:

```
$ grpcurl -plaintext -proto proto/ping/v1/ping.proto -d '{"message": "Hello"}' localhost:50051 ping.v1.PingService/Ping

{
  "message": "Pong"
}
```

This command requires you to specify the proto file. 

## Build the Docker image for the Server

Run this in the root directory of this project:

```
docker build -t server -f server/Dockerfile .
```