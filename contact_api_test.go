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
