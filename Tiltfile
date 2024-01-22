load('ext://restart_process', 'docker_build_with_restart')

# Compile command for the Go application
compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o grpc-server server/main.go'

# Compile command for the protobuf files
proto_compile_cmd = 'buf generate'

# Local resource to compile the Go application
local_resource(
  'compile',
  compile_cmd,
  deps=['./server', './generated'],
)

# Local resource to compile the protobuf files
local_resource(
  'proto_compile',
  proto_compile_cmd,
  deps=['./proto'],
)

# Docker build with restart for the Go application
docker_build_with_restart(
    'wcygan/docker-to-kubernetes-example-server',
    '.',
    dockerfile='server/Dockerfile',
    entrypoint='/app/grpc-server',
    live_update=[
        sync('./', '/app'),
    ]
)

# Kubernetes resources
k8s_yaml(['server/Deployment.yaml', 'server/service.yaml'])
k8s_resource('grpc-server', port_forwards=8080)