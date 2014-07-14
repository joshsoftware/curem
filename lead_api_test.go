package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestGetLeadsHandler(t *testing.T) {
	ts := httptest.NewServer(r)
	defer ts.Close()
	c, err := NewContact(
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
	l, err := NewLead(c.Slug, "Web", "Gautam", "Won", 3, 5, 2, "3rd July, 2014", nil)
	x := []lead{*l}
	y, err := json.Marshal(x)
	resp, err := http.Get(ts.URL + "/leads")
	if err != nil {
		t.Errorf("%s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("%s", err)
	}
	bodystring := strings.TrimSpace(string(body))
	if bodystring != string(y) {
		t.Errorf("expected %s, but got %s", string(y), bodystring)
	}
	dropCollections(t)
}

func TestGetLeadHandler(t *testing.T) {
	ts := httptest.NewServer(r)
	defer ts.Close()
	c, err := NewContact(
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
	l, err := NewLead(c.Slug, "Web", "Gautam", "Won", 3, 5, 2, "3rd July, 2014", nil)
	y, err := json.Marshal(l)
	resp, err := http.Get(ts.URL + "/leads/" + l.ID.Hex())
	if err != nil {
		t.Errorf("%s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("%s", err)
	}
	bodystring := strings.TrimSpace(string(body))
	if bodystring != string(y) {
		t.Errorf("expected %s, but got %s", string(y), bodystring)
	}
	dropCollections(t)
}