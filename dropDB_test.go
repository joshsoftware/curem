package main

import (
	"testing"

	"github.com/joshsoftware/curem/config"
)

func dropCollections(t *testing.T) {
	err := config.ContactsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
	err = config.LeadsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func dropContactsCollection(t *testing.T) {
	err := config.ContactsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func dropLeadsCollection(t *testing.T) {
	err := config.LeadsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}
