package posts

import "go.mongodb.org/mongo-driver/bson/primitive"

type post struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CompanyID primitive.ObjectID `bson:"c,omitempty"`
}
