package cmd

import (
	"log"

	"github.com/dankobgd/ecommerce-shop/apiv1"
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
	// server, err := app.NewServer()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer server.Shutdown()

	// api := apiv1.Init(server.Router)

	// serverErr = server.Start()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	server, err := app.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	api := apiv1.Init(server.Router)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		cancel()
	}()

	if err := serve(ctx); err != nil {
		log.Printf("failed to serve:+%v\n", err)
	}

	serverErr = server.Start(); serverErr != nil {
		log.Printf("failed to serve:+%v\n", err)
	}	
}
