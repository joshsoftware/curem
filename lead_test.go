package main

import (
	"fmt"
	"testing"

	"log"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func TestNewLead(t *testing.T) {
	sess, err := mgo.Dial("localhost")
	if err != nil {
		t.Errorf("%s", err)
	}
	defer sess.Close()
	sess.SetSafe(&mgo.Safe{})
	collection := sess.DB("test").C("newlead")
	err = collection.DropCollection() //Fresh test DB collection
	if err != nil {
		t.Errorf("%s", err)
	}
	f := fakeContactId()
	fakeLead, err := NewLead(
		collection,
		&mgo.DBRef{
			Collection: "newcontact",
			Id:         f.Hex(),
			Database:   "test",
		},
		"Web",
		"Hari",
		"Warming Up",
		2.5,
		20,
		3,
		"25th June, 2014",
		[]string{"Call back", "Based in mumbai"},
	)
	if err != nil {
		t.Errorf("%s", err)
	}

	fmt.Printf("%+v\n", fakeLead)
}

func fakeContactId() bson.ObjectId {
	sess, err := mgo.Dial("localhost")
	if err != nil {
		log.Println(err)
	}
	defer sess.Close()
	sess.SetSafe(&mgo.Safe{})
	collection := sess.DB("test").C("newcontact")
	err = collection.DropCollection() //Fresh test DB collection
	if err != nil {
		log.Println(err)
	}
	fakeContact, err := NewContact(
		collection,
		"Encom Inc.",
		"Flynn",
		"flynn@encom.com",
		"",
		"",
		"USA",
	)
	if err != nil {
		log.Println(err)
	}
	return fakeContact.Id
}
