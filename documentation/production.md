# Production Deployment

The application is deployed on Kubernetes using images from Docker Hub.

For an example of how to automatically push images to Docker Hub,
see [building Docker images and pushing them to Docker Hub using GitHub Actions](https://github.com/wcygan/github-actions-to-docker-registry).

## Deploy

Deploy the application to your target environment by running this in the root directory of this project:

```
./scripts/deploy.sh
```

## Undeploy

Undeploy the application from your target environment by running this in the root directory of this project:

```
./scripts/undeploy.sh
```

## Port Forwarding

When deploying to your target environment, you can
use [kubectl port-forward](https://kubernetes.io/docs/tasks/access-application-cluster/port-forward-access-application-cluster/#forward-a-local-port-to-a-port-on-the-pod)
to access the cluster from your local machine. This repository has a script, [port-forward.sh](../scripts/port-forward.sh), which captures how to expose `grpc-server` on port `8080` locally.

Run this script to enable port forwarding:

```
./scripts/port-forward.sh
```

## Testing

After following [deploy](#deploy) and [port forwarding](#port-forwarding), we can run [the client](testing.md#go-client) in another terminal window:

```
go run client/main.go ping --ip localhost --port 8080
```