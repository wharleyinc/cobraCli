package models

import (
	"time"
)

type ToDo struct {
	// ID        primitive.ObjectID `bson:"_id"` remove so that mongodb can auto generate ids for records
	Task      string    `bson:"task"`
	Status    string    `bson:"status"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

// Customer represents data about a customer.
type Customer struct {
	ID     string  `json:"id"`
	Username  string  `json:"username"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
	Transactions	[]TransactionHistory	`json:"transactions"`
	DateModified time.Time  `json:"date_modified"`
	LoginStatus bool  `json:"login_status"`
}

type TransactionHistory struct {
	ID     string  `json:"id"`
	Amount     float64  `json:"amount"`
	Transaction     TransactionType  `json:"transaction"`
	DateCaptured     time.Time  `json:"date_captured"`
}

/*type TransactionType struct {
	Debit  string  `json:"debit"`
	Credit  string  `json:"credit"`
	Transfer  string  `json:"transfer"`
}*/

// Customers to store a list of customers.
var Customers = []Customer{
	{ID: "1", Username: "olawale", Password: "oladapo", Balance: 2056.99, Transactions: []TransactionHistory{
		{
			ID:           "1",
			Amount:       24,
			Transaction:  TransactionTypeCredit,
			DateCaptured: time.Now(),
		},
		{
			ID:           "2",
			Amount:       2866,
			Transaction:  TransactionTypeTransfer,
			DateCaptured: time.Now(),
		},
		},
	LoginStatus: true, DateModified: time.Now()},
	{ID: "2", Username: "segun", Password: "mustpaha", Balance: 29022156.42, LoginStatus: false, DateModified: time.Now()},
	{ID: "3", Username: "jennifer", Password: "ezinne", Balance: 2004.80, LoginStatus: false, DateModified: time.Now()},
}

type TransactionType string

const (
	TransactionTypeCredit TransactionType = "Credit"
	TransactionTypeDebit  TransactionType = "Debit"
	TransactionTypeTransfer TransactionType = "Transfer"
)
