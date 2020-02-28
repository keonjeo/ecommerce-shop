package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "DB migrate",
	Long:  "Perform database migrations",
	RunE:  migrateCmdFn,
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func migrateCmdFn(cmd *cobra.Command, args []string) error {
	fmt.Println("migrate called")
	return nil
}
