package main

import (
	"testing"

	"github.com/joshsoftware/curem/config"
	"labix.org/v2/mgo/bson"
)

// This ensures that we use a separate test database when `go test` is run.
func init() {
	c := make(map[string]string)
	c["name"] = "test"
	c["url"] = "localhost"
	c["leads"] = "newlead"
	c["contacts"] = "newcontact"

	config.Configure(c)
}

func TestNewContact(t *testing.T) {
	fakeContact, err := NewContact(
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
	err = config.ContactsCollection.Find(bson.M{}).One(&fetchedContact)
	if err != nil {
		t.Errorf("%s", err)
	}

	// fakeContact is a pointer, because NewContact returns a pointer to a struct of contact type.
	// That's why we check fetchedContact with *fakeContact.

	if fetchedContact != *fakeContact {
		t.Errorf("inserted contact is not the fetched contact")
	}
	err = config.ContactsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestValidateNewContact(t *testing.T) {
	_, err := NewContact(
		"Encom Inc.",
		"",
		"flynn@encom.com",
		"",
		"",
		"USA",
	)
	if err == nil {
		t.Errorf("%s", "error shouldn't be nil when person is empty")
	}
	_, err = NewContact(
		"Encom Inc.",
		"Sam Flynn",
		"",
		"",
		"",
		"USA",
	)
	if err == nil {
		t.Errorf("%s", "error shouldn't be nil when email is empty")
	}
	_, err = NewContact(
		"Encom Inc.",
		"Sam Flynn",
		"x@.xyzc.com",
		"",
		"",
		"USA",
	)
	if err == nil {
		t.Errorf("%s", "error shouldn't be nil when email is invalid")
	}
}

func TestGetContact(t *testing.T) {
	fakeContact, err := NewContact(
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
	err = config.ContactsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestGetAllContacts(t *testing.T) {
	_, err := NewContact(
		"Encom Inc.",
		"Sam Flynn",
		"samflynn@encom.com",
		"103-345-456",
		"sam_flynn",
		"USA",
	)
	if err != nil {
		t.Errorf("%s", err)
	}
	_, err = NewContact(
		"Encom Inc.",
		"Kevin Flynn",
		"kevinflynn@encom.com",
		"234-877-988",
		"kevin_flynn",
		"USA",
	)
	if err != nil {
		t.Errorf("%s", err)
	}

	fetchedContacts, err := GetAllContacts()
	if err != nil {
		t.Errorf("%s", err)
	}
	if len(fetchedContacts) != 2 {
		t.Errorf("expected 2 contacts, but got %d", len(fetchedContacts))
	}
	err = config.ContactsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestUpdateContact(t *testing.T) {
	fakeContact, err := NewContact(
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
	fakeContact.Country = "India"
	fakeContact.Update()
	fetchedContact, err := GetContact(fakeContact.Id)
	if err != nil {
		t.Errorf("%s", err)
	}
	if fetchedContact.Country != "India" {
		t.Errorf("%s", "contact not updated")
	}
	err = config.ContactsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestDelete(t *testing.T) {
	fakeContact, err := NewContact(
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
	err = fakeContact.Delete()
	if err != nil {
		t.Errorf("%s", err)
	}
	n, err := config.ContactsCollection.Count()
	if err != nil {
		t.Errorf("%s", err)
	}
	if n != 0 {
		t.Errorf("expected 0 documents in the collection, but found %d", n)
	}
	err = config.ContactsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}
