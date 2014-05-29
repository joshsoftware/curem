package main

import (
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
	Id      bson.ObjectId `bson:"_id"`
	Company string        `bson:"company,omitempty"`
	Person  string        `bson:"person,omitempty"`
	Email   string        `bson:"email,omitempty"`
	Phone   string        `bson:"phone,omitempty"`
	SkypeId string        `bson:"skypeid,omitempty"`
	Country string        `bson:"country,omitempty"`
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
