module github.com/wcygan/docker-to-kubernetes-example/server

go 1.21

require (
	github.com/go-redis/redis/v8 v8.11.5
	github.com/wcygan/docker-to-kubernetes-example/generated/go v0.0.0
	google.golang.org/grpc v1.60.1
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.16.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231002182017-d307bd883b97 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace github.com/wcygan/docker-to-kubernetes-example/generated/go => ../generated/go
