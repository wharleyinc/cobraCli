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
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// transferCmd represents the transfer command
var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fundTransfer(args)
		fmt.Println("transfer called")
	},
}

func init() {
	rootCmd.AddCommand(transferCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// transferCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// transferCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func fundTransfer(args []string)  {

	var balance1, balance2 float64
	var transaction1, transaction2 models.TransactionHistory
	payingUsername := args[0]
	amount := args[1]
	receiverUsername := args[2]

	newAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		fmt.Println("please enter valid amount in digits")
	}

	fmt.Println("payingUsername supplied is: " + payingUsername)
	fmt.Println("receiverUsername supplied is: " + receiverUsername)
	fmt.Printf("\n\n\n attempting to transfer from  user: %s to  user %s. ", payingUsername, receiverUsername)

	for i := 0; i < len(models.Customers); i++ {
		if models.Customers[i].Username == payingUsername {
			for j := i + 1; j < len(models.Customers); j++ {
				if models.Customers[j].Username == receiverUsername {
					// compute a transaction object 1
					transaction1.Transaction = models.TransactionTypeDebit
					transaction1.Amount = newAmount
					transaction1.DateCaptured = time.Now()
					balance1 = models.Customers[i].Balance - newAmount
					fmt.Printf("new balance for %s  is: %0.2f", payingUsername, balance1)
					// add transaction to Customer object 1
					models.Customers[i].Transactions = append(models.Customers[i].Transactions, transaction1)

					// compute a transaction object 2
					transaction2.Transaction = models.TransactionTypeCredit
					transaction2.Amount = newAmount
					transaction2.DateCaptured = time.Now()
					balance2 = models.Customers[j].Balance + newAmount
					fmt.Printf("new balance for %s  is: %0.2f", receiverUsername, balance2)
					// add transaction to Customer object 1
					models.Customers[i].Transactions = append(models.Customers[i].Transactions, transaction1)
					return
				}
			}
		}
		fmt.Println("credentials entered are incorrect. try again")
	}

}

