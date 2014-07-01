package config

import (
	"log"

	"labix.org/v2/mgo"
)

var (
	Session                *mgo.Session
	Db                     *mgo.Database
	LeadsCollection        *mgo.Collection
	ContactsCollection     *mgo.Collection
	ContactsCollectionName string
)

func Configure(options map[string]string) {
	if options["name"] == "" {
		log.Fatalf("Configure requires the name of the MongoDB server")
	}
	if options["url"] == "" {
		log.Fatalf("Configure requires the url of the MongoDB server")
	}
	if options["leads"] == "" {
		log.Fatalf("Configure requires the name of the leads collection")
	}
	if options["contacts"] == "" {
		log.Fatalf("Configure requires the name of the contacts collection")
	}

	session, err := mgo.Dial(options["url"])
	if err != nil {
		panic(err)
	}

	// In the Monotonic consistency mode reads may not be entirely up-to-date,
	// but they will always see the history of changes moving forward,
	// the data read will be consistent across sequential queries in the same session,
	// and modifications made within the session will be observed in following queries.
	session.SetMode(mgo.Monotonic, true)

	// Make the session check for errors, without imposing further constraints.
	session.SetSafe(&mgo.Safe{})

	Session = session
	Db = session.DB(options["name"])
	LeadsCollection = Db.C(options["leads"])
	ContactsCollection = Db.C(options["contacts"])
	ContactsCollectionName = options["contacts"]
	index := mgo.Index{
		Key: []string{"$text:person"},
	}
	err = ContactsCollection.EnsureIndex(index)
	if err != nil {
		log.Fatalln(err)
	}
}
