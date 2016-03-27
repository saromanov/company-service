package db

import (
	"gopkg.in/mgo.v2"
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
func Create(user *company.Company) error {

}