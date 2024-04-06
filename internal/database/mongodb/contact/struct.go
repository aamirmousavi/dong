package contact

import "go.mongodb.org/mongo-driver/bson/primitive"

type ContactList []*Contact

type Contact struct {
	Id            primitive.ObjectID  `bson:"_id" json:"id"`
	UserId        primitive.ObjectID  `bson:"user_id" json:"user_id"`
	ContactUserId *primitive.ObjectID `bson:"contact_user_id,omitempty" json:"contact_user_id,omitempty"`
	FirstName     *string             `bson:"first_name,omitempty" json:"first_name,omitempty"`
	LastName      *string             `bson:"last_name,omitempty" json:"last_name,omitempty"`
	Number        *string             `bson:"number,omitempty" json:"number,omitempty"`
	Pic           *string             `bson:"pic,omitempty" json:"pic,omitempty"`
}

func New(
	userId primitive.ObjectID,
	contactUserId *primitive.ObjectID,
	firstName, lastName, number, pic *string,
) *Contact {
	return &Contact{
		UserId:        userId,
		ContactUserId: contactUserId,
		FirstName:     firstName,
		LastName:      lastName,
		Number:        number,
		Pic:           pic,
	}
}

func (c *Contact) SetId(id primitive.ObjectID) *Contact {
	c.Id = id
	return c
}

func (c *Contact) GenerateId() *Contact {
	c.Id = primitive.NewObjectID()
	return c
}
