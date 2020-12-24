package model

import (
	"context"
	"log"

	"github.com/bayu-aditya/myfacilities-backend/lib/tools"

	gmodel "github.com/bayu-aditya/myfacilities-backend/graph/model"
	"github.com/bayu-aditya/myfacilities-backend/lib/service"
	"github.com/bayu-aditya/myfacilities-backend/lib/variable"
	"go.mongodb.org/mongo-driver/bson"
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

// FindByID for organization
func (o *Organization) FindByID(id string) (found bool) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	err := o.Collection().FindOne(context.Background(), bson.M{"_id": objectID}).Decode(o)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
		log.Panicln("Error: model.Organization.FindByID")
	}
	return true
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

// // Delete Organization
// func (o *Organization) Delete(user *User) error {
// 	ctx := context.Background()

// 	o.Collection().FindOne(ctx)
// }

// Convert2GraphModel for graphql
func (o *Organization) Convert2GraphModel() *gmodel.Organization {
	return &gmodel.Organization{
		ID:      o.ID.Hex(),
		Name:    o.Name,
		Admins:  tools.ArrayObjectID2ArrayString(o.Admins),
		Members: tools.ArrayObjectID2ArrayString(o.Members),
	}
}
