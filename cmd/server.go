package cmd

import (
	"log"

	api "github.com/dankobgd/ecommerce-shop/api/v1"
	"github.com/dankobgd/ecommerce-shop/app"
	"github.com/dankobgd/ecommerce-shop/config"
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
	db, err := postgres.Connect()
	if err != nil {
		return err
	}

	pgst := postgres.NewStore(db)

	server, err := app.NewServer(pgst)
	if err != nil {
		log.Fatalf("could not create the app server: %v\n", err)
	}

	cfg := config.New()

	appOpts := []app.Option{
		app.SetConfig(cfg),
		app.SetServer(server),
	}

	a := app.New(appOpts...)

	api.Init(a, server.Router)

	if srvErr := server.Start(); srvErr != nil {
		log.Fatalf("could not start the server: %v\n", srvErr)
	}

	return nil
}
