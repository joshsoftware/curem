package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v1"

	"labix.org/v2/mgo"
)

var (
	Config  *AppConfig
	Session *mgo.Session
	Db      *mgo.Database
)

type AppConfig struct {
	Mongo struct {
		Name string
		Url  string
	}
}

func loadConfig() {
	file, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("error opening config.yaml: %s", err)
	}
	Config = new(AppConfig)
	if err := yaml.Unmarshal(file, &Config); err != nil {
		log.Fatalf("error parsing config.yaml: %s", err)
	}
}

func init() {
	loadConfig()
	session, err := mgo.Dial(Config.Mongo.Url)
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

	Db = session.DB(Config.Mongo.Name)
	Session = session
}
