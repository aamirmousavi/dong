package peroid

import "go.mongodb.org/mongo-driver/bson/primitive"

type FactorCalculatedBalanceList []*FactorCalculatedBalance

func (l *FactorCalculatedBalanceList) Find(userId primitive.ObjectID) (*FactorCalculatedBalance, bool) {
	for _, v := range *l {
		if v.UserId == userId {
			return v, true
		}
	}
	return nil, false
}

type FactorCalculatedBalance struct {
	PeroidId                         primitive.ObjectID           `bson:"peroid_id" json:"peroid_id"`
	UserId                           primitive.ObjectID           `bson:"user_id" json:"user_id"`
	Demand                           *int                         `bson:"demand" json:"demand"`
	Debt                             *int                         `bson:"debt" json:"debt"`
	ReletiveFactorCalculatedBalances *FactorCalculatedBalanceList `bson:"reletive_factor_calculated_balances" json:"reletive_factor_calculated_balances"`
}
