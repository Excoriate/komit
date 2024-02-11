package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/excoriate/komit/pkg/cli"

	"github.com/excoriate/komit/internal/app"
	"github.com/excoriate/komit/internal/ui"

	"github.com/excoriate/komit/cmd/generate"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var (
	// Command global flags.
	provider string
	apikey   string
	debug    bool
	model    string
)

const CLIName = "komit"

var rootCmd = &cobra.Command{
	Use:   CLIName,
	Short: "A CLI tool to generate different things in Git, with AI, for lazy developers.",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		apikeyViper := viper.GetString("apikey")
		if apikeyViper != "" {
			apikey = apikeyViper
		}

		apikeyEnv := os.Getenv("KOMIT_PROVIDER_APIKEY")
		if apikeyEnv != "" {
			apikey = apikeyEnv
		}

		// if it's empty, fail.
		if apikey == "" {
			ui.Error("", "The (AI) provider's API KEY is not set. You can either set it as an environment variable (KOMIT_PROVIDER_APIKEY) or pass it as a flag --apikey.", nil)
			os.Exit(1)
		}

		ctx := cmd.Context()
		app, err := app.New(ctx, &app.AIProviderOptions{
			Name:      provider,
			AuthToken: apikey,
			Model:     viper.GetString("model"),
		})

		if err != nil {
			ui.Error("", "Cannot initialize Komit", err)
			os.Exit(1)
		}

		// Storing the client/App in the context.
		cmdCtx := context.WithValue(context.Background(), cli.GetCtxKey(), app)
		cmd.SetContext(cmdCtx)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addPersistentFlags() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enabled debug mode")
	rootCmd.PersistentFlags().StringVarP(&provider, "provider", "", "openai", "AIProvider to use")
	rootCmd.PersistentFlags().StringVarP(&apikey, "apikey", "", "", "API Key to use")
	rootCmd.PersistentFlags().StringVarP(&model, "model", "", "gpt-3.5-turbo", "Model to use. Since the default provider is OpenAI, the default model is gpt-3.5-turbo")

	_ = viper.BindPFlags(rootCmd.PersistentFlags())
}

func initConfig() {
	homeDir, _ := os.UserHomeDir()
	configDir := fmt.Sprintf("%s/.komit", homeDir)

	viper.AddConfigPath(configDir)
	viper.SetConfigName(CLIName)
	viper.SetConfigType("yaml")
	viper.Set("CLI_NAME", CLIName)
	viper.SetEnvPrefix(strings.ToUpper(CLIName))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			_ = viper.SafeWriteConfigAs(fmt.Sprintf("%s/%s.yaml", configDir, CLIName))
		}
	}
}

func init() {
	rootCmd.AddCommand(generate.CMD)
	cobra.OnInitialize(initConfig)
	addPersistentFlags()
}
