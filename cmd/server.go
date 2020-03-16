package cmd

import (
	"fmt"
	"log"

	api "github.com/dankobgd/ecommerce-shop/api/v1"
	"github.com/dankobgd/ecommerce-shop/app"
	"github.com/dankobgd/ecommerce-shop/store/postgres"
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
	server, err := app.NewServer()
	if err != nil {
		log.Fatalf("could not create the app server: %v\n", err)
	}

	api.Init(server.Router)

	// #################################################
	db, _ := postgres.Connect()
	var res string
	_ = db.Get(&res, "select * from user;")
	fmt.Printf("result user: %v\n", res)
	// #################################################

	if srvErr := server.Start(); srvErr != nil {
		log.Fatalf("could not start the server: %v\n", srvErr)
	}

	return nil
}
