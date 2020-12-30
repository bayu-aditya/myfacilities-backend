package model

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/bayu-aditya/myfacilities-backend/lib/variable"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Base Model for MongoDB Document
type Base struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt *time.Time         `bson:"created_at,omitempty"`
	UpdatedAt *time.Time         `bson:"updated_at,omitempty"`

	tempQueryE bson.E
	findQueryD bson.D

	aggregateQueryA bson.A
}

// InitDateTimeForCreate for insert CreatedAt and UpdatedAt by datetime now
func (b *Base) InitDateTimeForCreate() {
	now := variable.DateTimeNow()
	b.CreatedAt = now
	b.UpdatedAt = now
}

func (b *Base) appendFindQueryD(query bson.E) {
	b.findQueryD = append(b.findQueryD, query)
}

// getFindQueryDAndClear method
func (b *Base) getFindQueryDAndClear() bson.D {
	result := b.findQueryD
	b.findQueryD = bson.D{}
	return result
}

// IsDocumentFound from error
func (b *Base) IsDocumentFound(err error) bool {
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
		log.Panicf("error model.Base.IsDocumentFound: %s \n", err)
	}
	return true
}
