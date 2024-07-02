package splitwise

type FactorList[T comparable] []Factor[T]

type Factor[T comparable] struct {
	Id    T
	Buyer T
	Price int
	Users []UserWithCoefficient[T]
}

func (f *Factor[T]) toExpense() Expense[T] {
	expense := make(Expense[T])

	coefficientsSum := 0
	for _, user := range f.Users {
		coefficientsSum += user.Coefficient
	}

	expense[f.Buyer] = -f.Price

	for _, user := range f.Users {
		if user.UserId == f.Buyer {
			expense[user.UserId] += (f.Price / coefficientsSum) * user.Coefficient
			continue
		}
		expense[user.UserId] = (f.Price / coefficientsSum) * user.Coefficient
	}

	return expense
}

type UserWithCoefficient[T comparable] struct {
	UserId      T
	Coefficient int
}

type TransactionList[T comparable] []*Transaction[T]

func createRelativeTransactions[T comparable](t *TransactionList[T]) {
	if t == nil || len(*t) == 0 {
		return
	}
	buyer := (*t)[len(*t)-1]
	buyerReletiveTransactions := make(TransactionList[T], 0)
	for i, balance := range *t {
		if i == len(*t)-1 {
			continue
		}
		if balance.ReletiveTransaction == nil {
			balance.ReletiveTransaction = &TransactionList[T]{}
		}
		balacneCopy := Transaction[T]{
			From:   balance.From,
			UserId: balance.UserId,
		}
		balacneCopy.Debt = nil
		balacneCopy.Demand = balance.Debt
		buyerReletiveTransactions = append(buyerReletiveTransactions, &Transaction[T]{From: buyer.UserId, UserId: balance.UserId, Debt: balance.Debt})
		*balance.ReletiveTransaction = append(*balance.ReletiveTransaction, &balacneCopy)
	}
	buyer.ReletiveTransaction = &buyerReletiveTransactions
	// fmt.Printf("t: %v\n", lib.ToJsonIndent(t))
}

type Transaction[T comparable] struct {
	From                T
	UserId              T
	Demand              *int
	Debt                *int
	ReletiveTransaction *TransactionList[T]
}
