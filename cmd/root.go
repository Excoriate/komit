package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Version string
	Commit  string
	Date    string
)

const CLIName = "komit"

var rootCmd = &cobra.Command{
	Version: "v0.0.1",
	Use:     CLIName,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(v string, c string, d string) {
	Version = v
	Commit = c
	Date = d
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addPersistentFlags() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringP("oai-key", "", "", "OpenAI API key")
}

func initConfig() {
	// Expose the CLI name to the Viper configuration.
	viper.Set("CLI_NAME_TITLE", CLIName)
}

func init() {
	addPersistentFlags()
	cobra.OnInitialize(initConfig)
}
