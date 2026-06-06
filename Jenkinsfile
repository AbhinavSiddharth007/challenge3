pipeline {
    agent any

    environment {
        // Unique image name (ttl.sh is public). Keep the :2h tag.
        IMAGE          = "ttl.sh/abhi-challenge3:2h"
        CONTAINER_NAME = "challenge3"
    }

    stages {
        stage('Build & Push') {
            steps {
                sh '''
                    set -eu
                    docker build --platform linux/amd64 -t "$IMAGE" .
                    docker push "$IMAGE"
                '''
            }
        }

        stage('Deploy on Docker VM') {
            // Runs on the docker VM. 'docker' is the usual node label for these
            // labs (matches the docker:4444 host in the task list). If running
            // this errors with "no node with label docker", check
            // Manage Jenkins -> Nodes for the real label and change it here.
            agent { label 'docker' }
            options { skipDefaultCheckout(true) }
            steps {
                sh '''
                    set -eu
                    docker pull "$IMAGE"
                    docker rm -f "$CONTAINER_NAME" >/dev/null 2>&1 || true
                    docker run -d --name "$CONTAINER_NAME" -p 4444:4444 "$IMAGE"
                '''
            }
        }
    }
}
