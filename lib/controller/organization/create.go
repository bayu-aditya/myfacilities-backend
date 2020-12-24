package organization

import (
	"context"
	"errors"
	"net/http"

	"github.com/bayu-aditya/myfacilities-backend/lib/middleware"
	"github.com/bayu-aditya/myfacilities-backend/lib/model"

	gmodel "github.com/bayu-aditya/myfacilities-backend/graph/model"
)

// Create Organization
func Create(ctx context.Context, input *gmodel.NewOrganization) (*gmodel.Organization, error) {
	c, _ := middleware.GinContextFromContext(ctx)
	claims := middleware.GetClaims(ctx)

	user := new(model.User)
	if found := user.FindByID(claims.UserID); found == false {
		c.AbortWithStatus(http.StatusNotFound)
		return nil, errors.New("user not found")
	}

	organization := &model.Organization{Name: input.Name}
	if err := organization.Create(user); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return nil, errors.New("cannot create organization")
	}

	return organization.Convert2GraphModel(), nil
}
