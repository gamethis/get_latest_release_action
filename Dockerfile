# Container image that runs your code
FROM golang:alpine3.20


# Copies your code file from your action repository to the filesystem path `/` of the container
COPY main.go /main.go
COPY entrypoint.sh /entrypoint.sh

# Code file to execute when the docker container starts up (`entrypoint.sh`)
ENTRYPOINT ["/entrypoint.sh"]
