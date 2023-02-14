package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Car struct {
	ID    int `bson:"_id" json:"id"`
	Seats int `bson:"seats" json:"seats"`
}

type Journey struct {
	ID        int                `bson:"_id" json:"id"`
	People    int                `bson:"people" json:"people"`
	CreatedOn primitive.DateTime `bson:"createdOn"`
	CarID     string             `bson:"carID"`
}
