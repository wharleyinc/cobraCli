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
	"cobraCli/models"
	"fmt"

	"github.com/spf13/cobra"
)

// userLoginCmd represents the userLogin command
var userLoginCmd = &cobra.Command{
	Use:   "userLogin",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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


	for _, a := range models.Customers {
		if a.Username == username {
			if a.Password == password {
				a.LoginStatus = true
				fmt.Println("login is successful")
				fmt.Println(username + ": Welcome to Wale's CLI using Cobra c/o Segun Mustapha")
			} else {
				fmt.Println("login unsuccessful. try again later")
				return
			}
		}
	}

}