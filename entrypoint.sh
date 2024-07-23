#!/bin/sh -l

REPO=$1
DESIRED_VERSION=$2

function getLatestRepoVersion() {
  LATEST_ARR=($(wget -q -O- https://api.github.com/repos/${REPO}/releases 2> /dev/null | awk '/tag_name/ { print $2 }' | cut -d '"' -f 2 | cut -d 'v' -f 2 | sort -V -r))
  for ver in "${LATEST_ARR[@]}"; do
    if [[ -n "${DESIRED_VERSION}" ]]; then
      if [[ $ver =~ $DESIRED_VERSION ]] && [[ ! $ver =~ beta ]] && [[ ! $ver =~ rc ]] && [[ ! $ver =~ alpha ]] ; then
        LATEST="$ver"
        break
      fi
    else
      if [[ ! $ver =~ beta ]] && [[ ! $ver =~ rc ]] && [[ ! $ver =~ alpha ]] ; then
        LATEST="$ver"
        break
      fi
    fi
  done
  echo -n "$LATEST"
}

VERSION=$(getLatestRepoVersion() "${REPO}" "${DESIRED_VERSION}")

echo "version=${VERSION}" >> $GITHUB_OUTPUT
