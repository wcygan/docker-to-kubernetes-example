docker_build('wcygan/docker-to-kubernetes-example-server', '.',
    dockerfile='server/Dockerfile')
k8s_yaml(['server/Deployment.yaml', 'server/service.yaml'])
k8s_resource('grpc-server', port_forwards=8080)