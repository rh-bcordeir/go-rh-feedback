package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Feedback struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	InterviewerID      primitive.ObjectID `bson:"interviewerId" `
	CandidateID        primitive.ObjectID `bson:"candidateId" `
	SelectionProcessID primitive.ObjectID `bson:"selectionProcessId" `
	Stage              string             `bson:"stage" ` //etapa do processo (Triagem, Técnica, Comportamental, etc)
	Comments           string             `bson:"comments" `
	Rating             int                `bson:"rating" `
	CreatedAt          time.Time          `bson:"createdAt" `
	UpdatedAt          time.Time          `bson:"updatedAt"`
}
