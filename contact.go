package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type contact struct {
	Id      bson.ObjectId `bson:"_id"`
	Company string        `bson:"company,omitempty"`
	Person  string        `bson:"person,omitempty"`
	Email   string        `bson:"email,omitempty"`
	Phone   string        `bson:"phone,omitempty"`
	SkypeID string        `bson:"skypeid,omitempty"`
	Country string        `bson:"country,omitempty"`
}

func NewContact(c *mgo.Collection, com, p, e, ph, s, cou string) (*contact, error) {
	doc := contact{Id: bson.NewObjectId(), Company: com, Person: p, Email: e, Phone: ph, SkypeID: s, Country: cou}
	err := c.Insert(doc)
	if err != nil {
		return &contact{}, err
	}
	return &doc, nil
}
