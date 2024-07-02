package peroid

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (hand *PeroidHandler) FactorCalculatedBalanceAdd(
	peroidId *primitive.ObjectID,
	factorCalculatedBalance *FactorCalculatedBalanceList,
) error {
	inters := make([]interface{}, len(*factorCalculatedBalance))
	for i, v := range *factorCalculatedBalance {
		inters[i] = v
	}
	if peroidId != nil {
		if _, err := hand.factorCalculatedBalance.DeleteMany(context.TODO(), bson.M{"peroid_id": peroidId}); err != nil {
			return err
		}
	}
	_, err := hand.factorCalculatedBalance.InsertMany(context.Background(), inters)
	if err != nil {
		return err
	}
	return nil
}

func (hand *PeroidHandler) FactorCalculatedBalanceGet(
	PeroidId primitive.ObjectID,
) (*FactorCalculatedBalanceList, error) {
	cursor, err := hand.factorCalculatedBalance.Find(
		context.Background(),
		bson.M{"peroid_id": PeroidId},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var list FactorCalculatedBalanceList
	if err = cursor.All(context.Background(), &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (hand *PeroidHandler) FactorCalculatedBalanceGetByUser(
	PeroidId primitive.ObjectID,
	UserId primitive.ObjectID,
) (*FactorCalculatedBalance, error) {
	var factorCalculatedBalance FactorCalculatedBalance
	err := hand.factorCalculatedBalance.FindOne(
		context.Background(),
		bson.M{"peroid_id": PeroidId, "user_id": UserId},
	).Decode(&factorCalculatedBalance)
	if err != nil {
		return nil, err
	}
	return &factorCalculatedBalance, nil
}
