package main

import (
	"testing"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func TestNewContact(t *testing.T) {
	sess, err := mgo.Dial("localhost")
	if err != nil {
		t.Errorf("%s", err)
	}
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
		t.Errorf("%s", err)
	}
	var fetchedContact contact
	err = sess.DB("test").C("newcontact").Find(bson.M{}).One(&fetchedContact)
	if err != nil {
		t.Errorf("%s", err)
	}

	// fakeContact is a pointer, because NewContact returns a pointer to a struct of contact type.
	// That's why we check fetchedContact with *fakeContact.

	if fetchedContact != *fakeContact {
		t.Errorf("inserted contact is not the fetched contact")
	}
	err = collection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestGetContact(t *testing.T) {
	sess, err := mgo.Dial("localhost")
	if err != nil {
		t.Errorf("%s", err)
	}
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
		t.Errorf("%s", err)
	}
	id := fakeContact.Id
	fetchedContact, err := GetContact(id)
	if err != nil {
		t.Errorf("%s", err)
	}
	if *fakeContact != *fetchedContact {
		t.Errorf("Expected %+v, but got %+v\n", *fakeContact, *fetchedContact)
	}
	err = collection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}
