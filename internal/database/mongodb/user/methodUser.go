package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (hand *UserHandler) GetId(
	ctx context.Context,
	number string,
) (*primitive.ObjectID, error) {
	usr := new(struct {
		Id primitive.ObjectID `bson:"_id"`
	})
	if err := hand.user.FindOne(
		ctx,
		bson.M{
			"number": number,
		},
	).Decode(usr); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &usr.Id, nil
}
