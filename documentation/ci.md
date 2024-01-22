# Continuous Integration

Stuff that happens after a commit is pushed to GitHub...

## Checking Tilt with CI

https://docs.tilt.dev/ci.html explains how to run Tilt in CI to ensure that your Tiltfile is valid & your local
development environment doesn't break

## Building & Pushing images to Docker Hub

Following the examples in https://github.com/wcygan/github-actions-to-docker-registry, I've
setup [publish-docker-image.yml](../.github/workflows/publish-docker-image.yml) to push images for this application to
Docker Hub.