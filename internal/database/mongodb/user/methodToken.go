package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (hand *UserHandler) CreateToken(
	ctx context.Context,
	tkn *Token,
) error {
	_, err := hand.token.InsertOne(
		ctx,
		tkn,
	)
	return err
}

func (hand *UserHandler) Logout(
	ctx context.Context,
	token string,
) error {
	_, err := hand.token.DeleteOne(
		ctx,
		bson.M{
			"access_token": token,
		},
	)
	return err
}

func (hand *UserHandler) GetUserByToken(
	ctx context.Context,
	token string,
) (*User, error) {
	cursor, err := hand.token.Aggregate(
		ctx,
		[]bson.M{
			{
				"$match": bson.M{
					"access_token": token,
				},
			},
			{
				"$lookup": bson.M{
					"from":         "user",
					"localField":   "user_id",
					"foreignField": "_id",
					"as":           "user",
				},
			},
			{
				"$unwind": "$user",
			},
			{
				"$project": bson.M{
					"_id":        "$user._id",
					"first_name": "$user.first_name",
					"last_name":  "$user.last_name",
					"number":     "$user.number",
					"pic":        "$user.pic",
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var usr User
	if cursor.Next(ctx) {
		err = cursor.Decode(&usr)
		if err != nil {
			return nil, err
		}
	}
	return &usr, nil
}
