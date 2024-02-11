package app

import "github.com/excoriate/komit/pkg/gitrepo"

type GitRepo interface {
	Diff(repoPath string) (*gitrepo.DiffOutput, error)
	DiffAsString(repoPath string) (string, error)
}

type GitRepository struct {
	app *App
}

func NewGitRepo(app *App) GitRepo {
	return &GitRepository{
		app: app,
	}
}

func (g *GitRepository) DiffAsString(repoPath string) (string, error) {

	diff, err := gitrepo.GetDiff(repoPath)
	if err != nil {
		return "", err
	}

	return diff.DiffString, nil
}

func (g *GitRepository) Diff(repoPath string) (*gitrepo.DiffOutput, error) {
	return getDiff(repoPath)
}

func getDiff(repoPath string) (*gitrepo.DiffOutput, error) {
	if repoPath == "" {
		repoPath = "."
	}

	if err := gitrepo.IsGitRepository(repoPath); err != nil {
		return nil, err
	}

	return gitrepo.GetDiff(repoPath)
}
