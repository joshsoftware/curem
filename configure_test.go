package main

import "github.com/joshsoftware/curem/config"

// This ensures that we use a separate test database when `go test` is run.
func init() {
	c := make(map[string]string)
	c["name"] = "test"
	c["url"] = "localhost"
	c["leads"] = "newlead"
	c["contacts"] = "newcontact"

	config.Configure(c)
}
