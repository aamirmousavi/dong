package peroid

import "go.mongodb.org/mongo-driver/bson/primitive"

type Peroid struct {
	Id      primitive.ObjectID   `bson:"_id" json:"id"`
	UserId  primitive.ObjectID   `bson:"user_id" json:"user_id"`
	Title   string               `bson:"title" json:"title"`
	UserIds []primitive.ObjectID `bson:"user_ids" json:"user_ids"`

	UserCount   uint64 `bson:"user_count" json:"user_count"`
	TotalPrice  uint64 `bson:"total_price" json:"total_price"`
	AvgPrice    uint64 `bson:"avg_price" json:"avg_price"`
	TotalFactor uint64 `bson:"total_factor" json:"total_factor"`
}

func NewPeroid(
	userId primitive.ObjectID,
	title string,
	userIds []primitive.ObjectID,
) *Peroid {
	return &Peroid{
		UserId:  userId,
		Title:   title,
		UserIds: userIds,
	}
}

func (p *Peroid) SetId(id primitive.ObjectID) *Peroid {
	p.Id = id
	return p
}

func (p *Peroid) GenerateId() *Peroid {
	p.Id = primitive.NewObjectID()
	return p
}
