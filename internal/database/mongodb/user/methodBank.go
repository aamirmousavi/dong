package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (hand *UserHandler) UpdateBank(
	b *Bank,
) error {
	_, err := hand.bank.UpdateOne(
		context.TODO(),
		bson.M{
			"_id": b.Id,
		},
		bson.M{
			"$set": b,
		},
		options.Update().SetUpsert(true),
	)
	return err
}

func (hand *UserHandler) GetBank(
	id primitive.ObjectID,
) (*Bank, error) {
	bank := new(Bank)
	if err := hand.bank.FindOne(
		context.TODO(),
		bson.M{
			"_id": id,
		},
	).Decode(&bank); err != nil {
		return nil, err
	}
	return bank, nil
}
