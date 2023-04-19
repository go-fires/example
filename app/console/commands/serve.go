package commands

import (
	"fmt"
	"github.com/go-fires/example/app/facades"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		err := facades.Http().Run("127.0.0.1:8081")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
