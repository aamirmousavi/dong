package contact

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (hand *ContactHandler) Add(c *Contact) error {
	_, err := hand.contact.InsertOne(context.TODO(), c)
	return err
}

func (hand *ContactHandler) GetByContactOf(contactOf primitive.ObjectID) ([]*Contact, error) {
	cursor, err := hand.contact.Find(context.TODO(), bson.M{"contact_of": contactOf})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	list := []*Contact{}
	for cursor.Next(context.Background()) {
		c := &Contact{}
		if err := cursor.Decode(c); err != nil {
			return nil, err
		}
		list = append(list, c)
	}
	return list, nil
}

func (hand *ContactHandler) Remove(id primitive.ObjectID) error {
	_, err := hand.contact.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

func (hand *ContactHandler) Update(c *Contact) error {
	_, err := hand.contact.ReplaceOne(context.TODO(), bson.M{"_id": c.Id}, c)
	return err
}

func (hand *ContactHandler) Get(id primitive.ObjectID) (*Contact, error) {
	c := &Contact{}
	err := hand.contact.FindOne(context.TODO(), bson.M{"_id": id}).Decode(c)
	return c, err
}
