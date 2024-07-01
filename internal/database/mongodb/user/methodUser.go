package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (hand *UserHandler) UserExists(
	ctx context.Context,
	number string,
) (bool, error) {
	count, err := hand.user.CountDocuments(
		ctx,
		bson.M{
			"number": number,
		},
	)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (hand *UserHandler) Create(
	ctx context.Context,
	usr *User,
) error {
	_, err := hand.user.InsertOne(
		ctx,
		usr,
	)
	return err
}

func (hand *UserHandler) GetById(
	ctx context.Context,
	id primitive.ObjectID,
) (*User, error) {
	usr := new(User)
	err := hand.user.FindOne(
		ctx,
		bson.M{
			"_id": id,
		},
	).Decode(usr)
	if err != nil {
		return usr, nil
	}
	return usr, err
}

func (hand *UserHandler) Get(
	ctx context.Context,
	number string,
) (*User, error) {
	usr := new(User)
	err := hand.user.FindOne(
		ctx,
		bson.M{
			"number": number,
		},
	).Decode(usr)
	return usr, err
}

func (hand *UserHandler) GetMany(ctx context.Context, ids []primitive.ObjectID) ([]*User, error) {
	cursor, err := hand.user.Find(ctx, bson.M{"_id": bson.M{"$in": ids}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var users []*User
	for cursor.Next(ctx) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (hand *UserHandler) Update(
	ctx context.Context,
	usr *User,
) error {
	_, err := hand.user.UpdateOne(
		ctx,
		bson.M{
			"number": usr.Number,
			"_id":    usr.Id,
		},
		bson.M{
			"$set": usr,
		},
	)
	return err
}
