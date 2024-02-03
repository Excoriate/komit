package main

import "github.com/excoriate/komit/cmd"

var (
	version = "dev"
	commit  = "HEAD"
	date    = "unknown"
)

func main() {
	cmd.Execute(version, commit, date)
}
