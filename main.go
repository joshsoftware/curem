package main

import (
	"github.com/codegangsta/negroni"
	"github.com/hariharan-uno/cors"
	"github.com/joshsoftware/curem/config"
)

func main() {

	c := make(map[string]string)
	c["name"] = "dev"
	c["url"] = "localhost"
	c["leads"] = "newlead"
	c["contacts"] = "newcontact"

	config.Configure(c)

	opts := cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PATCH", "DELETE"},
	}

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(opts.Allow))

	n.UseHandler(r) // r is a *mux.Router defined in contact_api.go
	n.Run(":3000")
}
