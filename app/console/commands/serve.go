package commands

import (
	"fmt"
	"github.com/go-fires/example/app/facades"
	"github.com/go-fires/example/app/providers/http"
	"github.com/go-fires/framework/facade"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		err := facades.Http().Run(facade.Config().Get("http").(*http.Config).Address)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
