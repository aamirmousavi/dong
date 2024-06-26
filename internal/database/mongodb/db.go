package mongodb

import (
	"context"

	mongodb_balance "github.com/aamirmousavi/dong/internal/database/mongodb/balance"
	mongodb_contact "github.com/aamirmousavi/dong/internal/database/mongodb/contact"
	mongodb_otp "github.com/aamirmousavi/dong/internal/database/mongodb/otp"
	mongodb_peroid "github.com/aamirmousavi/dong/internal/database/mongodb/peroid"
	mongodb_user "github.com/aamirmousavi/dong/internal/database/mongodb/user"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Handler struct {
	*mongodb_balance.BalanceHandler
	*mongodb_otp.OTPHandler
	*mongodb_user.UserHandler
	*mongodb_contact.ContactHandler
	*mongodb_peroid.PeroidHandler
}

func NewHandler(addr string) (*Handler, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(addr))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		return nil, err
	}
	return &Handler{
		mongodb_balance.NewHandler(
			client.Database(mongodb_balance.DATABASE_BALANCE),
		),
		mongodb_otp.NewHandler(
			client.Database(mongodb_user.DATABASE_USER).
				Collection(mongodb_otp.COLLECTION_OTP),
		),
		mongodb_user.NewHandler(
			client.Database(mongodb_user.DATABASE_USER),
		),
		mongodb_contact.NewHandler(
			client.Database(mongodb_contact.DATABASE_CONTACT),
		),
		mongodb_peroid.NewHandler(
			client.Database(mongodb_peroid.DATABASE_PEROID),
		),
	}, nil
}
