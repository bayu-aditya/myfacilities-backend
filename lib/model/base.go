package model

import (
	"time"

	"github.com/bayu-aditya/myfacilities-backend/lib/variable"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Base Model for MongoDB Document
type Base struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt *time.Time         `bson:"created_at,omitempty"`
	UpdatedAt *time.Time         `bson:"updated_at,omitempty"`
}

// InitDateTimeForCreate for insert CreatedAt and UpdatedAt by datetime now
func (b *Base) InitDateTimeForCreate() {
	now := variable.DateTimeNow()
	b.CreatedAt = now
	b.UpdatedAt = now
}
