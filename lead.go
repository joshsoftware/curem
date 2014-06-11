package main

import (
	"github.com/joshsoftware/curem/config"
	"labix.org/v2/mgo/bson"
)

type lead struct {
	Id                 bson.ObjectId `bson:"_id"                          json:"id"`
	ContactId          bson.ObjectId `bson:"contact,omitempty"            json:"contact,omitempty"`
	Source             string        `bson:"source,omitempty"             json:"source,omitempty"`
	Owner              string        `bson:"owner,omitempty"              json:"owner,omitempty"`
	Status             string        `bson:"status,omitempty"             json:"status,omitempty"`
	TeamSize           float64       `bson:"teamsize,omitempty"           json:"teamsize,omitempty"`
	RatePerHour        float64       `bson:"rateperhour,omitempty"        json:"rateperhour,omitempty"`
	DurationInMonths   float64       `bson:"durationinmonths,omitempty"   json:"durationinmonths,omitempty"`
	EstimatedStartDate string        `bson:"estimatedstartdate,omitempty" json:"estimatedstartdate,omitempty"`
	Comments           []string      `bson:"comments,omitempty"           json:"comments,omitempty"`
}

// NewLead takes the fields of a lead, initializes a struct of lead type and returns
// the pointer to that struct.
// Also, It inserts the lead object into the database.
func NewLead(cid bson.ObjectId, source, owner, status string, teamsize, rate, duration float64,
	start string, comments []string) (*lead, error) {
	doc := lead{
		Id:                 bson.NewObjectId(),
		ContactId:          cid,
		Source:             source,
		Owner:              owner,
		Status:             status,
		TeamSize:           teamsize,
		RatePerHour:        rate,
		DurationInMonths:   duration,
		EstimatedStartDate: start,
		Comments:           comments,
	}
	err := config.LeadsCollection.Insert(doc)
	if err != nil {
		return &lead{}, err
	}
	return &doc, nil
}

func GetLead(i bson.ObjectId) (*lead, error) {
	var l lead
	err := config.LeadsCollection.FindId(i).One(&l)
	if err != nil {
		return &lead{}, err
	}
	return &l, nil
}

func (l *lead) Update() error {
	_, err := config.LeadsCollection.UpsertId(l.Id, l)
	return err
}

func (l *lead) Delete() error {
	return config.LeadsCollection.RemoveId(l.Id)
}
