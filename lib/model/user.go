package model

import (
	"time"

	gmodel "github.com/bayu-aditya/myfacilities-backend/graph/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User Model
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt *time.Time         `bson:"created_at,omitempty"`
	UpdatedAt *time.Time         `bson:"updated_at,omitempty"`

	Name  string `bson:"name"`
	Email string `bson:"email"`
}

// Convert2GraphModel for graphQL
func (u *User) Convert2GraphModel() *gmodel.User {
	return &gmodel.User{
		ID:    u.ID.Hex(),
		Name:  u.Name,
		Email: u.Email,
	}
}
