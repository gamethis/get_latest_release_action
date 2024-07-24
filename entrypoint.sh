#!/bin/sh

REPO=$1
DESIRED_VERSION=$2

if [[ -z "$GITHUB_OUTPUT" ]]; then
  go run /main.go -repo_name=${REPO} -major=${DESIRED_VERSION}
else
  go run /main.go -repo_name=${REPO} -major=${DESIRED_VERSION} >> "$GITHUB_OUTPUT"
fi
  
