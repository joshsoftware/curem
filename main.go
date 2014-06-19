package main

import (
	"net/http"

	"github.com/joshsoftware/curem/config"
)

func main() {

	c := make(map[string]string)
	c["name"] = "test"
	c["url"] = "localhost"
	c["leads"] = "newlead"
	c["contacts"] = "newcontact"

	config.Configure(c)

	// r is a *mux.Router defined in contact_api.go
	http.ListenAndServe(":3000", r)
}
