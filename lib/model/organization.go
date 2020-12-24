package model

import (
	"context"

	"github.com/bayu-aditya/myfacilities-backend/lib/service"
	"github.com/bayu-aditya/myfacilities-backend/lib/variable"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Organization Model
type Organization struct {
	Base `bson:",inline"`

	Name    string               `bson:"name"`
	Admins  []primitive.ObjectID `bson:"admins"`
	Members []primitive.ObjectID `bson:"members"`
}

// Collection of organization
func (Organization) Collection() *mongo.Collection {
	return service.MongoDB.Collection(variable.Collection.Organization)
}

// Create organization by user
func (o *Organization) Create(user *User) error {
	o.InitDateTimeForCreate()

	o.Admins = append(o.Admins, user.ID)

	res, err := o.Collection().InsertOne(context.Background(), *o)
	if err != nil {
		return err
	}

	o.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
