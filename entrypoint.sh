#!/bin/sh -l

REPO=$1
DESIRED_VERSION=$2

which go
go run main.go -repo_name=${REPO} -major=${DESIRED_VERSION}  >> "$GITHUB_OUTPUT"
