package peroid

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (hand *PeroidHandler) FactorUpdate(factor *Factor) error {
	_, err := hand.factor.ReplaceOne(context.Background(), bson.M{"_id": factor.Id}, factor)
	if err != nil {
		return err
	}
	return nil
}

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

func (hand *PeroidHandler) FactorListWithContact(
	userId primitive.ObjectID,
	contactUserId primitive.ObjectID,
) ([]*Factor, error) {
	cursor, err := hand.factor.Find(
		context.Background(),
		bson.M{"$or": []bson.M{
			{"buyer": userId, "users.user_id": bson.M{"$in": []primitive.ObjectID{contactUserId}}},
			{"buyer": contactUserId, "users.user_id": bson.M{"$in": []primitive.ObjectID{userId}}},
		}},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var list []*Factor
	if err := cursor.All(context.Background(), &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (hand *PeroidHandler) FactorListByUser(userId primitive.ObjectID) ([]*Factor, error) {
	cursor, err := hand.factor.Find(
		context.TODO(),
		bson.M{
			"user_id": userId,
		},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var list []*Factor
	if err := cursor.All(context.Background(), &list); err != nil {
		return nil, err
	}
	return list, nil
}
