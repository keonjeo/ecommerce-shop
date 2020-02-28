package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "shop",
	Short: "E-commerce shop app example",
}

// Execute executes the root command and adds all child commands to it
// It is called in main.go
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	fmt.Println("Config Initialized")
}

func init() {
	cobra.OnInitialize(initConfig)
}
