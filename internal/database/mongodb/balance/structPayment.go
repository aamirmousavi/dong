package balance

import "go.mongodb.org/mongo-driver/bson/primitive"

type PaymentList []*Payment

type Payment struct {
	Id           primitive.ObjectID  `bson:"_id" json:"id"`
	Title        string              `bson:"title" json:"title"`
	PeroidId     *primitive.ObjectID `bson:"peroid_id,omitempty" json:"peroid_id,omitempty"`
	SourceUserId primitive.ObjectID  `bson:"source_user_id" json:"source_user_id"`
	TargetUserId primitive.ObjectID  `bson:"target_user_id" json:"target_user_id"`
	Amount       uint64              `bson:"amount" json:"amount"`
}

func NewPayment(
	title string,
	peroidId *primitive.ObjectID,
	sourceUserId primitive.ObjectID,
	targetUserId primitive.ObjectID,
	amount uint64,
) *Payment {
	return &Payment{
		Title:        title,
		PeroidId:     peroidId,
		SourceUserId: sourceUserId,
		TargetUserId: targetUserId,
		Amount:       amount,
	}
}

func (p *Payment) SetId(id primitive.ObjectID) *Payment {
	p.Id = id
	return p
}

func (p *Payment) GenerateId() *Payment {
	p.Id = primitive.NewObjectID()
	return p
}
