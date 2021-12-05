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
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// transferCmd represents the transfer command
var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "A quick command to make fund transfers from an account to another.",
	Long: `This command initiates a fund transfer from customer's account
 to another customer's account, identified by their usernames.
 This will be successful only if the usernames are in the datastore.
This should return a success message.

#usage: cobraCli transfer <source account username> <space> <amount> <space> <destination account username>`,
	Run: func(cmd *cobra.Command, args []string) {
		fundTransfer(args)
		// fmt.Println("transfer called")		// code clean-up
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
	var transaction1, transaction2 database.Transaction
	payingUsername := args[0]
	amount := args[1]
	receiverUsername := args[2]

	newAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		fmt.Println("please enter valid amount in digits")
		return
	}

	customer1, err := database.FindCustomer(payingUsername)
	if err != nil {
		fmt.Println("User with "+ payingUsername + " cannot be found")
		return
	}
	customer2, err := database.FindCustomer(receiverUsername)
	if err != nil {
		fmt.Println("User with "+ receiverUsername + " cannot be found")
		return
	}

	fmt.Println("payingUsername supplied is: " + payingUsername)
	fmt.Println("receiverUsername supplied is: " + receiverUsername)
	fmt.Printf("\n\n\n attempting to transfer from  user: %s to  user %s. ", payingUsername, receiverUsername)

	transaction1.Transaction = database.TransactionTypeDebit
	transaction1.Amount = newAmount
	transaction1.DateCaptured = time.Now()
	balance1 = customer1.Balance - newAmount
	fmt.Printf("new balance for %s  is: %0.2f", payingUsername, balance1)
	// add transaction to Customer object 1
	customer1.Transactions = append(customer1.Transactions, transaction1)

	// compute a transaction object 2
	transaction2.Transaction = database.TransactionTypeCredit
	transaction2.Amount = newAmount
	transaction2.DateCaptured = time.Now()
	balance2 = customer2.Balance - newAmount
	fmt.Printf("new balance for %s  is: %0.2f", receiverUsername, balance2)
	// add transaction to Customer object 1
	customer2.Transactions = append(customer2.Transactions, transaction2)
	return


}

