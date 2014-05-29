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
	err = collection.DropCollection() //Fresh test DB collection
	if err != nil {
		t.Errorf("%s", err)
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
		t.Errorf("%s", err)
	}
	var updatedcontact contact
	err = sess.DB("test").C("newcontact").Find(bson.M{}).One(&updatedcontact)
	if err != nil {
		t.Errorf("%s", err)
	}
	if updatedcontact != *fakeContact {
		t.Errorf("Inserted contact is not the fetched contact")
	}
}
