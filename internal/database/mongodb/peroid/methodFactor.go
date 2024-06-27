package peroid

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (hand *PeroidHandler) FactorAdd(
	factor *Factor,
) error {
	_, err := hand.factor.InsertOne(context.Background(), factor)
	if err != nil {
		return err
	}
	return nil
}

func (hand *PeroidHandler) FactorGet(
	id primitive.ObjectID,
) (*Factor, error) {
	factor := &Factor{}
	err := hand.factor.FindOne(context.Background(), bson.M{"_id": id}).Decode(factor)
	if err != nil {
		return nil, err
	}
	return factor, nil
}

func (hand *PeroidHandler) FactorList(peroidId primitive.ObjectID) ([]*Factor, error) {
	cursor, err := hand.factor.Find(context.Background(), bson.M{"peroid_id": peroidId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	list := []*Factor{}
	for cursor.Next(context.Background()) {
		factor := &Factor{}
		if err := cursor.Decode(factor); err != nil {
			return nil, err
		}
		list = append(list, factor)
	}
	return list, nil
}
