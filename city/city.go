package city

import (
	"github.com/nnqq/scr-parser/htmlfinder"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type City struct {
	ID    primitive.ObjectID        `bson:"_id,omitempty"`
	Title htmlfinder.NormalCaseCity `bson:"t,omitempty"`
	Slug  string                    `bson:"s,omitempty"`
}
