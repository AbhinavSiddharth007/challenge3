pipeline {
    agent any

    environment {
        IMAGE_NAME     = "ttl.sh/abhi-challenge3:2h"
        CONTAINER_NAME = "challenge3"
        DEPLOY_HOST    = "docker"
    }

    stages {
        stage('Build image') {
            steps {
                sh '''
                    set -eu
                    docker build --platform linux/amd64 -t "$IMAGE_NAME" .
                '''
            }
        }
        stage('Push image') {
            steps {
                sh '''
                    set -eu
                    docker push "$IMAGE_NAME"
                '''
            }
        }
        stage('Deploy on Docker VM') {
            steps {
                sh '''
                    set -eu
                    ssh -o StrictHostKeyChecking=no -o ConnectTimeout=10 "root@$DEPLOY_HOST" "
                        docker pull $IMAGE_NAME &&
                        docker rm -f $CONTAINER_NAME 2>/dev/null || true &&
                        docker run -d --name $CONTAINER_NAME -p 4444:4444 $IMAGE_NAME
                    "
                '''
            }
        }
        stage('Test deployment') {
            steps {
                sh '''
                    set -eu
                    sleep 3
                    ssh -o StrictHostKeyChecking=no "root@$DEPLOY_HOST" \
                        "docker run --rm --network host busybox wget -qO- http://localhost:4444/" | grep -q Hello
                '''
            }
        }
    }
}
