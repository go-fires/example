package commands

import (
	"github.com/go-fires/example/app/facades"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		facades.Http().Run("127.0.0.1:8081")
	},
}
