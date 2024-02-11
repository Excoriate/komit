package commit

import (
	"log"
	"os"

	"github.com/spf13/viper"

	"github.com/excoriate/komit/internal/ui"

	"github.com/excoriate/komit/internal/app"
	"github.com/excoriate/komit/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	commitType string
)

var CMD = &cobra.Command{
	Use:   "commit",
	Short: "Generate git commit messages with AI",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
		}
		ctx := cmd.Context().Value(cli.GetCtxKey())
		if ctx == nil {
			log.Fatal("Unable to get the client context.")
		}

		app, ok := ctx.(*app.App)
		if !ok {
			ui.Error("", "Unable to obtain the App (Client) from the command context.", nil)
			os.Exit(1)
		}

		commitTypeViper := viper.GetString("type")
		if commitTypeViper != "" {
			commitType = commitTypeViper
		}

		if commitType == "" {
			ui.Error("", "The commit type is not set. You can either set it as a flag --type, or in the configuration file.", nil)
			os.Exit(1)
		}

		r, err := app.Generate.GitCommitMessage(commitType)
		if err != nil {
			ui.Error("", "Cannot generate the commit message", err)
			os.Exit(1)
		}

		ui.Info("", r.GetCompiled())
	},
}

func init() {
	CMD.Flags().StringVarP(&commitType, "type", "", "conventional", "The type of the commit to generate. Possible values: simple, conventional, semantic, advanced, etc.")
	_ = viper.BindPFlags(CMD.Flags())
}
