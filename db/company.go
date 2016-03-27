package db

import (
	"gopkg.in/mgo.v2/bson"
)

// Company provides definition for model
// swagger:model company
type Company struct {
	ID      bson.ObjectId   `json:"id" bson:"_id"`
	Name    string          `json:"name" bson:"Name"`
	Owner   string          `json:"owner" bson:"Owner"`
	About   string          `json:"about" bson:"About"`
	City    string          `json:"city"  bson:"City"`
	Address string          `json:"address" bson:"Address"`
	Users   []bson.ObjectId `json:"users" bson:"Users"`
}
