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
// Also, It inserts the lead data into a mongoDB collection, which is passed as the first parameter.
func NewLead(cid bson.ObjectId, source, owner, status string, teamsize, rate, duration float64,
	start string, comments []string) (*lead, error) {
	collection := config.Db.C("newlead")
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
	err := collection.Insert(doc)
	if err != nil {
		return &lead{}, err
	}
	return &doc, nil
}

func GetLead(i bson.ObjectId) (*lead, error) {
	collection := config.Db.C("newlead")
	var l lead
	err := collection.FindId(i).One(&l)
	if err != nil {
		return &lead{}, err
	}
	return &l, nil
}

func (l *lead) Delete() error {
	i := l.Id
	collection := config.Db.C("newlead")
	err := collection.RemoveId(i)
	return err
}
