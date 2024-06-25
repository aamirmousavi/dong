package peroid

import "go.mongodb.org/mongo-driver/bson/primitive"

type Factor struct {
	Id     primitive.ObjectID    `bson:"_id" json:"id"`
	Title  string                `bson:"title" json:"title"`
	UserId primitive.ObjectID    `bson:"user_id" json:"user_id"`
	Price  uint64                `bson:"price" json:"price"`
	Buyer  primitive.ObjectID    `bson:"buyer" json:"buyer"`
	Users  []UserWithCoefficient `bson:"users" json:"users"`
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
) *Factor {
	return &Factor{
		Title:  title,
		UserId: userId,
		Price:  price,
		Buyer:  buyer,
		Users:  users,
	}
}

func (f *Factor) GenerateId() *Factor {
	f.Id = primitive.NewObjectID()
	return f
}