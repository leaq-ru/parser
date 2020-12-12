package cached_list

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type cachedList struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	MD5       string             `bson:"m,omitempty"`
	URL       string             `bson:"u,omitempty"`
	Kind      kind               `bson:"k,omitempty"`
	CreatedAt time.Time          `bson:"ca,omitempty"`
}

type kind uint8

const (
	_ kind = iota
	Kind_email
	Kind_phone
)
