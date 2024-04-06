package contact

import "go.mongodb.org/mongo-driver/mongo"

type ContactHandler struct {
	collection *mongo.Collection
}

func NewHandler(
	collection *mongo.Collection,
) *ContactHandler {
	return &ContactHandler{
		collection,
	}
}
