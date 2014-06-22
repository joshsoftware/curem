package main

import (
	"strconv"

	"log"

	"github.com/extemporalgenome/slug"
	"github.com/joshsoftware/curem/config"
	"labix.org/v2/mgo/bson"
)

// slugs is map with slugs pointing to their corresponding Ids.
// We are using an in-memory map because our current requirement is
// to cater to the order of thousands of customers.
// If the need increases, optimize the memory by storing the map separately.
var slugs map[string]bson.ObjectId = make(map[string]bson.ObjectId)

// contact type holds the fields related to a particular contact.
// omitempty tag will make sure the database doesn't contain content like:
//
//   {
//    _id: someId
//    company: ABC
//    Person: Xyz
//    Phone:
//    Skype:
//    Country:
//   }
// Instead, it will store the above data as:
//
//   {
//    _id: someId
//    company: ABC
//    Person: Xyz
//   }
type contact struct {
	Id      bson.ObjectId `bson:"_id"               json:"id"`
	Company string        `bson:"company,omitempty" json:"company,omitempty"`
	Person  string        `bson:"person,omitempty"  json:"person,omitempty"`
	Slug    string        `bson:"slug,omitempty"    json:"slug,omitempty"`
	Email   string        `bson:"email,omitempty"   json:"email,omitempty"`
	Phone   string        `bson:"phone,omitempty"   json:"phone,omitempty"`
	SkypeId string        `bson:"skypeid,omitempty" json:"skypeid,omitempty"`
	Country string        `bson:"country,omitempty" json:"country,omitempty"`
}

// NewContact takes the fields of a contact, initializes a struct of contact type and returns
// the pointer to that struct.
func NewContact(company, person, email, phone, skypeid, country string) (*contact, error) {
	doc := contact{
		Id:      bson.NewObjectId(),
		Company: company,
		Person:  person,
		Email:   email,
		Phone:   phone,
		SkypeId: skypeid,
		Country: country,
	}
	slugify(&doc)
	err := config.ContactsCollection.Insert(doc)
	if err != nil {
		return &contact{}, err
	}
	return &doc, nil
}

// GetContact takes the contact Id as an argument and returns a pointer to a contact object.
func GetContact(i bson.ObjectId) (*contact, error) {
	var c contact
	err := config.ContactsCollection.FindId(i).One(&c)
	if err != nil {
		return &contact{}, err
	}
	return &c, nil
}

// GetAllContacts fetches all the contacts from the database.
func GetAllContacts() ([]contact, error) {
	var c []contact
	err := config.ContactsCollection.Find(nil).All(&c)
	if err != nil {
		return []contact{}, err
	}
	return c, nil
}

// Update updates the contact in the database.
// First, fetch a contact from the database and change the necessary fields.
// Then call the Update method on that contact object.
func (c *contact) Update() error {
	err := config.ContactsCollection.UpdateId(c.Id, c)
	return err
}

// Delete deletes the contact from the database.
func (c *contact) Delete() error {
	return config.ContactsCollection.RemoveId(c.Id)
}

func slugify(c *contact) {
	base := slug.SlugAscii(c.Person)
	temp := base
	i := 1
	for {
		if _, ok := slugs[temp]; ok {
			temp = base + "-" + strconv.Itoa(i)
			i++
		} else {
			c.Slug = temp
			slugs[temp] = c.Id
			return
		}

	}
}

// cacheSlugs fetches all the records from an existing contacts
// collection and initialize a map with all the slugs pointing to
// their corresponding Ids.
func cacheSlugs() {
	var c []contact
	err := config.ContactsCollection.Find(nil).All(&c)
	if err != nil {
		log.Fatalf("%s", err)
	}
	for _, val := range c {
		slugs[val.Slug] = val.Id
	}
}
