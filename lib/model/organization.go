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

// MultipleOrganizations for MongoDB
type MultipleOrganizations struct {
	model Organization
}

// GetByUser for this user
// either as admin or member
func (q *MultipleOrganizations) GetByUser(user *User) []*Organization {
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
		log.Panicf("Error model.MultipleOrganizations.GetByUser: %s \n", err)
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

// Where for create single bsonE query
func (o *Organization) Where(key string) *Organization {
	o.tempQueryE.Key = key
	return o
}

// Is for Where method
func (o *Organization) Is(value interface{}) *Organization {
	o.tempQueryE.Value = value

	o.appendFindQueryD(o.tempQueryE)
	o.tempQueryE = bson.E{}
	return o
}

// FindByID for organization
func (o *Organization) FindByID(id string) (found bool) {
	objectID, _ := primitive.ObjectIDFromHex(id)

	o.Where("_id").Is(objectID)
	err := o.Collection().FindOne(context.Background(), o.getFindQueryDAndClear()).Decode(o)
	return o.IsDocumentFound(err)
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
		bson.M{"$addToSet": bson.M{"admins": targetUser.ID}},
	).Err()
}

// IsAdmin checker for user in this organization
func (o *Organization) IsAdmin(user *User) bool {
	o.Where("_id").Is(o.ID)
	o.Where("admins").Is(bson.M{"$in": bson.A{user.ID}})

	err := o.Collection().FindOne(context.Background(), o.getFindQueryDAndClear()).Err()
	return o.IsDocumentFound(err)
}

// IsMember checker for user in this organization
func (o *Organization) IsMember(user *User) bool {
	o.Where("_id").Is(o.ID)
	o.Where("members").Is(bson.M{"$in": bson.A{user.ID}})

	err := o.Collection().FindOne(context.Background(), o.getFindQueryDAndClear()).Err()
	return o.IsDocumentFound(err)
}

// IsUserContain this organization?
func (o *Organization) IsUserContain(user *User) bool {
	o.Where("_id").Is(o.ID)
	o.Where("$or").Is(bson.A{
		bson.M{"admins": bson.M{"$in": bson.A{user.ID}}},
		bson.M{"members": bson.M{"$in": bson.A{user.ID}}},
	})

	err := o.Collection().FindOne(context.Background(), o.getFindQueryDAndClear()).Err()
	return o.IsDocumentFound(err)
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
