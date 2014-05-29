package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type lead struct {
	Id                 bson.NewObjectId `bson:"_id"`
	Contact            *mgo.DBRef       `bson:"contact,omitempty"`
	Source             string           `bson:"source,omitempty"`
	Owner              string           `bson:"owner,omitempty"`
	Status             string           `bson:"status,omitempty"`
	TeamSize           float            `bson:"teamsize,omitempty"`
	RatePerHour        float            `bson:"rateperhour,omitempty"`
	DurationInMonths   float            `bson:"durationinmonths,omitempty"`
	EstimatedStartDate string           `bson:"estimatedstartdate,omitempty"`
	//Here we choose not to use time.Time because omitempty isn't supported for time.Time
	Comments []string `bson:"comments,omitempty"`
}

// NewLead takes the fields of a lead, initializes a struct of lead type and returns
// the pointer to that struct.
// Also, It inserts the lead data into a mongoDB collection, which is passed as the first parameter.
func NewLead(c *mgo.Collection, r *mgo.DBRef, source, owner, status,
	teamsize, rate, duration, start string, comments []string) (*lead, error) {

	doc := lead{
		Id:                 bson.NewObjectId(),
		Contact:            r,
		Source:             source,
		Owner:              owner,
		Status:             status,
		TeamSize:           teamsize,
		RatePerHour:        rate,
		DurationInMonths:   duration,
		EstimatedStartDate: start,
		Comments:           comments,
	}
	err := c.Insert(doc)
	if err != nil {
		return &lead{}, err
	}
	return &doc, nil
}
