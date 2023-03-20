/*
Copyright © 2023 Harry M harry.morgan@birdie.care

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dbc",
	Short: "Connect securely to RDS Databases",
	Long: `
Commands:

connect - Open a connection to a database
		
	dbc connect -h $database_host -u $user 
	`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
