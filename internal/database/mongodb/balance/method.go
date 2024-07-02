package balance

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (hand *BalanceHandler) Add(
	peroidId *primitive.ObjectID,
	balance BalanceList,
) error {
	interfaces := make([]interface{}, len(balance))
	for i, b := range balance {
		interfaces[i] = b
	}
	if peroidId != nil {
		if _, err := hand.balance.DeleteMany(context.TODO(), bson.M{
			"peroid_id":  peroidId,
			"is_payment": false,
		}); err != nil {
			return err
		}
	}
	_, err := hand.balance.InsertMany(
		context.TODO(),
		interfaces,
	)
	return err
}
