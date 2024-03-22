package user

import "go.mongodb.org/mongo-driver/mongo"

type UserHandler struct {
	database *mongo.Database
	user     *mongo.Collection
	token    *mongo.Collection
}

func NewHandler(
	database *mongo.Database,
) *UserHandler {
	return &UserHandler{
		database,
		database.Collection(_COLLECTION_USER),
		database.Collection(_COLLECTION_TOKEN),
	}
}
