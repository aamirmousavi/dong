package contact

import "go.mongodb.org/mongo-driver/bson/primitive"

type Contact struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	Number    string             `bson:"number" json:"number"`
	FirstName string             `bson:"first_name" json:"first_name"`
	LastName  string             `bson:"last_name" json:"last_name"`
	Pic       *string            `bson:"pic" json:"pic"`
	UserId    primitive.ObjectID `bson:"user_id" json:"user_id"`
	ContactOf primitive.ObjectID `bson:"contact_of" json:"contact_of"`

	Demand *uint64 `bson:"demand,omitempty" json:"demand,omitempty"`
	Debt   *uint64 `bson:"debt,omitempty" json:"debt,omitempty"`
}

func NewContact(
	number, firstName, lastName string,
	pic *string,
	userId primitive.ObjectID,
) *Contact {
	return &Contact{
		Number:    number,
		FirstName: firstName,
		LastName:  lastName,
		Pic:       pic,
		UserId:    userId,
	}
}

func (c *Contact) SetId(id primitive.ObjectID) *Contact {
	c.Id = id
	return c
}

func (c *Contact) SetUserId(userId primitive.ObjectID) *Contact {
	c.UserId = userId
	return c
}

func (c *Contact) GenerateId() *Contact {
	c.Id = primitive.NewObjectID()
	return c
}
