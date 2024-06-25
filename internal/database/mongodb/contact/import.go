package contact

import "go.mongodb.org/mongo-driver/mongo"

type ContactHandler struct {
	database *mongo.Database
	contact  *mongo.Collection
}

func NewHandler(
	database *mongo.Database,
) *ContactHandler {
	return &ContactHandler{
		database,
		database.Collection(_COLLECTION_CONTACT),
	}
}
