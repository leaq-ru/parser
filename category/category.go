package category

import (
	"github.com/jbrukh/bayesian"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Title bayesian.Class     `bson:"t,omitempty"`
	Slug  string             `bson:"s,omitempty"`
}
