package peroid

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (hand *PeroidHandler) Add(
	peroid *Peroid,
) error {
	_, err := hand.peroid.InsertOne(context.Background(), peroid)
	if err != nil {
		return err
	}
	return nil
}

func (hand *PeroidHandler) Get(
	id primitive.ObjectID,
) (*Peroid, error) {
	peroid := &Peroid{}
	err := hand.peroid.FindOne(context.Background(), bson.M{"_id": id}).Decode(peroid)
	if err != nil {
		return nil, err
	}
	return peroid, nil
}

func (hand *PeroidHandler) GetByUserId(
	userId primitive.ObjectID,
) ([]*Peroid, error) {
	cursor, err := hand.peroid.Find(context.Background(), bson.M{"$or": []bson.M{{"user_id": userId}, {"user_ids": bson.M{"$in": []primitive.ObjectID{userId}}}}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	list := []*Peroid{}
	for cursor.Next(context.Background()) {
		peroid := &Peroid{}
		if err := cursor.Decode(peroid); err != nil {
			return nil, err
		}
		list = append(list, peroid)
	}
	return list, nil
}

func (hand *PeroidHandler) Remove(
	id primitive.ObjectID,
) error {
	_, err := hand.peroid.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}

func (hand *PeroidHandler) Update(
	peroid *Peroid,
) error {
	_, err := hand.peroid.ReplaceOne(context.Background(), bson.M{"_id": peroid.Id}, peroid)
	if err != nil {
		return err
	}
	return nil
}

func (hand *PeroidHandler) UpdateUserIds(
	id primitive.ObjectID,
	userIds []primitive.ObjectID,
) error {
	_, err := hand.peroid.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": bson.M{"user_ids": userIds}})
	if err != nil {
		return err
	}
	return nil
}

func (hand *PeroidHandler) AddUser(
	id primitive.ObjectID,
	userId primitive.ObjectID,
) error {
	_, err := hand.peroid.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{
			"$inc":  bson.M{"user_count": 1},
			"$push": bson.M{"user_ids": userId},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (hand *PeroidHandler) UpdateTotalPrice(
	id primitive.ObjectID,
	totalPrice uint64,
) error {
	_, err := hand.peroid.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": bson.M{"total_price": totalPrice}})
	if err != nil {
		return err
	}
	return nil
}

func (hand *PeroidHandler) UpdateAvgPrice(
	id primitive.ObjectID,
	avgPrice uint64,
) error {
	_, err := hand.peroid.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": bson.M{"avg_price": avgPrice}})
	if err != nil {
		return err
	}
	return nil
}

func (hand *PeroidHandler) UpdateTotalFactor(
	id primitive.ObjectID,
	totalFactor uint64,
) error {
	_, err := hand.peroid.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": bson.M{"total_factor": totalFactor}})
	if err != nil {
		return err
	}
	return nil
}

func (hand *PeroidHandler) UpdateAll(
	peroid *Peroid,
) error {
	_, err := hand.peroid.ReplaceOne(context.Background(), bson.M{"_id": peroid.Id}, peroid)
	if err != nil {
		return err
	}
	return nil
}
