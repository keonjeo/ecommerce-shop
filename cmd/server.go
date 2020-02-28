package cmd

import (
	"context"
	"fmt"

	"github.com/dankobgd/ecommerce-shop/app"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run server",
	Long:  "Run the API web server",
}

func init() {
	rootCmd.AddCommand(serverCmd)
	rootCmd.RunE = serverCmdFn
}

func serverCmdFn(command *cobra.Command, args []string) error {
	fmt.Println("RUN THE SERVER")

	// appOptions = []AppOption{"", "", ""}
	// server, err := app.NewServer()
	// api := apiv1.Init(server, server.AppOptions, server.Router)

	api := apiv1.New()
	server, err := app.NewServer()
	if err != nil {
		fmt.Println(err)
	}

	server.Start()
	defer server.Stop(context.Background())

	return nil
}
