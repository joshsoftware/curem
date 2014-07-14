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
		"Encom Ic.",
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

func TestGetContactByID(t *testing.T) {
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
	id := fakeContact.ID
	fetchedContact, err := GetContactByID(id)
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

func TestGetNonExistingContactByID(t *testing.T) {
	_, err := GetContactByID(bson.ObjectIdHex("53b112bde3bdea2642000002"))
	if err == nil {
		t.Errorf("%s", "error shouldn't be nil when we try a fetch a non existent contact")
	}
}

func TestGetContactBySlug(t *testing.T) {
	c, err := NewContact(
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
	f, err := GetContactBySlug(c.Slug)
	if err != nil {
		t.Errorf("%s", err)
	}
	if *c != *f {
		t.Errorf("expected %+v, but got %+v", *c, *f)
	}
	err = config.ContactsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestGetNonExistingContactBySlug(t *testing.T) {
	_, err := GetContactBySlug("nlvnjrelvenliqas")
	if err == nil {
		t.Errorf("%s", "error shouldn't be nil when we try a fetch a non existent contact")
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
	fetchedContact, err := GetContactByID(fakeContact.ID)
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

func TestSlugifyContact(t *testing.T) {
	c, err := NewContact(
		"Encom Inc.",
		"Sam Flynn",
		"samflynn@encom.com",
		"103-345-456",
		"sam_flynn",
		"USA",
	)
	if c.Slug != "sam-flynn" {
		t.Errorf("expected slug to be %s, but got %s", "sam-flynn", c.Slug)
	}
	d := &contact{
		Person: "Sam Flynn",
		Email:  "sam@example.com",
	}
	slugifyContact(d)
	if d.Slug == "sam-flynn" {
		t.Errorf("expected something other than %s as slug", "sam-flynn")
	}
	err = config.ContactsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestContactSlugExists(t *testing.T) {
	c, err := NewContact(
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
	if !contactSlugExists(c.Slug) {
		t.Errorf("%s", "expected contactSlugExists to return true but returns false")
	}
	err = config.ContactsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}
