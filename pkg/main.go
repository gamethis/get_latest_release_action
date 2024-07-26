package main

import (
	"cmp"
	json "encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

var resp *http.Response
var err error
var digitCheck = regexp.MustCompile(`^[0-9]+$`)

// Repo represents a repository.
type Repo struct {
	RepoName      *string
	Major         *string
	Minor         string
	Patch         string
	Latest        string
	LatestRelease Release
}

// Release represents a GitHub release.
type Release struct {
	URL         string    `json:"url"`
	AssetsURL   string    `json:"assets_url"`
	UploadURL   string    `json:"upload_url"`
	HTMLURL     string    `json:"html_url"`
	ID          int       `json:"id"`
	TagName     string    `json:"tag_name"`
	Name        string    `json:"name"`
	Draft       bool      `json:"draft"`
	Prerelease  bool      `json:"prerelease"`
	CreatedAt   time.Time `json:"created_at"`
	PublishedAt time.Time `json:"published_at"`
	Assets      []any     `json:"assets"`
	TarballURL  string    `json:"tarball_url"`
	ZipballURL  string    `json:"zipball_url"`
	Body        string    `json:"body"`
}

// getReleases retrieves the releases for the repository.
func (repo Repo) getReleases() []Release {
	var releases []Release
	wesbsite := fmt.Sprintf("https://api.github.com/repos/%s/releases", *repo.RepoName)
	//fmt.Println(wesbsite)
	resp, err = http.Get(wesbsite)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &releases)
	if err != nil {
		log.Fatal(err)
	}
	return repo.filterReleases(releases)

	//fmt.Println(repo.LatestRelease.Body)
}

// GetMajor returns the major version of the release.
func (release Release) GetMajor() (major int) {
	SemVerSplit := strings.Split(release.TagName, ".")[0]
	majorSplit := strings.Split(SemVerSplit, "-")
	for _, value := range majorSplit {
		if digitCheck.MatchString(value) {
			major, err = strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			break
		}
	}
	return major
}

// filterReleases filters the releases based on the major version.
func (repo Repo) filterReleases(releases []Release) []Release {
	var filteredReleases []Release
	for _, release := range releases {
		switch *repo.Major {
		case "None":
			if !release.Draft && !release.Prerelease && release.TagName != "" {
				filteredReleases = append(filteredReleases, release)
			}
		default:
			major := release.GetMajor()
			if strings.Contains(strconv.Itoa(major), *repo.Major) && !release.Draft && !release.Prerelease && release.TagName != "" {
				//fmt.Println(release.TagName)
				filteredReleases = append(filteredReleases, release)
			}
		}
	}
	if len(filteredReleases) == 0 {
		log.Fatalf(fmt.Sprintf("No releases found matching Major Version: %s", *repo.Major))
	}
	return filteredReleases
}

// getLatest retrieves the latest release.
func (repo *Repo) getLatest(releases []Release) {
	if *repo.Major == "None" {
		major_list := []int{}
		for _, release := range releases {
			major_list = append(major_list, release.GetMajor())
		}
		slices.Sort(major_list)
		majorVersion := major_list[len(major_list)-1]
		*repo.Major = strconv.Itoa(majorVersion)
		releases = repo.filterReleases(releases)
	}
	slices.SortFunc(releases, func(i, j Release) int {
		if strings.Contains(i.TagName, "-") && strings.Contains(j.TagName, "-") {
			return cmp.Compare(strings.Split(strings.ToLower(i.TagName), "-")[1], strings.Split(strings.ToLower(j.TagName), "-")[1]) // sort by tag name
		} else {
			return cmp.Compare(strings.ToLower(i.TagName), strings.ToLower(j.TagName)) // sort by tag name
		}
	})

	repo.Latest = releases[len(releases)-1].TagName
	repo.LatestRelease = releases[len(releases)-1]
}

func main() {
	var repo Repo
	// repo.RepoName = "gamethis/bonkeywonkers"
	//repo.RepoName = "frrouting/frr"
	//repo.Major = "10"
	repo.RepoName = flag.String("repo_name", "", "The name of the repository in the format of 'owner/repo'")
	repo.Major = flag.String("major", "None", "The major version of the release")
	repo.Minor = *flag.String("minor", "None", "The minor version of the release")
	repo.Patch = *flag.String("patch", "None", "The patch version of the release")
	flag.Parse()
	if *repo.RepoName == "" {
		log.Fatalf("Please provide a repository name")
	}
	filteredReleases := repo.getReleases()
	repo.getLatest(filteredReleases)
	fmt.Printf("version=%s\n", repo.Latest)
	//fmt.Println(repo.LatestRelease.Body)

}
