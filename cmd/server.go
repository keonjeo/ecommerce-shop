package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

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
	server, err := app.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	apiv1.Init(server.Router)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		log.Printf("system call:%+v", <-done)
		cancel()
	}()

	if serverErr := server.Start(ctx); serverErr != nil {
		log.Printf("failed to serve:+%v\n", err)
	}

	return nil
}
