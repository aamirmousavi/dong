package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	FirstName string             `bson:"first_name" json:"first_name"`
	LastName  string             `bson:"last_name" json:"last_name"`
	Number    string             `bson:"number" json:"number"`
	Pic       string             `bson:"pic" json:"pic"`
}

func NewUser(
	firstName, lastName, number, pic string,
) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Number:    number,
		Pic:       pic,
	}
}

func (u *User) SetId(id primitive.ObjectID) *User {
	u.Id = id
	return u
}

func (u *User) GenerateId() *User {
	u.Id = primitive.NewObjectID()
	return u
}

func (u *User) GenerateToken() *Token {
	return newTokenFromUser(u)
}
