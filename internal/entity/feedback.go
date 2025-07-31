package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Feedback struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	InterviewerID      string             `bson:"interviewerId" json:"interviewerId"`
	CandidateID        string             `bson:"candidateId" json:"candidateId"`
	SelectionProcessID string             `bson:"selectionProcessId" json:"selectionProcessId"`
	Stage              string             `bson:"stage" json:"stage"`
	Comments           string             `bson:"comments" json:"comments"`
	Rating             int                `bson:"rating" json:"rating"`
	CreatedAt          time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt          time.Time          `bson:"updatedAt" json:"updatedAt"`

	//Interviewer      User             `bson:"interviewer" json:"interviewer"`
	//Candidate        Candidate        `bson:"candidate" json:"candidate"`
	//SelectionProcess SelectionProcess `bson:"selectionProcess" json:"selectionProcess"`
}
