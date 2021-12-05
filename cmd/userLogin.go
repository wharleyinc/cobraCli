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
	"strings"

	"github.com/spf13/cobra"
)

// userLoginCmd represents the userLogin command
var userLoginCmd = &cobra.Command{
	Use:   "userLogin",
	Short: "A command to access the banking app cli",
	Long: `You're required to pass in your valid username and password,
so as to authenticate your identity on the banking app cli.
This should return a success/fail message'

#usage: cobraCli userLogin <username> <space> <password>`,
	Run: func(cmd *cobra.Command, args []string) {
		login(args)
	},
}

func init() {
	rootCmd.AddCommand(userLoginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userLoginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userLoginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func login(args []string)  {

	username := args[0]
	password := args[1]
	fmt.Println("username supplied is: " + username)
	fmt.Println("password supplied is: " + password)

	customer, err := database.FindCustomer(username)
	if err != nil {
		fmt.Println("User with "+ username + " cannot be found")
		return
	}

	if strings.EqualFold(customer.Password, password) {
		fmt.Println("login is successful")
		fmt.Println(username + ": Welcome to Wale's CLI using Cobra c/o Segun Mustapha")
		return
	} else {
		fmt.Println("incorrect password.")
		return
	}
	fmt.Println("login unsuccessful. try again later")
	return
/
}