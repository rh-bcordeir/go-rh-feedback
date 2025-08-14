package database

import (
	"context"
	"errors"
	"log"

	"github.com/brunocordeiro180/go-rh-feedback/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CandidateDB struct {
	collection *mongo.Collection
}

func NewCandidateDB(client *mongo.Client) *CandidateDB {
	return &CandidateDB{
		collection: client.Database("feedback_db").Collection("candidates"),
	}
}

func (c *CandidateDB) CreateCandidate(candidate *entity.Candidate) error {
	candidate.ID = primitive.NewObjectID()
	_, err := c.collection.InsertOne(context.Background(), candidate)
	if err != nil {
		log.Println(err)
		return errors.New("unable to insert candidate")
	}
	return nil
}

func (c *CandidateDB) ListCandidates() ([]entity.Candidate, error) {
	var candidates []entity.Candidate
	cursor, err := c.collection.Find(context.Background(), primitive.D{})
	if err != nil {
		log.Println(err)
		return nil, errors.New("unable to list candidates")
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var candidate entity.Candidate
		if err := cursor.Decode(&candidate); err != nil {
			log.Println(err)
			return nil, errors.New("error decoding candidate")
		}
		candidates = append(candidates, candidate)
	}

	if err := cursor.Err(); err != nil {
		log.Println(err)
		return nil, errors.New("cursor error")
	}

	return candidates, nil
}
