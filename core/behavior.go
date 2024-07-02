package core

import (
	"container/heap"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Calculate(
	expenses map[primitive.ObjectID]Expense,
) []PaymentNode {
	balances := simulateExpense(expenses)
	paymentGraph := makePaymentGraph(balances)
	return paymentGraph
}

// abs helper function for absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// makePaymentGraph builds a payment graph based on user balances
func makePaymentGraph(balances map[primitive.ObjectID]int) []PaymentNode {
	firstHeap := &MaxHeap{}
	secondHeap := &MaxHeap{}

	graph := []PaymentNode{}

	// Initialize heaps with user balances (absolute value for second heap)
	for userID, amount := range balances {
		node := &Node{UserID: userID, FinalBalance: amount}
		heap.Push(firstHeap, node)
		heap.Push(secondHeap, &Node{UserID: userID, FinalBalance: abs(amount)})
	}

	// Loop until both heaps are empty
	for firstHeap.Len() > 0 {
		receiver := heap.Pop(firstHeap).(*Node)
		sender := heap.Pop(secondHeap).(*Node)

		// Calculate transfer amount (minimum of sender and receiver balance)
		amountTransferred := min(sender.FinalBalance, receiver.FinalBalance)

		// Add payment node to graph
		graph = append(graph, PaymentNode{From: sender.UserID, To: receiver.UserID, Amount: amountTransferred})

		// Update sender and receiver balances
		sender.FinalBalance -= amountTransferred
		receiver.FinalBalance -= amountTransferred

		// Push back non-zero balances to respective heaps
		if sender.FinalBalance != 0 {
			heap.Push(secondHeap, sender)
		}
		if receiver.FinalBalance != 0 {
			heap.Push(firstHeap, receiver)
		}
	}

	return graph
}

// min helper function for finding minimum
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// simulateExpense simulates an expense with user balances
func simulateExpense(expense map[primitive.ObjectID]Expense) map[primitive.ObjectID]int {
	balances := map[primitive.ObjectID]int{}
	for _, expense := range expense {
		for UserID, amount := range expense.Balances {
			balances[UserID] += amount
		}
	}
	return balances
}
