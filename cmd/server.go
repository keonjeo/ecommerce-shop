package cmd

import (
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
		log.Fatal(err)
	}

	api.Init(server.Router)

	if srvErr := server.Start(); srvErr != nil {
		log.Fatalf("could not start the server: %v\n", srvErr)
	}

	// #################################################
	log.Printf("STARTED CONNECTING")
	db, err := postgres.Connect()
	if err != nil {
		log.Fatalln(err)
		return err
	}
	res, _ := db.Query("select * from user;")
	log.Printf("result: %v", res)
	log.Printf("END CONNECTING")
	// #################################################

	return nil
}
