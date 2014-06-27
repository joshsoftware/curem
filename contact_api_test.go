package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/joshsoftware/curem/config"
)

func TestGetContactsHandler(t *testing.T) {
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
	x := []contact{*c}
	y, err := json.Marshal(x)
	resp, err := http.Get(ts.URL + "/contacts")
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
	err = config.ContactsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestPostContactsHandler(t *testing.T) {
	ts := httptest.NewServer(r)
	defer ts.Close()
	var b bytes.Buffer
	b.Write([]byte(`{"person":"Hari haran","email":"hari@example.com"}`))
	resp, err := http.Post(ts.URL+"/contacts", "encoding/json", &b)
	if err != nil {
		t.Errorf("%s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 201 {
		x, err := GetAllContacts()
		if err != nil {
			t.Errorf("%s", err)
		}
		if len(x) == 1 {
			if x[0].Person != "Hari haran" {
				t.Errorf("expected person to be Hari haran, but got %s", x[0].Person)
			}
			if x[0].Email != "hari@example.com" {
				t.Errorf("expected email to be hari@example.com, but got %s", x[0].Email)
			}
		} else {
			t.Errorf("expected 1 contact, but got %d contacts", len(x))
		}
		err = config.ContactsCollection.DropCollection()
		if err != nil {
			t.Errorf("%s", err)
		}

	} else {
		t.Errorf("expected response code to be 201 but got %d", resp.StatusCode)
	}
}

func TestGetContactHandler(t *testing.T) {
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
	y, err := json.Marshal(c)
	resp, err := http.Get(ts.URL + "/contacts/" + c.Slug)
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
	err = config.ContactsCollection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}
