package generate

import (
	"log"
	"os"

	"github.com/excoriate/komit/internal/ui"

	"github.com/excoriate/komit/internal/app"
	"github.com/excoriate/komit/pkg/cli"
	"github.com/spf13/cobra"
)

var CMD = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new commit (conventional) message",
	Run: func(cmd *cobra.Command, args []string) {
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

// func init() {
//}
