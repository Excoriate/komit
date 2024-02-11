package app

import "strings"

const (
	// Enhanced prompt for generating a conventional commit message
	promptCommitConventional = `
Generate a conventional commit message that clearly describes the changes made.
Start with the type of change (e.g., feat, fix), optionally include the scope,
and add a concise description in imperative mood. Format: <type>(<scope>): <description>.
Example: "feat(parser): add ability to parse arrays."
Include a body for detailed text and a footer for issue references if applicable.
Consider the following DIFF to provide context for the commit message, and understand the changes made.
`

	// Enhanced prompt for generating a simple commit message
	promptCommitSimple = `
Write a simple commit message summarizing the changes succinctly.
Begin with an imperative verb (e.g., Add, Fix, Update) to describe the action taken.
Keep it brief yet informative. Example: "Update README with new installation instructions."
Consider the following DIFF to provide context for the commit message, and understand the changes made.
`

	// Enhanced prompt for generating an advanced commit message
	promptCommitAdvanced = `
Create an advanced commit message following the conventional format.
Begin with the change type and scope, followed by a short description.
Elaborate on the reasons for the changes, related issues, and impact in the message body.
End with a footer for "BREAKING CHANGE" or issue references.
Example: "feat(database): add 'archive' option to 'users' table for data retention. Resolves #123, relates to #456."
Consider the following DIFF to provide context for the commit message, and understand the changes made.
`

	// New prompt for generating a semantic commit message
	promptCommitSemantic = `
Construct a semantic commit message that precisely conveys the changes' impact on the versioning and codebase.
Begin with a change type (feat for minor feature, fix for patch, BREAKING CHANGE for major change),
optionally specify the scope, and provide a succinct description.
Format: <type>(<scope>): <description>. Include reasons, related issues, and impacts in the body.
Footer may list breaking changes or issue references.
Example: "BREAKING CHANGE(auth): overhaul authentication system to improve security. Affects all login endpoints. Fixes #789."
Consider the following DIFF to provide context for the commit message, and understand the changes made.
`
)

type Commit interface {
	GetSimple() string
	GetConventional() string
	GetAdvanced() string
	GetSemantic() string // Added getter for semantic commit
	AddDiff(prompt, diff string) string
	AddDiffWithFiles(prompt, files []string) string
	GetCompiled() string
	GetTemplate() string
}

type GitCommit struct {
	messageTemplate string
	messageCompiled string
}

func (gc *GitCommit) GetSimple() string {
	return promptCommitSimple
}

func (gc *GitCommit) GetConventional() string {
	return promptCommitConventional
}

func (gc *GitCommit) GetAdvanced() string {
	return promptCommitAdvanced
}

func (gc *GitCommit) GetSemantic() string { // Implement getter for semantic commit
	return promptCommitSemantic
}

func (gc *GitCommit) AddDiff(prompt, diff string) string {
	// Implementation to append diff details to the prompt, enhancing the context for generating commit messages
	return prompt + "\n\nDiff:\n" + diff
}

func (gc *GitCommit) AddDiffWithFiles(prompt string, files []string) string {
	// Implementation to append diff details to the prompt, enhancing the context for generating commit messages
	return prompt + "\n\nFiles:\n" + strings.Join(files, "\n")
}

func (gc *GitCommit) GetCompiled() string {
	return gc.messageCompiled
}

func (gc *GitCommit) GetTemplate() string {
	return gc.messageTemplate
}
