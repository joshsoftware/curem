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

func TestPostLeadHandler(t *testing.T) {
	_, err := NewContact(
		"Encom Inc.",
		"Hari haran",
		"hari@encom.com",
		"",
		"",
		"USA",
	)
	ts := httptest.NewServer(r)
	defer ts.Close()
	var b bytes.Buffer
	b.Write([]byte(`{"contactSlug":"hari-haran","source":"Web","owner":"Gautam","status":"Warming Up"}`))
	resp, err := http.Post(ts.URL+"/leads", "encoding/json", &b)
	if err != nil {
		t.Errorf("%s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 201 {
		x, err := GetAllLeads()
		if err != nil {
			t.Errorf("%s", err)
		}
		if len(x) == 1 {
			if x[0].ContactSlug != "hari-haran" {
				t.Errorf("expected contact slug to be hari-haran, but got %s", x[0].ContactSlug)
			}
			if x[0].Source != "Web" {
				t.Errorf("expected source to be Web, but got %s", x[0].Source)
			}
		} else {
			t.Errorf("expected 1 lead, but got %d contacts", len(x))
		}
		dropCollections(t)

	} else {
		t.Errorf("expected response code to be 201 but got %d", resp.StatusCode)
	}
}

func TestPostLeadHandlerError(t *testing.T) {
	ts := httptest.NewServer(r)
	defer ts.Close()
	var b bytes.Buffer
	b.Write([]byte(`{"contactSlug":"hari-haran","owner":"Gautam","status":"Warming Up"}`))
	resp, err := http.Post(ts.URL+"/leads", "encoding/json", &b)
	if err != nil {
		t.Errorf("%s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 500 {
		t.Errorf("expecting response code 500, but got %d", resp.StatusCode)
	}
}

func TestDeleteLeadHandler(t *testing.T) {
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
	req, err := http.NewRequest("DELETE", ts.URL+"/leads/"+l.ID.Hex(), nil)
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
	x, err := GetAllLeads()
	if err != nil {
		t.Errorf("%s", err)
	}
	if len(x) != 0 {
		t.Errorf("expected 0 leads after delete, but got %d", len(x))
	}
	dropCollections(t)
}
