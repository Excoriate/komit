package app

import (
	"errors"
	"fmt"

	"github.com/excoriate/komit/internal/erroer"
)

var generateCommitMessageAllowedOptions = []string{
	"simple", "conventional", "advanced", "semantic",
}

func isCommitMessageTypeOptionValid(option string) error {
	for _, v := range generateCommitMessageAllowedOptions {
		if v == option {
			return nil
		}
	}

	return fmt.Errorf("commit message type option %s not supported", option)
}

type Generate interface {
	GitCommitMessage(commitMsgTypeOption string) (Commit, error)
}

type GenerateAction struct {
	app     *App
	gitRepo GitRepo
}

func NewGenerate(app *App) Generate {
	return &GenerateAction{
		app:     app,
		gitRepo: NewGitRepo(app),
	}
}

func (g *GenerateAction) GitCommitMessage(commitMsgTypeOption string) (Commit, error) {
	if err := isCommitMessageTypeOptionValid(commitMsgTypeOption); err != nil {
		return nil, err
	}

	// New commit
	commit := &GitCommit{}
	switch commitMsgTypeOption {
	case "simple":
		commit.messageTemplate = commit.GetSimple()

	case "advanced":
		commit.messageTemplate = commit.GetAdvanced()

	case "semantic":
		commit.messageTemplate = commit.GetSemantic()

	case "conventional":
		commit.messageTemplate = commit.GetConventional()

	default:
		return nil, fmt.Errorf("commit message type option %s not supported", commitMsgTypeOption)
	}

	//Getting the repository difference.
	diff, err := g.gitRepo.Diff(".")
	if err != nil {
		return nil, err
	}

	g.app.Log.Debug("Diff: %s", diff)

	prompt := commit.AddDiff(commit.messageTemplate, diff.DiffString)
	commit.messageCompiled = prompt

	g.app.Log.Debug("Generated commit message: %s", prompt)

	// Files changed
	var filesOnly []string
	var promptWithFilesOnly string

	// Calling AI provider to get the completion.
	commitGeneratedFromAI, aiErr := g.app.AIProvider.GetCompletion(g.app.Ctx, commit.GetCompiled())
	if aiErr != nil {
		// If error is due to max tokens, we can handle it by using only the files changed to re-adequately generate the commit message.
		if errors.Is(aiErr, &erroer.ErrMaxTokensExceedApplicationLimit{}) {
			g.app.Log.Warn("Max tokens exceed application limit. Retrying with only the files changed.")

			// Storing only the files (to in the go-git terminology)
			for _, fileDiff := range diff.FileDiffs {
				_, toFile := fileDiff.Files()
				if toFile != nil {
					filesOnly = append(filesOnly, toFile.Path())
				}
			}

			promptWithFilesOnly = commit.AddDiffWithFiles(commit.GetTemplate(), filesOnly)
		} else {
			return nil, aiErr
		}
	}

	commitGeneratedFromAIWithFilesOnly, aiFilesErr := g.app.AIProvider.GetCompletion(g.app.Ctx, promptWithFilesOnly)
	if aiFilesErr != nil {
		return nil, aiFilesErr
	}

	commit.messageCompiled = commitGeneratedFromAIWithFilesOnly
	return commit, nil
}
