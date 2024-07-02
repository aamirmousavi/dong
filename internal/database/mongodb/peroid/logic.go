package peroid

import (
	"fmt"

	"github.com/aamirmousavi/dong/core"
	"github.com/aamirmousavi/dong/internal/database/mongodb/balance"
	"github.com/aamirmousavi/dong/lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (p *Peroid) Recalculate() {
	FactorCalculatedBalances := p.calculateFactorBalances()
	p.Balances = FactorCalculatedBalances
}

func (p *Peroid) AddFactor(factor *Factor) {
	if p.Factors == nil {
		p.Factors = &FactorList{factor}
	}
	*p.Factors = append(*p.Factors, factor)
	FactorCalculatedBalances := p.calculateFactorBalances()
	p.Balances = FactorCalculatedBalances
	if p.MoneySpend == nil {
		p.MoneySpend = make(map[primitive.ObjectID]int)
	}
	if _, ok := p.MoneySpend[factor.Buyer]; !ok {
		p.MoneySpend[factor.Buyer] = 0
	}
	p.MoneySpend[factor.Buyer] += factor.Price
	p.TotalPrice += factor.Price
	p.TotalFactor++
	p.AvgPrice = p.TotalPrice / p.TotalFactor
}

func (p *Peroid) AddPeyment(payment *balance.Payment) {
	if p.Payments == nil {
		p.Payments = &balance.PaymentList{payment}
	}
	*p.Payments = append(*p.Payments, payment)
	PaymentCalculatedBalances := p.calculatePaymentBalances()
	p.Balances = PaymentCalculatedBalances
}

func (p *Peroid) addPaymentBalanceMap(expenses *map[primitive.ObjectID]core.Expense) {
	for _, payment := range *p.Payments {
		expense := core.Expense{
			ExpenseID: payment.Id,
			Balances:  map[primitive.ObjectID]int{},
		}
		expense.Balances[payment.SourceUserId] = -payment.Amount
		expense.Balances[payment.TargetUserId] = payment.Amount
		(*expenses)[payment.Id] = expense
	}
}

func (p *Peroid) calculatePaymentBalances() *FactorCalculatedBalanceList {
	expenses := make(map[primitive.ObjectID]core.Expense)
	p.addFactorToExpensMap(&expenses)
	p.addPaymentBalanceMap(&expenses)

	paymentGraph := core.Calculate(expenses)
	listOfFactorCalculatedBalances := make(FactorCalculatedBalanceList, 0)
	for i, node := range paymentGraph {
		var Demand, Debt *int
		if node.Amount < 0 {
			node.Amount *= -1
			Demand = &node.Amount
		} else if node.Amount > 0 {
			Debt = &node.Amount
		} else {
			continue
		}
		reletiveFactorCalculatedBalances := make(FactorCalculatedBalanceList, 0)
		for j, reletiveNode := range paymentGraph {
			if i == j {
				continue
			}
			if reletiveNode.To != node.To && reletiveNode.From != node.To {
				continue
			}
			targetUser := reletiveNode.To
			if reletiveNode.To == node.To {
				targetUser = reletiveNode.From
			}
			var RelDemand, RelDebt *int
			if reletiveNode.Amount < 0 {
				reletiveNode.Amount *= -1
				RelDemand = &reletiveNode.Amount
			} else if reletiveNode.Amount > 0 {
				RelDebt = &reletiveNode.Amount
			} else {
				continue
			}
			reletiveFactorCalculatedBalances = append(reletiveFactorCalculatedBalances, &FactorCalculatedBalance{
				PeroidId: p.Id,
				UserId:   targetUser,
				Demand:   RelDemand,
				Debt:     RelDebt,
			})
		}
		listOfFactorCalculatedBalances = append(listOfFactorCalculatedBalances, &FactorCalculatedBalance{
			PeroidId:                         p.Id,
			UserId:                           node.To,
			Demand:                           Demand,
			Debt:                             Debt,
			ReletiveFactorCalculatedBalances: &reletiveFactorCalculatedBalances,
		})
	}
	return &listOfFactorCalculatedBalances
}

func (p *Peroid) addFactorToExpensMap(expenses *map[primitive.ObjectID]core.Expense) {
	for _, factor := range *p.Factors {
		expense := convertFactorToExpense(*factor)
		(*expenses)[factor.Id] = expense
	}
}

func (p *Peroid) calculateFactorBalances() *FactorCalculatedBalanceList {
	expenses := make(map[primitive.ObjectID]core.Expense)
	p.addFactorToExpensMap(&expenses)
	p.addPaymentBalanceMap(&expenses)
	paymentGraph := core.Calculate(expenses)
	listOfFactorCalculatedBalances := make(FactorCalculatedBalanceList, 0)
	fmt.Printf("paymentGraph: %v\n", lib.ToJsonIndent(paymentGraph))
	for i, node := range paymentGraph {
		var Demand, Debt *int
		if node.Amount < 0 {
			node.Amount *= -1
			Demand = &node.Amount
		} else if node.Amount > 0 {
			Debt = &node.Amount
		} else {
			continue
		}
		reletiveFactorCalculatedBalances := make(FactorCalculatedBalanceList, 0)
		for j, reletiveNode := range paymentGraph {
			if i == j {
				continue
			}
			if reletiveNode.To != node.To && reletiveNode.From != node.To {
				continue
			}
			var RelDemand, RelDebt *int
			if reletiveNode.Amount < 0 {
				reletiveNode.Amount *= -1
				RelDemand = &reletiveNode.Amount
			} else if reletiveNode.Amount > 0 {
				RelDebt = &reletiveNode.Amount
			} else {
				continue
			}
			reletiveFactorCalculatedBalances = append(reletiveFactorCalculatedBalances, &FactorCalculatedBalance{
				PeroidId: p.Id,
				UserId:   reletiveNode.To,
				Demand:   RelDemand,
				Debt:     RelDebt,
			})
		}
		// fmt.Printf("reletiveFactorCalculatedBalances: %v\n", lib.ToJsonIndent(reletiveFactorCalculatedBalances))
		listOfFactorCalculatedBalances = append(listOfFactorCalculatedBalances, &FactorCalculatedBalance{
			PeroidId:                         p.Id,
			UserId:                           node.To,
			Demand:                           Demand,
			Debt:                             Debt,
			ReletiveFactorCalculatedBalances: &reletiveFactorCalculatedBalances,
		})
		// fmt.Printf("listOfFactorCalculatedBalances: %v\n", lib.ToJsonIndent(listOfFactorCalculatedBalances))
	}
	return &listOfFactorCalculatedBalances
}

func convertFactorToExpense(factor Factor) core.Expense {
	expense := core.Expense{
		ExpenseID: factor.Id,
		Balances:  map[primitive.ObjectID]int{},
	}

	sumOfCoefficients := 0
	for _, user := range factor.Users {
		sumOfCoefficients += user.Coefficient
	}

	expense.Balances[factor.Buyer] = -factor.Price

	for _, user := range factor.Users {
		if user.UserId == factor.Buyer {
			expense.Balances[user.UserId] += (factor.Price / sumOfCoefficients) * user.Coefficient
			continue
		}
		expense.Balances[user.UserId] = (factor.Price / sumOfCoefficients) * user.Coefficient
	}

	return expense
}
