package balance

import "go.mongodb.org/mongo-driver/bson/primitive"

type BalanceList []*Balance

type Balance struct {
	Id           primitive.ObjectID `bson:"_id" json:"id"`
	SourceUserId primitive.ObjectID `bson:"source_user_id" json:"source_user_id"`
	TargetUserId primitive.ObjectID `bson:"target_user_id" json:"target_user_id"`
	Amount       uint64             `bson:"amount" json:"amount"`
}

func NewBalance(
	sourceUserId primitive.ObjectID,
	targetUserId primitive.ObjectID,
	amount uint64,
) *Balance {
	return &Balance{
		Id:           primitive.NewObjectID(),
		SourceUserId: sourceUserId,
		TargetUserId: targetUserId,
		Amount:       amount,
	}
}
