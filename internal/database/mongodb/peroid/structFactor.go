package peroid

import "go.mongodb.org/mongo-driver/bson/primitive"

type Factor struct {
	Id        primitive.ObjectID    `bson:"_id" json:"id"`
	Title     string                `bson:"title" json:"title"`
	UserId    primitive.ObjectID    `bson:"user_id" json:"user_id"`
	Price     uint64                `bson:"price" json:"price"`
	Buyer     primitive.ObjectID    `bson:"buyer" json:"buyer"`
	BuyerName *string               `bson:"buyer_name,omitempty" json:"buyer_name,omitempty"`
	Users     []UserWithCoefficient `bson:"users" json:"users"`
	PeroidId  primitive.ObjectID    `bson:"peroid_id" json:"peroid_id"`
}

type UserWithCoefficient struct {
	UserId      primitive.ObjectID `bson:"user_id" json:"user_id"`
	Coefficient uint64             `bson:"coefficient" json:"coefficient"`
}

func NewFactor(
	title string,
	userId primitive.ObjectID,
	price uint64,
	buyer primitive.ObjectID,
	users []UserWithCoefficient,
	peroidId primitive.ObjectID,
) *Factor {
	return &Factor{
		Title:    title,
		UserId:   userId,
		Price:    price,
		Buyer:    buyer,
		Users:    users,
		PeroidId: peroidId,
	}
}

func (f *Factor) GenerateId() *Factor {
	f.Id = primitive.NewObjectID()
	return f
}
