package otp

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OTP struct {
	UserId    *primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	Number    string              `bson:"number" json:"number"`
	Code      int                 `bson:"code" json:"code"`
	Used      bool                `bson:"used" json:"used"`
	ExpiredAt primitive.DateTime  `json:"expired_at" bson:"expired_at"`
}

func New(
	number string,
	code int,
) *OTP {
	return &OTP{
		Number:    number,
		Code:      code,
		Used:      false,
		ExpiredAt: primitive.NewDateTimeFromTime(time.Now().Add(_EXPIRE_TIME)),
	}
}

func (o *OTP) SetUserId(id primitive.ObjectID) *OTP {
	o.UserId = &id
	return o
}
