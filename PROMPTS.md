# PROMPTS

For Challenge 3, I wanted the Go service to be packaged as a proper Docker image, pushed to `ttl.sh`, and deployed from Jenkins onto the remote Docker VM on port `4444`.

I chose a multi-stage Dockerfile with a static build so the final image stays small and predictable. I also pinned the Jenkins build to `linux/amd64` because the VM that runs the verification is x86_64, and I did not want the host machine’s architecture to leak into the deployment.

I pushed back on a single-stage image and on any deploy flow that assumed the Docker VM was available locally from Jenkins. The image still needs to be pulled on the target host, and the container has to be started there with the correct port mapping.
