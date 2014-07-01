package main

import (
	"testing"

	"github.com/joshsoftware/curem/config"

	"log"

	"labix.org/v2/mgo/bson"
)

func TestNewLead(t *testing.T) {
	f := fakeContactID()
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

	var refContact contact
	err = config.ContactsCollection.FindId(fakeLead.ContactID).One(&refContact)
	if err != nil {
		t.Errorf("%s", err)
	}

	err = config.LeadsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}

	// Drop collection created by fakeContactID()
	err = config.ContactsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func fakeContactID() bson.ObjectId {
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
	return fakeContact.ID
}

func TestGetLead(t *testing.T) {
	f := fakeContactID()
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
	id := fakeLead.ID
	fetchedLead, err := GetLead(id)
	if err != nil {
		t.Errorf("%s", err)
	}
	if fetchedLead.ID != fakeLead.ID {
		t.Errorf("Expected id of %v, but got %v", fakeLead.ID, fetchedLead.ID)
	}
	err = config.LeadsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
	// Drop collection created by fakeContactID()
	err = config.ContactsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestGetAllLeads(t *testing.T) {
	f := fakeContactID()
	_, err := NewLead(
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
	f = fakeContactID()
	_, err = NewLead(
		f,
		"Referral",
		"Hari",
		"Won",
		4,
		20,
		3,
		"5th July, 2014",
		[]string{"Discuss technical constraints"},
	)
	if err != nil {
		t.Errorf("%s", err)
	}
	fetchedLeads, err := GetAllLeads()
	if err != nil {
		t.Errorf("%s", err)
	}
	if len(fetchedLeads) != 2 {
		t.Errorf("expected 2 leads, but got %d", len(fetchedLeads))
	}
	err = config.ContactsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
	err = config.LeadsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestUpdateLead(t *testing.T) {
	f := fakeContactID()
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
	fetchedLead, err := GetLead(fakeLead.ID)
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
	f := fakeContactID()
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
