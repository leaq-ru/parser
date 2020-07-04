package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type People struct {
	ID         primitive.ObjectID   `bson:"_id"`
	CompanyIDs []primitive.ObjectID `bson:"c,omitempty"`
	Name       string               `bson:"n,omitempty"`
	Avatar     avatar               `bson:"a,omitempty"`
	Phone      int                  `bson:"p,omitempty"`
	Email      string               `bson:"e,omitempty"`
	VkID       int                  `bson:"v,omitempty"`
}

func (p People) validate() error {
	return validation.ValidateStruct(
		&p,
		validation.Field(&p.ID, validation.Required),
		validation.Field(&p.CompanyIDs, validation.Required),
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Avatar, validation.Required),
	)
}
