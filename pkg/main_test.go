package main

import (
	"testing"
)

func TestGetMajor(t *testing.T) {
	var release Release
	release.TagName = "v1.2.3"

	expected := 1
	actual := release.GetMajor()

	if actual != expected {
		t.Errorf("Expected major version %d, but got %d", expected, actual)
	}
}
func TestFilterReleases(t *testing.T) {
	var RepoName *string = new(string)
	*RepoName = "owner/repo"
	var Major *string = new(string)
	*Major = "1"
	var repo = Repo{
		RepoName: RepoName,
		Major:    Major,
	}

	releases := []Release{
		{
			TagName:    "v1.0.0",
			Draft:      false,
			Prerelease: false,
		},
		{
			TagName:    "v1.1.0",
			Draft:      false,
			Prerelease: false,
		},
		{
			TagName:    "v2.0.0",
			Draft:      false,
			Prerelease: false,
		},
		{
			TagName:    "v2.1.0",
			Draft:      false,
			Prerelease: false,
		},
	}

	filteredReleases := repo.filterReleases(releases)

	expected := 2
	actual := len(filteredReleases)

	if actual != expected {
		t.Errorf("Expected %d filtered releases, but got %d", expected, actual)
	}
}
