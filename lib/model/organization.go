package model

import (
	"context"
	"errors"
	"log"

	"github.com/bayu-aditya/myfacilities-backend/lib/tools"

	gmodel "github.com/bayu-aditya/myfacilities-backend/graph/model"
	"github.com/bayu-aditya/myfacilities-backend/lib/service"
	"github.com/bayu-aditya/myfacilities-backend/lib/variable"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// OrganizationQuery for MongoDB
type OrganizationQuery struct {
	model Organization
}

// GetOrganizations for this user
// either as admin or member
func (q *OrganizationQuery) GetOrganizations(user *User) []*Organization {
	var result []*Organization

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cur, err := q.model.Collection().Find(
		ctx,
		bson.D{
			{Key: "$or", Value: bson.A{
				bson.M{"admins": bson.M{
					"$in": bson.A{user.ID},
				}},
				bson.M{"members": bson.M{
					"$in": bson.A{user.ID},
				}},
			}},
		},
	)
	if err != nil {
		log.Panicf("Error model.OrganizationQuery.GetOrganizations: %s \n", err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var org Organization
		cur.Decode(&org)
		result = append(result, &org)
	}
	return result
}

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

// AddAdmin targetUser by adminUser
func (o *Organization) AddAdmin(adminUser *User, targetUser *User) error {
	if o.IsAdmin(adminUser) == false {
		return errors.New("you're not admin in this organization")
	}

	return o.Collection().FindOneAndUpdate(
		context.Background(),
		bson.M{"_id": o.ID},
		bson.M{"$push": bson.M{"admins": targetUser.ID}},
	).Decode(o)
}

// IsAdmin checker for user in this organization
func (o *Organization) IsAdmin(user *User) bool {
	err := o.Collection().FindOne(
		context.Background(),
		bson.D{
			{Key: "_id", Value: o.ID},
			{Key: "admins", Value: bson.M{
				"$in": bson.A{user.ID},
			}},
		},
	).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
		log.Println("Error model.Organization.IsAdmin")
	}

	return true
}

// IsUserContain this organization?
func (o *Organization) IsUserContain(user *User) bool {
	err := o.Collection().FindOne(
		context.Background(),
		bson.D{
			{Key: "_id", Value: o.ID},
			{Key: "$or", Value: bson.A{
				bson.M{"admins": bson.M{
					"$in": bson.A{user.ID},
				}},
				bson.M{"members": bson.M{
					"$in": bson.A{user.ID},
				}},
			}},
		},
	).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
		log.Println("Error model.Organization.IsUserContain")
	}

	return true
}

// Convert2GraphModel for graphql
func (o *Organization) Convert2GraphModel() *gmodel.Organization {
	return &gmodel.Organization{
		ID:      o.ID.Hex(),
		Name:    o.Name,
		Admins:  tools.ArrayObjectID2ArrayString(o.Admins),
		Members: tools.ArrayObjectID2ArrayString(o.Members),
	}
}
