#!/bin/bash
CONTAINER_NAME=""
IMAGE_NAME=""
IMAGE_URL=""


# Drop old containers
docker stop $CONTAINER_NAME
docker rm $CONTAINER_NAME

# Pull latest image
# TODO

# Start new containers
docker run -p 80:80 -d -t --name $CONTAINER_NAME $IMAGE_NAME
