package user

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const _EXPIRE_DAY = 30

var _secret = ""

func InitSetSecret(secret string) {
	_secret = secret
}

type Token struct {
	AccessToken          string             `bson:"access_token" json:"access_token"`
	CreatedAt            primitive.DateTime `bson:"created_at" json:"created_at"`
	AccessTokenExpiresIn primitive.DateTime `bson:"access_token_expires_in" json:"access_token_expires_in"`
	UserId               primitive.ObjectID `bson:"user_id" json:"user_id"`
}

func newTokenFromUser(usr *User) *Token {
	return &Token{
		AccessToken:          makeSha512(_secret, "ACCESS"+time.Now().String()+usr.Id.Hex()+primitive.NewObjectID().Hex()),
		CreatedAt:            primitive.NewDateTimeFromTime(time.Now()),
		AccessTokenExpiresIn: primitive.NewDateTimeFromTime(time.Now().Add(time.Hour * 24 * _EXPIRE_DAY)),
		UserId:               usr.Id,
	}
}
func makeSha512(secret, data string) string {
	h := hmac.New(sha512.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
