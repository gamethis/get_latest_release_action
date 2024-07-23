# Get latest release docker action

[![Run pre-commit](https://github.com/gamethis/get_latest_release_action/actions/workflows/pre-commit.yaml/badge.svg)](https://github.com/gamethis/get_latest_release_action/actions/workflows/pre-commit.yaml)

This action gets the laest GitHub release.

## Inputs

## `repo`

**Required** The Repository you want to get the latest version from.  Should be specified in org/repo fasion.
IE.. `gamethis/get_latest_release_action`

## `major_version`

**Optional** The major version you want to constrain the release version to.

## Outputs

## `version`

The latest release version of the repo targeted.

## Example usage

- Get latest version with no constraint
```yaml
uses: gamethis/get_latest_release_action@v1
with:
  repo: get_latest_release_action
```

- Get latest version with `major_version` contraint
 ```yaml
uses: gamethis/get_latest_release_action@v1
with:
  repo: get_latest_release_action
  major_version: 1
``` 
