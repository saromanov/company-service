package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	company "github.com/saromanov/company-service/proto/company"
)

var (
	database string
	collection string
	session *mgo.Session
	coll *mgo.Collection
)

func Init(){
	var err error
	session, err = mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	// Collection People
	coll = session.DB("test").C("people")
}

func Create(user *company.CompanyItem) (*company.CreateResponse, error) {
	id := bson.NewObjectId()
	comp := Company {
		ID: id,
		Name: user.Name,
		Owner: user.Owner,
	}

	err := coll.Insert(comp)
	if err != nil {
		return nil, err
	}

	resp := &company.CreateResponse {
		Id: id.Hex(),
	}
	return resp, err
}

func Read(id string) (*company.CompanyItem, error) {
	var comp Company
	err := coll.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&comp)
	if err != nil {
		return nil, err
	}

	item := &company.CompanyItem {
		Name: comp.Name,
		Owner: comp.Owner,
	}
	return item, nil
}

// Remove by profileid and company id
func Remove(id string) error {
	if !bson.IsObjectIdHex(id) {
		return fmt.Errorf("%s is not a ObjectId value", id)
	}
	return coll.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}