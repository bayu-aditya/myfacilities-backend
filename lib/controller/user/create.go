package user

import (
	"context"

	gmodel "github.com/bayu-aditya/myfacilities-backend/graph/model"
	"github.com/bayu-aditya/myfacilities-backend/lib/model"
)

// Create user controller
func Create(ctx context.Context, input *gmodel.NewUser) (*gmodel.User, error) {
	user := new(model.User)

	user.ConvertFromGraphModel(input)

	_, err := user.Create()
	if err != nil {
		return nil, err
	}

	return user.Convert2GraphModel(), nil
}
