package city

import (
	"github.com/leaq-ru/parser/htmlfinder"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type City struct {
	ID    primitive.ObjectID        `bson:"_id,omitempty"`
	Title htmlfinder.NormalCaseCity `bson:"t,omitempty"`
	Slug  string                    `bson:"s,omitempty"`
}
