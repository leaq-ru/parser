package city

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type City struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	VkCityID int                `bson:"v,omitempty"`
	Title    string             `bson:"t,omitempty"`
	Slug     string             `bson:"s,omitempty"`
}
