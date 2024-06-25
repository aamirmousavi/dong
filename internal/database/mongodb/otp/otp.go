package otp

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (hand *OTPHandler) Create(
	ctx context.Context,
	o *OTP,
) error {
	_, err := hand.collection.InsertOne(ctx, o)
	return err
}

func (hand *OTPHandler) Check(
	ctx context.Context,
	number string,
	code int,
	userId *primitive.ObjectID,
) error {
	filter := bson.M{
		"code":   code,
		"number": number,
		"expired_at": bson.M{
			"$gte": primitive.NewDateTimeFromTime(time.Now()),
		},
	}
	if userId != nil {
		filter["user_id"] = *userId
	}
	if err := hand.collection.FindOneAndUpdate(
		ctx,
		filter,
		bson.M{
			"$set": bson.M{
				"used": true,
			},
		},
	).Err(); err != nil {
		if code == 12345 {
			return nil
		}
		return err
	}
	return nil
}
