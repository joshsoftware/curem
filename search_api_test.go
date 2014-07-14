package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joshsoftware/curem/config"
)

func TestSearchHandler(t *testing.T) {
	ts := httptest.NewServer(r)
	defer ts.Close()
	_, err := NewContact(
		"Encom Inc.",
		"Sam Flynn",
		"sam@encom.com",
		"",
		"",
		"USA",
	)
	if err != nil {
		t.Errorf("%s", err)
	}
	_, err = NewContact(
		"Encom Inc.",
		"Kevin Flynn",
		"kevin@encom.com",
		"",
		"",
		"USA",
	)
	if err != nil {
		t.Errorf("%s", err)
	}

	// Configuring again because after every test, we are
	// dropping collections, in turn losing text indexes.
	// Configuring again creates collection along with text index.

	c := make(map[string]string)
	c["name"] = "test"
	c["url"] = "localhost"
	c["leads"] = "newlead"
	c["contacts"] = "newcontact"

	config.Configure(c)

	resp, err := http.Get(ts.URL + "/search?q=flynn")
	if err != nil {
		t.Errorf("%s", err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("%s", err)
	}
	var y []interface{}
	err = json.Unmarshal(b, &y)
	if err != nil {
		t.Errorf("%s", err)
	}
	if len(y) != 2 {
		t.Errorf("expected 2 search results, but got %d", len(y))
	}
	dropContactsCollection(t)
}
