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

// checkHistoryCmd represents the checkHistory command
var checkHistoryCmd = &cobra.Command{
	Use:   "checkHistory",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		checkHistory(args)
		fmt.Println("checkHistory called")
	},
}

func init() {
	rootCmd.AddCommand(checkHistoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkHistoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkHistoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func checkHistory(args []string)  {

	var sum int

	username := args[0]

	fmt.Println("username supplied is: " + username)
	fmt.Println("\n\n\n attempting to check transation history for user: " + username + " ...")


	customer, err := database.FindCustomer(username)
	if err != nil {
		fmt.Println("User with "+ username + " cannot be found")
		return
	}
	for _, b := range customer.Transactions {
		fmt.Println(b)
		sum += int(b.Amount)
	}
	fmt.Printf("Current Account Balance for User %s  is : %0.2f\n", customer.Username, sum)

	/*for _, a := range models.Customers {
		if a.Username == username {
			if a.LoginStatus {
				for _, b := range a.Transactions {
					fmt.Println(b)
					sum += int(b.Amount)
				}
				fmt.Printf("Current Account Balance for User %s  is : %0.2f\n", a.Username, sum)
			} else {
				fmt.Println("You're not logged. try again later")
				return
			}
		} else {
			fmt.Println("Username not found. kindly retry")
			return
		}
	}*/

}
