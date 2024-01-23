load('ext://restart_process', 'docker_build_with_restart')

# Compile command for the Go application
compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../build/grpc-server .'

# Compile command for the protobuf files
proto_compile_cmd = 'buf generate'

# Local resource to compile the Go application
local_resource(
  'compile',
  compile_cmd,
  deps=['./server', './generated'],
  dir='./server'  # Set the working directory for the command
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
    entrypoint='/app/build/grpc-server',
    ignore=['./client'],
    live_update=[
        sync('./', '/app'),
    ],
)

# Kubernetes resources
k8s_yaml(['server/Deployment.yaml', 'server/service.yaml', 'storage/redis-deployment.yaml', 'storage/redis-service.yaml'])
k8s_resource('redis-deployment', new_name='redis')
k8s_resource('grpc-server', port_forwards=8080, resource_deps=['redis'])