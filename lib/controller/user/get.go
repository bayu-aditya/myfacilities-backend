package user

import (
	"context"

	gmodel "github.com/bayu-aditya/myfacilities-backend/graph/model"
)

func Get(ctx context.Context) (*gmodel.User, error) {
	return &gmodel.User{ID: "13", Name: "Bayu Aditya", Email: "test@gmail.com"}, nil
}
