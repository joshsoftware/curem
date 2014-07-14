package main

import (
	"testing"

	"github.com/joshsoftware/curem/config"

	"log"
)

func TestIsStatusValid(t *testing.T) {
	if isStatusValid("asdfgh") {
		t.Errorf("status asdfgh should be invalid")
	}
	if !isStatusValid("Won") {
		t.Errorf("status Won should be valid")
	}
}

func TestNewLead(t *testing.T) {
	f := fakeContactSlug()
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

	_, err = GetContactBySlug(fakeLead.ContactSlug)
	if err != nil {
		t.Errorf("%s", err)
	}

	dropCollections(t)
}

func fakeContactSlug() string {
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
	return fakeContact.Slug
}

func TestGetLead(t *testing.T) {
	f := fakeContactSlug()
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
	dropCollections(t)
}

func TestGetAllLeads(t *testing.T) {
	f := fakeContactSlug()
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
	f = fakeContactSlug()
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
	dropCollections(t)
}

func TestUpdateLead(t *testing.T) {
	f := fakeContactSlug()
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
	dropCollections(t)
}

func TestDeleteLead(t *testing.T) {
	f := fakeContactSlug()
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
	dropLeadsCollection(t)

	n, err = config.ContactsCollection.Count()
	if err != nil {
		t.Errorf("%s", err)
	}
	if n != 1 {
		t.Errorf("expected 1 document in the collection, but found %d", n)
	}
	dropContactsCollection(t)
}
