# Docker Commands

All the information was from this [blog post](https://www.docker.com/blog/developing-go-apps-docker/)

## This builds the image:
docker build --rm -t [image-name]:alpha .
--rm removes intermediate containers created by each layer of the build process
-t Tags the image

## This lists the current images:
docker image ls

## This runs the image:
docker run -d -p 5173:8080 --name go-docker-app go-todo-app:latest

-d means dettached, so the app runs in the background
-p assigns Ports -> HOSTPORT:CONTAINERPORT
--name gives a name to the container
and then you choose the name of the image
