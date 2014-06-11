package main

import (
	"fmt"
	"testing"

	"github.com/joshsoftware/curem/config"

	"log"

	"labix.org/v2/mgo/bson"
)

func TestNewLead(t *testing.T) {
	f := fakeContactId()
	fakeLead, err := NewLead(
		f,
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
	err = config.ContactsCollection.FindId(fakeLead.ContactId).One(&refContact)
	if err != nil {
		t.Errorf("%s", err)
	}
	fmt.Printf("%+v\n", refContact)

	err = config.LeadsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}

	// Drop collection created by fakeContactId()
	err = config.ContactsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func fakeContactId() bson.ObjectId {
	fakeContact, err := NewContact(
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
	f := fakeContactId()
	fakeLead, err := NewLead(
		f,
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
	err = config.LeadsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
	// Drop collection created by fakeContactId()
	err = config.ContactsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestUpdateLead(t *testing.T) {
	f := fakeContactId()
	fakeLead, err := NewLead(
		f,
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
	fakeLead.Status = "Won"
	err = fakeLead.Update()
	if err != nil {
		t.Errorf("%s", err)
	}
	fetchedLead, err := GetLead(fakeLead.Id)
	if err != nil {
		t.Errorf("%s", err)
	}
	if fetchedLead.Status != "Won" {
		t.Errorf("%s", "lead not updated")
	}
	err = config.LeadsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
	err = config.ContactsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestDeleteLead(t *testing.T) {
	f := fakeContactId()
	fakeLead, err := NewLead(
		f,
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
	err = fakeLead.Delete()
	if err != nil {
		t.Errorf("%s", err)
	}
	n, err := config.LeadsCollection.Count()
	if err != nil {
		t.Errorf("%s", err)
	}
	if n != 0 {
		t.Errorf("expected 0 documents in the collection, but found %d", n)
	}
	err = config.LeadsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}

	n, err = config.ContactsCollection.Count()
	if err != nil {
		t.Errorf("%s", err)
	}
	if n != 1 {
		t.Errorf("expected 1 document in the collection, but found %d", n)
	}
	err = config.ContactsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}
