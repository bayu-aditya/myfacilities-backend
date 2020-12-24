package user

import (
	"context"

	gmodel "github.com/bayu-aditya/myfacilities-backend/graph/model"
)

func Create(ctx context.Context, input *gmodel.NewUser) (*gmodel.User, error) {
	return &gmodel.User{
		ID:    "13",
		Name:  input.Name,
		Email: input.Email,
	}, nil
}
