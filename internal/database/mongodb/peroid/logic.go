package peroid

import (
	"github.com/aamirmousavi/dong/internal/splitwise"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (p *Peroid) EditFactor(factor *Factor, oldTata *Factor, isPayment bool) {
	userWithCoefficient := make([]splitwise.UserWithCoefficient[primitive.ObjectID], 0)
	for _, user := range factor.Users {
		userWithCoefficient = append(userWithCoefficient, splitwise.UserWithCoefficient[primitive.ObjectID]{
			UserId:      user.UserId,
			Coefficient: user.Coefficient,
		})
	}
	Factor := splitwise.Factor[primitive.ObjectID]{
		Id:    factor.Id,
		Buyer: factor.Buyer,
		Price: factor.Price,
		Users: userWithCoefficient,
	}
	p.PeroidSplitwise.EditFactors(Factor)

	FactorCalculatedBalances := p.PeroidSplitwise.Transactions
	dbFactorCalculatedBalanceList := make(FactorCalculatedBalanceList, 0)
	for _, factorCalculatedBalance := range *FactorCalculatedBalances {

		ReletiveFactorCalculatedBalances := make(FactorCalculatedBalanceList, 0)
		if factorCalculatedBalance.ReletiveTransaction != nil {
			for _, reletiveFactorCalculatedBalance := range *factorCalculatedBalance.ReletiveTransaction {
				ReletiveFactorCalculatedBalances = append(ReletiveFactorCalculatedBalances, &FactorCalculatedBalance{
					UserId: reletiveFactorCalculatedBalance.UserId,
					Demand: reletiveFactorCalculatedBalance.Demand,
					Debt:   reletiveFactorCalculatedBalance.Debt,
				})
			}
		}

		dbFactorCalculatedBalanceList = append(dbFactorCalculatedBalanceList, &FactorCalculatedBalance{
			PeroidId:                         p.Id,
			UserId:                           factorCalculatedBalance.UserId,
			Demand:                           factorCalculatedBalance.Demand,
			Debt:                             factorCalculatedBalance.Debt,
			ReletiveFactorCalculatedBalances: &ReletiveFactorCalculatedBalances,
		})
	}
	p.Balances = &dbFactorCalculatedBalanceList

	if isPayment {
		return
	}
	if p.MoneySpend == nil {
		p.MoneySpend = make(map[primitive.ObjectID]int)
	}
	if _, ok := p.MoneySpend[factor.Buyer]; !ok {
		p.MoneySpend[factor.Buyer] = 0
	}

	if oldTata.Price != factor.Price {
		p.MoneySpend[factor.Buyer] -= oldTata.Price
		p.MoneySpend[factor.Buyer] += factor.Price
		p.TotalPrice -= oldTata.Price
		p.TotalPrice += factor.Price
		p.AvgPrice = p.TotalPrice / p.TotalFactor
	}
}

func (p *Peroid) AddFactor(factor *Factor, isPayment bool) {
	userWithCoefficient := make([]splitwise.UserWithCoefficient[primitive.ObjectID], 0)
	for _, user := range factor.Users {
		userWithCoefficient = append(userWithCoefficient, splitwise.UserWithCoefficient[primitive.ObjectID]{
			UserId:      user.UserId,
			Coefficient: user.Coefficient,
		})
	}
	Factor := splitwise.Factor[primitive.ObjectID]{
		Id:    factor.Id,
		Buyer: factor.Buyer,
		Price: factor.Price,
		Users: userWithCoefficient,
	}
	p.PeroidSplitwise.AddFactors(Factor)

	FactorCalculatedBalances := p.PeroidSplitwise.Transactions
	dbFactorCalculatedBalanceList := make(FactorCalculatedBalanceList, 0)
	for _, factorCalculatedBalance := range *FactorCalculatedBalances {

		ReletiveFactorCalculatedBalances := make(FactorCalculatedBalanceList, 0)
		if factorCalculatedBalance.ReletiveTransaction != nil {
			for _, reletiveFactorCalculatedBalance := range *factorCalculatedBalance.ReletiveTransaction {
				ReletiveFactorCalculatedBalances = append(ReletiveFactorCalculatedBalances, &FactorCalculatedBalance{
					UserId: reletiveFactorCalculatedBalance.UserId,
					Demand: reletiveFactorCalculatedBalance.Demand,
					Debt:   reletiveFactorCalculatedBalance.Debt,
				})
			}
		}

		dbFactorCalculatedBalanceList = append(dbFactorCalculatedBalanceList, &FactorCalculatedBalance{
			PeroidId:                         p.Id,
			UserId:                           factorCalculatedBalance.UserId,
			Demand:                           factorCalculatedBalance.Demand,
			Debt:                             factorCalculatedBalance.Debt,
			ReletiveFactorCalculatedBalances: &ReletiveFactorCalculatedBalances,
		})
	}
	p.Balances = &dbFactorCalculatedBalanceList

	if isPayment {
		return
	}

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
