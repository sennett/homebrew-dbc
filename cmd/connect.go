/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"flag"
	"fmt"

	"github.com/birdiecare/dbc/handler"
	"github.com/spf13/cobra"
)

var host string
var port string
var region string
var user string
var iam bool

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a DB",
	Long: `Connect to a DB.
Opens a conneciton with the given database @ the given port at localhost:localport (-lp).
This command will output a Token to use as a password when connecting to your database.	

	psql -h localhost -p ${localport} -U ${user} -d ${database} --password

Then paste the token`,

	Run: func(cmd *cobra.Command, args []string) {
		flag.Parse()

		if iam {
			token := handler.HandleToken(host, port, region, user)
			fmt.Println("Token: ", token)
		}
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

	//Flags
	connectCmd.Flags().StringVarP(&host, "host", "H", "", "Hostname of the Database to open a connection to. If a hostname is not provided, a fuzzyfind list with be presented to select a database, and subsequently, a user to connect to`")
	connectCmd.Flags().StringVarP(&port, "port", "p", "5432", "Port of the Datbase to open a connection to (default 5432)")
	connectCmd.Flags().StringVarP(&region, "region", "r", "eu-west-2", "Region of the Datbase to open a connection to (default eu-west-2)")
	connectCmd.Flags().StringVarP(&user, "user", "u", "", "The DB User to open a connection with")
	connectCmd.Flags().BoolVarP(&iam, "iam", "I", false, "Bool: Use IAM Authentication for Database Connection - Generates a password token using IAM Authentication")
}
