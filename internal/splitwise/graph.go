package splitwise

import (
	"container/heap"
	"fmt"

	"github.com/aamirmousavi/dong/lib"
)

type node[T comparable] struct {
	id      T
	balance int
}

type graph[T comparable] struct {
	payments []paymentNode[T]
}

func (g *graph[T]) toTransactions() *TransactionList[T] {
	transactions := make(TransactionList[T], 0)
	for _, node := range g.payments {

		var Demand, Debt *int
		if node.Amount < 0 {
			node.Amount *= -1
			Demand = &node.Amount
		} else if node.Amount > 0 {
			Debt = &node.Amount
		} else {
			continue
		}

		transactions = append(transactions, &Transaction[T]{
			From:   node.From,
			UserId: node.To,
			Demand: Demand,
			Debt:   Debt,
		})
	}
	return &transactions
}

type paymentNode[T comparable] struct {
	From   T
	To     T
	Amount int
}

func newGraph[T comparable](expenses map[T]Expense[T]) *graph[T] {
	balances := simulateExpense(expenses)
	fmt.Printf("balances: %v\n", lib.ToJsonIndent(balances))
	paymentNodes := calculatePaymentNodes(balances)
	fmt.Printf("paymentNodes: %v\n", lib.ToJsonIndent(paymentNodes))
	return &graph[T]{payments: paymentNodes}
}

func simulateExpense[T comparable](expenses map[T]Expense[T]) map[T]int {
	balances := map[T]int{}
	for _, expense := range expenses {
		for id, amount := range expense {
			balances[id] += amount
		}
	}
	return balances
}

func calculatePaymentNodes[T comparable](balances map[T]int) []paymentNode[T] {
	firstHeap := &maxHeap[T]{}
	secondHeap := &maxHeap[T]{}

	graph := []paymentNode[T]{}

	for id, balance := range balances {
		_node := &node[T]{id: id, balance: balance}
		heap.Push(firstHeap, _node)
		heap.Push(secondHeap, &node[T]{id: id, balance: abs(balance)})
	}

	for firstHeap.Len() > 0 {
		receiver := heap.Pop(firstHeap).(*node[T])
		sender := heap.Pop(secondHeap).(*node[T])

		amountTransferred := min(sender.balance, receiver.balance)

		graph = append(graph, paymentNode[T]{From: sender.id, To: receiver.id, Amount: amountTransferred})

		sender.balance -= amountTransferred
		receiver.balance -= amountTransferred

		if sender.balance != 0 {
			heap.Push(secondHeap, sender)
		}
		if receiver.balance != 0 {
			heap.Push(firstHeap, receiver)
		}
	}

	return graph
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
