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
	sess.SetMode(mgo.Monotonic, true)
	sess.SetSafe(&mgo.Safe{})
	collection := sess.DB("test").C("newlead")
	f := fakeContactId(sess)
	fakeLead, err := NewLead(
		collection,
		&mgo.DBRef{
			Collection: "newcontact",
			Id:         f,
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

	var refContact contact
	err = sess.FindRef(fakeLead.Contact).One(&refContact)
	if err != nil {
		t.Errorf("%s", err)
	}
	fmt.Printf("%+v\n", refContact)

	err = collection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
	// Drop collection created by fakeContactId()
	err = sess.DB("test").C("newcontact").DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func fakeContactId(s *mgo.Session) bson.ObjectId {
	sess := s.Copy()
	defer sess.Close()
	sess.SetSafe(&mgo.Safe{})
	collection := sess.DB("test").C("newcontact")
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

func TestGetLead(t *testing.T) {
	sess, err := mgo.Dial("localhost")
	if err != nil {
		t.Errorf("%s", err)
	}
	defer sess.Close()
	sess.SetMode(mgo.Monotonic, true)
	sess.SetSafe(&mgo.Safe{})
	collection := sess.DB("test").C("newlead")
	f := fakeContactId(sess)
	fakeLead, err := NewLead(
		collection,
		&mgo.DBRef{
			Collection: "newcontact",
			Id:         f,
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
	id := fakeLead.Id
	fetchedLead, err := GetLead(id)
	if err != nil {
		t.Errorf("%s", err)
	}
	if fetchedLead.Id != fakeLead.Id {
		t.Errorf("Expected id of %v, but got %v", fakeLead.Id, fetchedLead.Id)
	}
	err = collection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
	// Drop collection created by fakeContactId()
	err = sess.DB("test").C("newcontact").DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestDeleteLead(t *testing.T) {
	sess, err := mgo.Dial("localhost")
	if err != nil {
		t.Errorf("%s", err)
	}
	defer sess.Close()
	sess.SetMode(mgo.Monotonic, true)
	sess.SetSafe(&mgo.Safe{})
	collection := sess.DB("test").C("newlead")
	f := fakeContactId(sess)
	fakeLead, err := NewLead(
		collection,
		&mgo.DBRef{
			Collection: "newcontact",
			Id:         f,
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
	id := fakeLead.Id
	err = DeleteLead(id)
	if err != nil {
		t.Errorf("%s", err)
	}
	n, err := collection.Count()
	if err != nil {
		t.Errorf("%s", err)
	}
	if n != 0 {
		t.Errorf("expected 0 documents in the collection, but found %d", n)
	}
	err = collection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
	contactCollection := sess.DB("test").C("newcontact")
	n, err = contactCollection.Count()
	if err != nil {
		t.Errorf("%s", err)
	}
	if n != 1 {
		t.Errorf("expected 1 document in the collection, but found %d", n)
	}
	err = contactCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}
