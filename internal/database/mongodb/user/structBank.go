package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bank struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	CardNumber  string             `bson:"card_number" json:"card_number"`
	BankName    string             `bson:"bank_name" json:"bank_name"`
	AccountName string             `bson:"account_name" json:"account_name"`
	ShebaNumber string             `bson:"sheba_number" json:"sheba_number"`
}

func NewBank(
	userId primitive.ObjectID,
	cardNumber, bankName, accountName, shebaNumber string,
) *Bank {
	return &Bank{
		Id:          userId,
		CardNumber:  cardNumber,
		BankName:    bankName,
		AccountName: accountName,
		ShebaNumber: shebaNumber,
	}
}
