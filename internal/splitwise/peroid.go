package splitwise

type Peroid[T comparable] struct {
	Id           T
	Transactions *TransactionList[T]
	Factors      *FactorList[T]
}

func (p *Peroid[T]) EditFactors(factor Factor[T]) {
	for i, f := range *p.Factors {
		if f.Id == factor.Id {
			(*p.Factors)[i] = factor
			break
		}
	}

	expenses := make(map[T]Expense[T])
	for _, factor := range *p.Factors {
		expense := factor.toExpense()
		expenses[factor.Id] = expense
	}
	graph := newGraph(expenses)
	transactions := graph.toTransactions()
	createRelativeTransactions(transactions)
	p.Transactions = transactions
}

func (p *Peroid[T]) AddFactors(factors Factor[T]) {
	expenses := make(map[T]Expense[T])
	if p.Factors == nil {
		p.Factors = &FactorList[T]{}
	}
	*p.Factors = append(*p.Factors, factors)
	for _, factor := range *p.Factors {
		expense := factor.toExpense()
		expenses[factor.Id] = expense
	}
	graph := newGraph(expenses)
	transactions := graph.toTransactions()
	createRelativeTransactions(transactions)
	p.Transactions = transactions
}
