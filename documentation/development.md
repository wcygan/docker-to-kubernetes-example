# Local Development

[Tilt](https://tilt.dev/) is used for local development and testing.

```bash
tilt up
```

This command is captured in [tilt.sh](../scripts/tilt.sh).

## Port Forwarding

When you run `tilt up`, the tunnel is automatically created for you; no additional setup is required.

Port Forwarding is configured in [Tiltfile](../Tiltfile).

This means that `grpc-server` should be accessible on port `8080` from your local machine.
