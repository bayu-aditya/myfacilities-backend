package model

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/bayu-aditya/myfacilities-backend/lib/tools"

	gmodel "github.com/bayu-aditya/myfacilities-backend/graph/model"
	"github.com/bayu-aditya/myfacilities-backend/lib/service"
	"github.com/bayu-aditya/myfacilities-backend/lib/variable"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// User Model
type User struct {
	Base `bson:",inline"`

	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

// Collection for user document
func (User) Collection() *mongo.Collection {
	return service.MongoDB.Collection(variable.Collection.User)
}

// FindByID and return status found
func (u *User) FindByID(id string) (found bool) {
	empty := User{}

	objectID, _ := primitive.ObjectIDFromHex(id)
	u.Collection().FindOne(context.Background(), bson.M{"_id": objectID}).Decode(u)
	return !(*u == empty)
}

// FindByEmail and return status found
func (u *User) FindByEmail(email string) (found bool) {
	empty := User{}

	u.Collection().FindOne(context.Background(), bson.M{"email": email}).Decode(u)
	return !(*u == empty)
}

// Create User to MongoDB
func (u *User) Create() (*mongo.InsertOneResult, error) {
	u.InitDateTimeForCreate()

	res, err := u.Collection().InsertOne(context.Background(), *u)
	if err != nil {
		return nil, err
	}

	u.ID = res.InsertedID.(primitive.ObjectID)
	return res, nil
}

// ConvertFromGraphModel for graphQL
func (u *User) ConvertFromGraphModel(input interface{}) {
	switch v := input.(type) {
	case *gmodel.NewUser:
		u.Name = v.Name
		u.Email = v.Email
		u.Password = tools.PasswordFromPlainText(v.Password)
	default:
		log.Fatalf("User.ConvertFromGraphModel, wrong type %T", v)
	}
	return
}

// Convert2GraphModel for graphQL
func (u *User) Convert2GraphModel() *gmodel.User {
	return &gmodel.User{
		ID:    u.ID.Hex(),
		Name:  u.Name,
		Email: u.Email,
	}
}
