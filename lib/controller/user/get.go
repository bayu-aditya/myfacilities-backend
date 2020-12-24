package user

import (
	"context"
	"errors"
	"net/http"

	gmodel "github.com/bayu-aditya/myfacilities-backend/graph/model"
	"github.com/bayu-aditya/myfacilities-backend/lib/middleware"
	"github.com/bayu-aditya/myfacilities-backend/lib/model"
)

// Get user information
func Get(ctx context.Context) (*gmodel.User, error) {
	c, _ := middleware.GinContextFromContext(ctx)
	claims := middleware.GetClaims(ctx)

	user := new(model.User)
	if found := user.FindByID(claims.UserID); found == false {
		c.AbortWithStatus(http.StatusNotFound)
		return nil, errors.New("user not found")
	}

	return user.Convert2GraphModel(), nil
}

// GetByEmail for user information
func GetByEmail(ctx context.Context, email *string) (*gmodel.User, error) {
	c, _ := middleware.GinContextFromContext(ctx)

	user := new(model.User)
	if found := user.FindByEmail(*email); found == false {
		c.AbortWithStatus(http.StatusNotFound)
		return nil, errors.New("user not found")
	}

	return user.Convert2GraphModel(), nil
}
