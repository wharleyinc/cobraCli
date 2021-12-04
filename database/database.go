package database

import (
	"encoding/json"
	"os"
	"strings"
	"time"
)

// Customer represents data about a customer.
type Customer struct {
	//ID     string  `json:"id"`	// not re
	Username  string  `json:"username"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
	Transactions	[]Transaction	`json:"transactions"`
	DateModified time.Time  `json:"date_modified"`
	LoginStatus bool  `json:"login_status"`
}

type Transaction struct {
	//ID     string  `json:"id"`
	Amount     float64  `json:"amount"`
	Transaction     TransactionType  `json:"transaction"`
	DateCaptured     time.Time  `json:"date_captured"`
}


type TransactionType string

const (
	TransactionTypeCredit TransactionType = "Credit"
	TransactionTypeDebit  TransactionType = "Debit"
	TransactionTypeTransfer TransactionType = "Transfer"
)


func GetCustomers() ([]Customer, error) {
	data, err := os.ReadFile("database/database.json")
	var customers []Customer
	if err == nil {
		json.Unmarshal(data, &customers)
	}
	return customers, err
}
func UpdateDB(data []Customer) {
	bytes, err := json.Marshal(data)
	if err == nil {
		err := os.WriteFile("database/database.json", bytes, 0644)
		if err != nil {
			return
		}
	}
}
func FindCustomer(username string) (*Customer, error) {
	customers, err := GetCustomers()
	if err == nil {
		for i := 0; i < len(customers); i++ {
			customer := customers[i]
			if strings.EqualFold(customer.Username, username) {
				return &customer, nil
			}
		}
	}
	return nil, err
}
func CreateCustomer(username, password string) (*Customer, error) {
	customer, err := FindCustomer(username)
	if customer == nil {
		var newCustomer Customer
		newCustomer.Username = strings.ToLower(username)
		newCustomer.Balance = 0
		newCustomer.Password = password
		newCustomer.Transactions = []Transaction{}
		newCustomer.LoginStatus = false
		newCustomer.DateModified = time.Now()
		customers, err := GetCustomers()
		if err == nil {
			customers = append(customers, newCustomer)
			UpdateDB(customers)
		}
		return &newCustomer, err
	}
	return customer, err
}
func UpdateCustomer(customer *Customer) {
	// Update the json with this modified customer information
	customers, err := GetCustomers()
	if err == nil {
		for i := 0; i < len(customers); i++ {
			if strings.EqualFold(customers[i].Username, customer.Username) {
				// Update the user details
				customers[i] = *customer
			}
		}
		// update database
		UpdateDB(customers)
	}
}
