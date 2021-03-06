/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"cobraCli/database"
	"fmt"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "creates a new customer record taking in a simple username and password string entries.",
	Long: `This command creates a new customer record passing in the username 
and password. This will be successful only if the username is not already 
in the datastore.
This should return a success message.
#usage: cobraCli create <username> <space> <password>`,
	Run: func(cmd *cobra.Command, args []string) {
		create(args)
		// fmt.Println("create called")		// code clean-up
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func create(args []string)  {

	username := args[0]
	password := args[1]
	fmt.Println("username supplied is: " + username)
	fmt.Println("password supplied is: " + password)

	customer, err := database.CreateCustomer(username, password)
	if err != nil {
		fmt.Println("User with "+ username + "already exist. please login ")
		return
	}
	fmt.Println(customer.Username + ": Welcome to Wale's CLI using Cobra c/o Segun Mustapha")

}