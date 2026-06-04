# DEBUG

## Failure hypotheses

1. The image was built for the wrong CPU family, most likely `arm64`, while the Docker VM expects `amd64`.
2. The Go binary inside the image was compiled without fixing the target platform, so the build host influenced the artifact.
3. The image was valid, but the executable copied into it could not run on the remote machine.

## Hypothesis verification

```bash
docker image inspect ttl.sh/abhi-challenge3:2h --format '{{.Architecture}}'
file main
docker inspect ttl.sh/abhi-challenge3:2h
```

To verify the fix locally, rebuild with the target architecture pinned and inspect the image metadata again:

```bash
docker build --platform linux/amd64 -t ttl.sh/abhi-challenge3:2h .
docker image inspect ttl.sh/abhi-challenge3:2h --format '{{.Architecture}}'
```

## Solution

The fix is to make the Docker build explicitly target `linux/amd64` and to cross-compile the Go binary with `CGO_ENABLED=0`, `GOOS=linux`, and `GOARCH=amd64`. That keeps the release artifact aligned with the x86_64 Docker VM instead of whatever platform happened to build it.

## Underlying lesson

A successful build only matters if the artifact can actually execute on the host you deploy to.
