# Testing

Once the primary application is running (either following [Development](development.md) or [Production](production.md)),
you can test it using a client application.

## Clients

### Go Client

After following [proto.md](proto.md) to generate Go code from the proto files, you can use
the [Go client](../client/main.go) to test the server:

```
go run client/main.go ping --ip localhost --port 8080
```

### grpcurl

Alternatively, you can use [grpcurl](https://github.com/fullstorydev/grpcurl) to test the server:

```
grpcurl -plaintext -proto proto/ping/v1/ping.proto -d '{"message": "Hello"}' localhost:8080 ping.v1.PingService/Ping

{
  "message": "Pong"
}
```

This command requires you to specify the proto file.

## Port Forwarding

Typically, your Kubernetes cluster will not be accessible from your local machine. You can use port forwarding to access
the cluster from your local machine, which enables you to use the clients to interact with the application.

## Local Development with Tilt

See [Development Port Forwarding](development.md#port-forwarding).

## Production Deployment

See [Production Port Forwarding](production.md#port-forwarding).

