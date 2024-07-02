package balance

import "go.mongodb.org/mongo-driver/bson/primitive"

type BalanceList []*Balance

type Balance struct {
	Id           primitive.ObjectID  `bson:"_id" json:"id"`
	PeroidId     *primitive.ObjectID `bson:"peroid_id,omitempty" json:"peroid_id,omitempty"`
	SourceUserId primitive.ObjectID  `bson:"source_user_id" json:"source_user_id"`
	TargetUserId primitive.ObjectID  `bson:"target_user_id" json:"target_user_id"`
	Amount       int                 `bson:"amount" json:"amount"`
	IsPayment    bool                `bson:"is_payment" json:"is_payment"`
}

func NewBalance(
	peroidId *primitive.ObjectID,
	sourceUserId primitive.ObjectID,
	targetUserId primitive.ObjectID,
	amount int,
	isPayment bool,
) *Balance {
	return &Balance{
		Id:           primitive.NewObjectID(),
		PeroidId:     peroidId,
		SourceUserId: sourceUserId,
		TargetUserId: targetUserId,
		Amount:       amount,
		IsPayment:    isPayment,
	}
}
