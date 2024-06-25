package balance

import "go.mongodb.org/mongo-driver/mongo"

type BalanceHandler struct {
	database *mongo.Database
	balance  *mongo.Collection
}

func NewHandler(
	database *mongo.Database,
) *BalanceHandler {
	return &BalanceHandler{
		database,
		database.Collection(_COLLECTION_BALANCE),
	}
}
