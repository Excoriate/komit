package gitrepo

import (
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5"
)

// IsGitRepository checks if the current directory is a git repository.
func IsGitRepository(repoPath string) error {
	if repoPath == "" {
		return fmt.Errorf("no repository path provided")
	}
	_, err := git.PlainOpen(".")
	if err != nil && errors.Is(err, git.ErrRepositoryNotExists) {
		return fmt.Errorf("not a git repository: %w", err)
	}

	return nil
}
