package contact

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (hand *ContactHandler) Create(
	ctx context.Context,
	cnt *Contact,
) error {
	_, err := hand.collection.InsertOne(
		ctx,
		cnt,
	)
	return err
}

func (hand *ContactHandler) Remove(
	ctx context.Context,
	id primitive.ObjectID,
	userId primitive.ObjectID,
) error {
	_, err := hand.collection.DeleteOne(
		ctx,
		bson.M{
			"_id":     id,
			"user_id": userId,
		},
	)
	return err
}

func (hand *ContactHandler) List(
	ctx context.Context,
	userId primitive.ObjectID,
	limit, offest int64,
) (*ContactList, error) {
	cursor, err := hand.collection.Find(
		ctx,
		bson.M{
			"user_id": userId,
		},
		options.Find().
			SetLimit(limit).
			SetSkip(offest),
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	list := make(ContactList, 0)
	if err := cursor.All(ctx, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (hand *ContactHandler) Get(
	ctx context.Context,
	id primitive.ObjectID,
	userId primitive.ObjectID,
) (*Contact, error) {
	cnt := new(Contact)
	err := hand.collection.FindOne(
		ctx,
		bson.M{
			"_id":     id,
			"user_id": userId,
		},
	).Decode(cnt)
	return cnt, err
}

func (hand *ContactHandler) Edit(
	ctx context.Context,
	cnt *Contact,
) error {
	_, err := hand.collection.UpdateByID(
		ctx,
		cnt.Id,
		bson.M{
			"$set": cnt,
		},
	)
	return err
}
