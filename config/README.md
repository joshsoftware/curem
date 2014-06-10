##config

Package config initializes variables of Db, Session, etc.

They should be specified by calling the `Configure()` function with a map[string]string as an argument.

For example,

~~~go
package main

import (
	"fmt"

	"github.com/joshsoftware/curem/config"
)

func main() {
	c := make(map[string]string)
	c["name"] = "hello"
	c["url"] = "localhost"
	c["leads"] = "newlead"
	c["contacts"] = "newcontact"

	config.Configure(c)

	fmt.Printf("%+v\n", config.Db)
	fmt.Printf("%+v\n", config.Session)
	fmt.Printf("%+v\n", config.LeadsCollection)
	fmt.Printf("%+v\n", config.ContactsCollection)
}
~~~