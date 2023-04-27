package cmd

import (
	"github.com/spf13/cobra"
	"go1/pkg/console"
	"go1/pkg/helpers"
)

var CmdKey = &cobra.Command{
	Use:   "key",
	Short: "generate key ...",
	Run:   runKeyGenerate,
	Args:  cobra.NoArgs,
}

func runKeyGenerate(cmd *cobra.Command, args []string) {
	console.Success("---")
	console.Success("App Key:")
	console.Success(helpers.RandomString(32))
	console.Success("---")
	console.Warning("please go to .env file to change the APP_KEY option")
}
