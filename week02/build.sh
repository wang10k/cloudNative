#!/bin/bash
cd /data/projects/cloudNative/cloudNative/week02
GOOS=linux GOARCH=amd64 go build .
docker build . -t timeapy/week02
# need login to dockerhub
# docker login
# start container.
# docker run -d --name week02 -p8001:80 week02

# push image to docker  hub
# docker push