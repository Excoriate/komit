package gitrepo

import (
	"context"
	"fmt"

	"github.com/go-git/go-git/v5/plumbing/format/diff"

	"github.com/go-git/go-git/v5"
)

// DiffOutput stores information about the git diff.
type DiffOutput struct {
	DiffString string
	FileDiffs  []diff.FilePatch
}

// GetDiff returns the diff between the current commit and its parent.
func GetDiff(repoPath string) (*DiffOutput, error) {
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open git repository: %w", err)
	}

	// Get the HEAD commit to compare against.
	head, err := r.Head()
	if err != nil {
		return nil, fmt.Errorf("failed to get HEAD: %w", err)
	}

	commit, err := r.CommitObject(head.Hash())
	if err != nil {
		return nil, fmt.Errorf("failed to get commit object: %w", err)
	}

	// Get the parent commit of the current commit.
	parentCommit, err := commit.Parent(0)
	if err != nil {
		return nil, fmt.Errorf("failed to get parent commit: %w", err)
	}

	// Get the diff between the current commit and its parent.
	patch, err := parentCommit.PatchContext(context.Background(), commit)
	if err != nil {
		return nil, fmt.Errorf("failed to get diff: %w", err)
	}

	// Prepare the output
	output := &DiffOutput{
		DiffString: patch.String(),
		FileDiffs:  patch.FilePatches(),
	}

	return output, nil
}
