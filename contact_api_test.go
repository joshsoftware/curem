package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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
	dropContactsCollection(t)
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
		dropContactsCollection(t)

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
	dropContactsCollection(t)
}

func TestPatchContactHandler(t *testing.T) {
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
	var b bytes.Buffer
	b.Write([]byte(`{"person":"Hari haran","email":"hari@example.com","country":""}`))
	req, err := http.NewRequest("PATCH", ts.URL+"/contacts/"+c.Slug, &b)
	if err != nil {
		t.Errorf("%s", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("%s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
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
			if x[0].Country != "" {
				t.Errorf(`expected country to be "", but got %s`, x[0].Country)
			}
		} else {
			t.Errorf("expected 1 contact, but got %d contacts", len(x))
		}
		dropContactsCollection(t)
	} else {
		t.Errorf("expected response code to be 200, but got %d", resp.StatusCode)
	}
}

func TestDeleteContactHandler(t *testing.T) {
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
	req, err := http.NewRequest("DELETE", ts.URL+"/contacts/"+c.Slug, nil)
	if err != nil {
		t.Errorf("%s", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("%s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 204 {
		t.Errorf("expected response status code to be 204, but got %d", resp.StatusCode)
	}
	x, err := GetAllContacts()
	if err != nil {
		t.Errorf("%s", err)
	}
	if len(x) != 0 {
		t.Errorf("expected 0 contacts after delete, but got %d", len(x))
	}
	dropContactsCollection(t)
}
