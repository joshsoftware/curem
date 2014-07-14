package main

import (
	"errors"
	"math/rand"
	"net/mail"
	"strconv"
	"time"

	"log"

	"github.com/extemporalgenome/slug"
	"github.com/joshsoftware/curem/config"
	"labix.org/v2/mgo/bson"
)

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
	ID        bson.ObjectId `bson:"_id"                  json:"id"`
	Company   string        `bson:"company,omitempty"    json:"company,omitempty"`
	Person    string        `bson:"person,omitempty"     json:"person,omitempty"`
	Slug      string        `bson:"slug,omitempty"       json:"slug,omitempty"`
	Email     string        `bson:"email,omitempty"      json:"email,omitempty"`
	Phone     string        `bson:"phone,omitempty"      json:"phone,omitempty"`
	SkypeID   string        `bson:"skypeID,omitempty"    json:"skypeID,omitempty"`
	Country   string        `bson:"country,omitempty"    json:"country,omitempty"`
	CreatedAt time.Time     `bson:"createdAt,omitempty"  json:"createdAt,omitempty"`
	UpdatedAt time.Time     `bson:"updatedAt,omitempty"  json:"updatedAt,omitempty"`
}

// incomingContact type is used for handling PATCH requests.
// To understand why we are using pointer types for fields,
// refer to http://blog.golang.org/json-and-go and
// https://willnorris.com/2014/05/go-rest-apis-and-pointers .
// Using pointer types, we can differentiate intentional nil value fields
// and empty fields.
// This type is used *only* for decoding json obtained from a PATCH request.
// If you don't want a field to be updated, remove the field from this type.
type incomingContact struct {
	ID      *bson.ObjectId `json:"id"`
	Company *string        `json:"company,omitempty"`
	Person  *string        `json:"person,omitempty"`
	Slug    *string        `json:"slug,omitempty"`
	Email   *string        `json:"email,omitempty"`
	Phone   *string        `json:"phone,omitempty"`
	SkypeID *string        `json:"skypeID,omitempty"`
	Country *string        `json:"country,omitempty"`
}

// copyIncomingFields copies the fields from an incomingContact into a
// contact object.
// Using multiple if statements provides more granularity, while allowing
// update of only some specific fields.
// TODO(Hari): Use reflect package instead of multiple if statements
func (c *contact) copyIncomingFields(i *incomingContact) error {
	if i.ID != nil {
		if *i.ID != c.ID {
			return errors.New("id doesn't match")
		}
	}
	if i.Slug != nil {
		if *i.Slug != c.Slug {
			return errors.New("slug can't be updated")
		}
	}
	if i.Company != nil {
		c.Company = *i.Company
	}
	if i.Person != nil {
		c.Person = *i.Person
	}
	if i.Slug != nil {
		c.Slug = *i.Slug
	}
	if i.Email != nil {
		c.Email = *i.Email
	}
	if i.Phone != nil {
		c.Phone = *i.Phone
	}
	if i.SkypeID != nil {
		c.SkypeID = *i.SkypeID
	}
	if i.Country != nil {
		c.Country = *i.Country
	}
	return nil
}

func (c *contact) Validate() error {
	if c.Person == "" {
		err := errors.New("person can't be empty")
		return err
	}
	if c.Email == "" {
		err := errors.New("email can't be empty")
		return err
	}
	if _, err := mail.ParseAddress(c.Email); err != nil {
		return err
	}

	return nil
}

// NewContact takes the fields of a contact, initializes a struct of contact type and returns
// the pointer to that struct.
func NewContact(company, person, email, phone, skypeid, country string) (*contact, error) {
	doc := contact{
		ID:      bson.NewObjectId(),
		Company: company,
		Person:  person,
		Email:   email,
		Phone:   phone,
		SkypeID: skypeid,
		Country: country,
	}

	if err := (&doc).Validate(); err != nil {
		return &contact{}, err
	}
	slugifyContact(&doc)

	doc.CreatedAt = doc.ID.Time()
	doc.UpdatedAt = doc.CreatedAt

	err := config.ContactsCollection.Insert(doc)
	if err != nil {
		return &contact{}, err
	}
	return &doc, nil
}

// GetContactByID takes the contact ID as an argument and returns a pointer to the contact object.
func GetContactByID(i bson.ObjectId) (*contact, error) {
	var c contact
	err := config.ContactsCollection.FindId(i).One(&c)
	if err != nil {
		return &contact{}, err
	}
	return &c, nil
}

// GetContactBySlug takes the contact slug as an argument and returns a pointer to the contact object.
func GetContactBySlug(slug string) (*contact, error) {
	var c []contact
	err := config.ContactsCollection.Find(bson.M{"slug": slug}).All(&c)
	if err != nil {
		return &contact{}, err
	}

	if len(c) == 0 {
		return &contact{}, errors.New("no contact")
	}

	if len(c) > 1 {
		return &contact{}, errors.New("more than 1 contact found")
	}
	return &c[0], nil
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
	if err := c.Validate(); err != nil {
		return err
	}
	c.UpdatedAt = bson.Now()
	err := config.ContactsCollection.UpdateId(c.ID, c)
	return err
}

// Delete deletes the contact from the database.
func (c *contact) Delete() error {
	return config.ContactsCollection.RemoveId(c.ID)
}

// slugifyContact takes a pointer to a contact as the argument and
// creates a slug and populates the slug field of the contact.
func slugifyContact(c *contact) {
	base := slug.SlugAscii(c.Person)
	temp := base
	rand.Seed(time.Now().UnixNano()) // takes the current time in nanoseconds as the seed
	i := rand.Intn(10000)
	for {
		if contactSlugExists(temp) {
			temp = base + "-" + strconv.Itoa(i)
			i = rand.Intn(10000)
		} else {
			c.Slug = temp
			return
		}

	}
}

// contactSlugExists returns true if a slug for a contact already exists in the database
// else returns false.
func contactSlugExists(slug string) bool {
	var c []contact
	err := config.ContactsCollection.Find(bson.M{"slug": slug}).All(&c)
	if err != nil {
		log.Println(err)
		return true // returns true just to be safe.
	}
	if len(c) == 0 {
		return false
	}
	return true
}

// Leads returns all the leads of a particular contact.
func (c *contact) Leads() ([]lead, error) {
	var l []lead
	err := config.LeadsCollection.Find(bson.M{"contactSlug": c.Slug}).All(&l)
	if err != nil {
		return l, err
	}
	return l, nil
}
