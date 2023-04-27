package cmd

import (
	"github.com/spf13/cobra"
	"go1/pkg/console"
)

var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "start ..",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {
	console.Success("success ...")
}
