package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type City struct {
	ID       primitive.ObjectID `bson:"_id"`
	VkCityID int                `bson:"v"`
	Title    *title             `bson:"t"`
}

type title struct {
	Ru string `bson:"r"`
	En string `bson:"e"`
}

func Create() {

}
