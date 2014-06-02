package main

import (
	"github.com/joshsoftware/curem/config"
	"labix.org/v2/mgo"
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
//Instead, it will store the above data as:
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
	Email   string        `bson:"email,omitempty"   json:"email,omitempty"`
	Phone   string        `bson:"phone,omitempty"   json:"phone,omitempty"`
	SkypeId string        `bson:"skypeid,omitempty" json:"skypeid,omitempty"`
	Country string        `bson:"country,omitempty" json:"country,omitempty"`
}

// NewContact takes the fields of a contact, initializes a struct of contact type and returns
// the pointer to that struct.
// Also, It inserts the contact data into a mongoDB collection, which is passed as the first parameter.
func NewContact(c *mgo.Collection, company, person, email, phone, skypeid, country string) (*contact, error) {
	doc := contact{
		Id:      bson.NewObjectId(),
		Company: company,
		Person:  person,
		Email:   email,
		Phone:   phone,
		SkypeId: skypeid,
		Country: country,
	}
	err := c.Insert(doc)
	if err != nil {
		return &contact{}, err
	}
	return &doc, nil
}

// TODO(Hari): Move session logic into a config file and a separate function
func GetContact(i bson.ObjectId) (*contact, error) {
	collection := config.Db.C("newcontact")
	var c contact
	err := collection.FindId(i).One(&c)
	if err != nil {
		return &contact{}, err
	}
	return &c, nil
}

func DeleteContact(i bson.ObjectId) error {
	collection := config.Db.C("newcontact")
	err := collection.RemoveId(i)
	return err
}
