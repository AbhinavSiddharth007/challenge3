# Reflection — Docker build and deploy

## What did I do?

I turned the Go service into a containerized delivery flow: it now builds inside a multi-stage Dockerfile, gets pushed to `ttl.sh`, and is started on the remote Docker VM through Jenkins on port `4444`.

## What was most surprising?

The most surprising part was how easy it is for the same source code to fail at runtime if the image was built for the wrong architecture. Treating `linux/amd64` as part of the release contract, not just a build detail, made the deployment much more reliable.

## What’s still unclear?

I still want to get more comfortable with the trade-offs between a plain `docker build` and a `docker buildx build` workflow, especially when the target machine architecture is different from the machine that runs Jenkins.
