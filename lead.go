package main

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/joshsoftware/curem/config"
	"labix.org/v2/mgo/bson"
)

var statusOptions = []string{"Won", "Lost", "Warming Up", "Cooling Down", "Tech", "Negotiating", "Proposal", "Contract"}

type lead struct {
	ID                 bson.ObjectId `bson:"_id"                          json:"id"`
	ContactID          bson.ObjectId `bson:"contactID,omitempty"          json:"contactID,omitempty"`
	Source             string        `bson:"source,omitempty"             json:"source,omitempty"`
	Owner              string        `bson:"owner,omitempty"              json:"owner,omitempty"`
	Status             string        `bson:"status,omitempty"             json:"status,omitempty"`
	TeamSize           float64       `bson:"teamSize,omitempty"           json:"teamSize,omitempty"`
	RatePerHour        float64       `bson:"ratePerHour,omitempty"        json:"ratePerHour,omitempty"`
	DurationInMonths   float64       `bson:"durationInMonths,omitempty"   json:"durationInMonths,omitempty"`
	EstimatedStartDate string        `bson:"estimatedStartDate,omitempty" json:"estimatedStartDate,omitempty"`
	Comments           []string      `bson:"comments,omitempty"           json:"comments,omitempty"`
	CreatedAt          time.Time     `bson:"createdAt,omitempty"          json:"createdAt,omitempty"`
	UpdatedAt          time.Time     `bson:"updatedAt,omitempty"          json:"updatedAt,omitempty"`
}

type incomingLead struct {
	ID                 *bson.ObjectId `json:"id"`
	ContactID          *bson.ObjectId `json:"contactID,omitempty"`
	Source             *string        `json:"source,omitempty"`
	Owner              *string        `json:"owner,omitempty"`
	Status             *string        `json:"status,omitempty"`
	TeamSize           *float64       `json:"teamSize,omitempty"`
	RatePerHour        *float64       `json:"ratePerHour,omitempty"`
	DurationInMonths   *float64       `json:"durationInMonths,omitempty"`
	EstimatedStartDate *string        `json:"estimatedStartDate,omitempty"`
	Comments           *[]string      `json:"comments,omitempty"`
}

func (l *lead) copyIncomingFields(i *incomingLead) error {
	if i.ID != nil {
		if *i.ID != l.ID {
			return errors.New("id doesn't match")
		}
	}
	if i.ContactID != nil {
		l.ContactID = *i.ContactID
	}
	if i.Source != nil {
		l.Source = *i.Source
	}
	if i.Owner != nil {
		l.Owner = *i.Owner
	}
	if i.Status != nil {
		l.Status = *i.Status
	}
	if i.TeamSize != nil {
		l.TeamSize = *i.TeamSize
	}
	if i.RatePerHour != nil {
		l.RatePerHour = *i.RatePerHour
	}
	if i.DurationInMonths != nil {
		l.DurationInMonths = *i.DurationInMonths
	}
	if i.EstimatedStartDate != nil {
		l.EstimatedStartDate = *i.EstimatedStartDate
	}
	if i.Comments != nil {
		l.Comments = *i.Comments
	}
	return nil
}

func (l *lead) Validate() error {
	if l.ContactID == "" {
		return errors.New("contact ID can't be empty")
	}
	if l.Source == "" {
		return errors.New("source can't be empty")
	}
	if l.Owner == "" {
		return errors.New("owner can't be empty")
	}
	if l.Status == "" {
		return errors.New("status can't be empty")
	}
	if !isStatusValid(l.Status) {
		return fmt.Errorf("status is invalid. Can only be one of %s", statusOptions)
	}
	return nil
}

func isStatusValid(s string) bool {
	for _, v := range statusOptions {
		if v == s {
			return true
		}
	}
	return false
}

// NewLead takes the fields of a lead, initializes a struct of lead type and returns
// the pointer to that struct.
// Also, It inserts the lead object into the database.
func NewLead(cid bson.ObjectId, source, owner, status string, teamsize, rate, duration float64,
	start string, comments []string) (*lead, error) {
	doc := lead{
		ID:                 bson.NewObjectId(),
		ContactID:          cid,
		Source:             strings.Title(source),
		Owner:              strings.Title(owner),
		Status:             strings.Title(status),
		TeamSize:           teamsize,
		RatePerHour:        rate,
		DurationInMonths:   duration,
		EstimatedStartDate: start,
		Comments:           comments,
	}

	if err := (&doc).Validate(); err != nil {
		return &lead{}, err
	}

	doc.CreatedAt = doc.ID.Time()
	doc.UpdatedAt = doc.CreatedAt

	err := config.LeadsCollection.Insert(doc)
	if err != nil {
		return &lead{}, err
	}
	return &doc, nil
}

// GetLead takes the lead ID as an argument and returns a pointer to a lead object.
func GetLead(i bson.ObjectId) (*lead, error) {
	var l lead
	err := config.LeadsCollection.FindId(i).One(&l)
	if err != nil {
		return &lead{}, err
	}
	return &l, nil
}

// GetAllLeads fetches all the leads from the database.
func GetAllLeads() ([]lead, error) {
	var l []lead
	err := config.LeadsCollection.Find(nil).All(&l)
	if err != nil {
		return []lead{}, err
	}
	return l, nil
}

// Update updates the lead in the database.
// First, fetch a lead from the database and change the necessary fields.
// Then call the Update method on that lead object.
func (l *lead) Update() error {
	if err := l.Validate(); err != nil {
		return err
	}
	l.UpdatedAt = bson.Now()
	err := config.LeadsCollection.UpdateId(l.ID, l)
	return err
}

// Delete deletes the lead from the database.
func (l *lead) Delete() error {
	return config.LeadsCollection.RemoveId(l.ID)
}
