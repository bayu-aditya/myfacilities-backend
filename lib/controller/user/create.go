package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/bayu-aditya/myfacilities-backend/lib/middleware"

	gmodel "github.com/bayu-aditya/myfacilities-backend/graph/model"
	"github.com/bayu-aditya/myfacilities-backend/lib/model"
)

// Create user controller
func Create(ctx context.Context, input *gmodel.NewUser) (*gmodel.User, error) {
	c, _ := middleware.GinContextFromContext(ctx)
	user := new(model.User)

	if found := user.FindByEmail(input.Email); found == true {
		c.AbortWithStatus(http.StatusBadRequest)
		return nil, errors.New("email has been registered")
	}

	user.ConvertFromGraphModel(input)

	_, err := user.Create()
	if err != nil {
		return nil, err
	}

	return user.Convert2GraphModel(), nil
}
