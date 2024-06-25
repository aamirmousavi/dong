package peroid

import "go.mongodb.org/mongo-driver/mongo"

type PeroidHandler struct {
	database *mongo.Database
	peroid   *mongo.Collection
}

func NewHandler(
	database *mongo.Database,
) *PeroidHandler {
	return &PeroidHandler{
		database,
		database.Collection(_COLLECTION_PEROID),
	}
}
