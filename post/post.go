package post

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Post struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CompanyID primitive.ObjectID `bson:"c,omitempty"`
	Date      time.Time          `bson:"d,omitempty"`
	Text      string             `bson:"t,omitempty"`
	Photos    []Photo            `bson:"p,omitempty"`
}

type Photo struct {
	URLm string `bson:"um,omitempty"`
	URLr string `bson:"ur,omitempty"`
}
