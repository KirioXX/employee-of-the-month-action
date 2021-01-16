package main

import (
	"fmt"
	"github.com/thatisuday/github-actions-golang-module/internal/file"
	"github.com/thatisuday/github-actions-golang-module/internal/giphy"
	"github.com/thatisuday/github-actions-golang-module/internal/github"
	"os"
	"os/signal"
	"syscall"
)

const ghUser = "Honkytonk123"

func main() {
	// Get env variables
	ghActor := os.Getenv("GITHUB_ACTOR")
	if ghActor == "" {
		fmt.Println("GITHUB_ACTOR environment variable is not set")
		os.Exit(1)
	}

	ghRepo := os.Getenv("GITHUB_REPOSITORY")
	if ghRepo == "" {
		fmt.Println("GITHUB_REPOSITORY environment variable is not set")
		os.Exit(1)
	}

	giphyApiKey := os.Getenv("GIPHY_API_KEY")
	if giphyApiKey == "" {
		fmt.Println("GIPHY_API_KEY environment variable is not set")
		os.Exit(1)
	}

	ghToken := os.Getenv("GITHUB_TOKEN")
	if ghToken == "" {
		fmt.Println("GITHUB_TOKEN environment variable is not set")
		os.Exit(1)
	}

	// Get input args
	tag, title, page := os.Args[1], os.Args[2], os.Args[3]

	// Get image
	getRandom := giphy.Init(giphyApiKey)
	image, err := getRandom(tag)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	wikiCommitMessage := os.Getenv("WIKI_COMMIT_MESSAGE")
	if wikiCommitMessage == "" {
		fmt.Println("WIKI_COMMIT_MESSAGE not set, using default")
		wikiCommitMessage = "Automatically publish wiki"
	}
	gitRepoURL := "https://github.com/" + ghRepo + ".wiki.git"
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		file.CleanDir()
		os.Exit(1)
	}()

	// Github Clone
	github.CloneRepo(file.Dir, gitRepoURL, ghUser, ghToken)

	// File manipulation
	res := file.ReadFile(page)
	temp, _ := file.GenTemplate(title, image)
	if file.HasMarkers(res) {
		newFileContent := file.ReplaceMarker(string(res), string(temp))
		file.WriteFile(page, []byte(newFileContent))
	} else {
		newFileContent := string(res) + string(temp)
		file.WriteFile(page, []byte(newFileContent))
	}

	// Push
	github.AddFile(page)
	github.Commit(wikiCommitMessage, ghUser)
	github.Push(ghUser, ghToken)

	file.CleanDir()
}
