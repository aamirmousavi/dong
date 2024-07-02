package core

import "go.mongodb.org/mongo-driver/bson/primitive"

// Expense struct (simplified for this example)
type Expense struct {
	ExpenseID primitive.ObjectID         `json:"expense_id"`
	Balances  map[primitive.ObjectID]int `json:"balances"`
}

// Node struct for building the payment graph
type Node struct {
	UserID       primitive.ObjectID `json:"user_id"`
	FinalBalance int                `json:"final_balance"`
}

// PaymentNode struct for representing payment relationships
type PaymentNode struct {
	From   primitive.ObjectID `json:"from"`
	To     primitive.ObjectID `json:"to"`
	Amount int                `json:"amount"`
}
