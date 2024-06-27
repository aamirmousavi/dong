package balance

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (hand *BalanceHandler) PaymentUpdate(payment *Payment) error {
	_, err := hand.payment.ReplaceOne(
		context.Background(),
		bson.M{"_id": payment.Id},
		payment,
	)
	return err
}

func (hand *BalanceHandler) PaymentAdd(payment ...*Payment) error {
	interfaces := make([]interface{}, len(payment))
	for i, p := range payment {
		interfaces[i] = p
	}
	_, err := hand.payment.InsertMany(
		context.Background(),
		interfaces,
	)
	return err
}

func (hand *BalanceHandler) PaymentGet(id primitive.ObjectID) (*Payment, error) {
	var payment Payment
	if err := hand.payment.FindOne(
		context.Background(),
		bson.M{"_id": id},
	).Decode(&payment); err != nil {
		return nil, err
	}
	return &payment, nil
}

func (hand *BalanceHandler) PaymentList(
	peroidId *primitive.ObjectID,
	userId *primitive.ObjectID,
) (PaymentList, error) {
	fitler := bson.M{}
	if peroidId != nil {
		fitler["peroid_id"] = peroidId
	}
	if userId != nil {
		fitler["$or"] = []bson.M{
			{"source_user_id": userId},
			{"target_user_id": userId},
		}
	}
	cursor, err := hand.payment.Find(
		context.Background(),
		fitler,
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var list PaymentList
	if err := cursor.All(context.Background(), &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (hand *BalanceHandler) PaymentListWithContact(
	userId primitive.ObjectID,
	contactUserId primitive.ObjectID,
) (PaymentList, error) {
	cursor, err := hand.payment.Find(
		context.Background(),
		bson.M{"$or": []bson.M{
			{"source_user_id": userId, "target_user_id": contactUserId},
			{"source_user_id": contactUserId, "target_user_id": userId},
		}},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var list PaymentList
	if err := cursor.All(context.Background(), &list); err != nil {
		return nil, err
	}
	return list, nil
}
