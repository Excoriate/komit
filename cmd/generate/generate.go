package generate

import (
	"log"
	"os"

	"github.com/spf13/viper"

	"github.com/excoriate/komit/cmd/generate/commit"

	"github.com/excoriate/komit/internal/ui"

	"github.com/excoriate/komit/internal/app"
	"github.com/excoriate/komit/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	ignoreTokenLimit bool
)

var CMD = &cobra.Command{
	Use:   "generate",
	Short: "Generate several related things, such as conventional commit messages, pull request titles, etc.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
		}

		ctx := cmd.Context().Value(cli.GetCtxKey())
		if ctx == nil {
			log.Fatal("Unable to get the client context.")
		}

		_, ok := ctx.(*app.App)
		if !ok {
			ui.Error("", "Unable to obtain the App (Client) from the command context.", nil)
			os.Exit(1)
		}
	},
}

func init() {
	CMD.AddCommand(commit.CMD)
	CMD.Flags().BoolVarP(&ignoreTokenLimit, "ignore-token-limit", "", false, "Ignore the tokens limit (by default, 2048 tokens) when generating data through the AI provider.")

	_ = viper.BindPFlags(CMD.Flags())
}
