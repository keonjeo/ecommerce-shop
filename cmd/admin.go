nteger generated always as identity primary key,package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "admin",
	Short: "Managment of users",
}

var createSuperAdminCmd = &cobra.Command{
	Use:     "super",
	Short:   "Create super admin",
	Long:    "Creates the super admin",
	Example: "  admin super --email user@example.com --username userexample --password password123",
	RunE:    createUserFn,
}

var createUserCmd = &cobra.Command{
	Use:     "createuser",
	Short:   "Create user",
	Long:    "Creates new user",
	Example: "  admin createuser --email user@example.com --username userexample --password password123",
	RunE:    createUserFn,
}

var deleteUserCmd = &cobra.Command{
	Use:     "deleteuser",
	Short:   "Delete user",
	Long:    "Deletes the user with the given id",
	Example: "  admin deleteuser --id 12345",
	RunE:    deleteUserFn,
}

func init() {
	createSuperAdminCmd.Flags().StringP("email", "e", "", "Required. The email address for the new user account.")
	createSuperAdminCmd.Flags().StringP("username", "u", "", "Required. Username for the new user account.")
	createSuperAdminCmd.Flags().StringP("password", "p", "", "Required. The password for the new user account.")
	createUserCmd.Flags().StringP("email", "e", "", "Required. The email address for the new user account.")
	createUserCmd.Flags().StringP("username", "u", "", "Required. Username for the new user account.")
	createUserCmd.Flags().StringP("password", "p", "", "Required. The password for the new user account.")
	deleteUserCmd.Flags().Int("id", 0, "Required. The ID for deleting the user.")

	userCmd.AddCommand(createSuperAdminCmd, createUserCmd, deleteUserCmd)
	rootCmd.AddCommand(userCmd)
}

func createSuperAdminFn(command *cobra.Command, args []string) error {
	email, erre := command.Flags().GetString("email")
	if erre != nil || email == "" {
		return errors.New("Email is required")
	}
	username, erru := command.Flags().GetString("username")
	if erru != nil || username == "" {
		return errors.New("Username is required")
	}
	password, errp := command.Flags().GetString("password")
	if errp != nil || password == "" {
		return errors.New("Password is required")
	}

	fmt.Println("CREATE SUPER ADMIN")
	return nil
}

func createUserFn(command *cobra.Command, args []string) error {
	email, erre := command.Flags().GetString("email")
	if erre != nil || email == "" {
		return errors.New("Email is required")
	}
	username, erru := command.Flags().GetString("username")
	if erru != nil || username == "" {
		return errors.New("Username is required")
	}
	password, errp := command.Flags().GetString("password")
	if errp != nil || password == "" {
		return errors.New("Password is required")
	}

	fmt.Println("CREATE USER")
	return nil
}

func deleteUserFn(command *cobra.Command, args []string) error {
	email, err := command.Flags().GetInt("id")
	if err != nil || email == 0 {
		return errors.New("ID is required")
	}
	fmt.Println("DELETE USER")
	return nil
}
