package github

import (
	"fmt"
	git "github.com/go-git/go-git/v5"
	http "github.com/go-git/go-git/v5/plumbing/transport/http"
	"os"
)

var repo *git.Repository
var worktree *git.Worktree

func CloneRepo(dir string, repoURL string, ghUser string, ghToken string) {

	// Clone the given repository to the given directory
	fmt.Printf("git clone %s\n", repoURL)

	r, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL: repoURL,
		Auth: &http.BasicAuth{
			Username: ghUser, // yes, this can be anything except an empty string
			Password: ghToken,
		},
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Repository cloned")
	repo = r

	w, err := r.Worktree()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	worktree = w
}

func AddFile(path string) {
	fmt.Println(worktree.Status())
	worktree.Add(path)
	fmt.Println(worktree.Status())
}

func Commit(message string) {
	worktree.Commit(message, &git.CommitOptions{})
	fmt.Printf("Committed: %s\n", message)
}

func Push(ghUser string, ghToken string) {
	err := repo.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth: &http.BasicAuth{
			Username: ghUser, // yes, this can be anything except an empty string
			Password: ghToken,
		},
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Remote updated.")
}
