package otp

import "go.mongodb.org/mongo-driver/mongo"

type OTPHandler struct {
	collection *mongo.Collection
}

func NewHandler(
	collection *mongo.Collection,
) *OTPHandler {
	return &OTPHandler{
		collection,
	}
}
