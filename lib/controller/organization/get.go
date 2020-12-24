package organization

import (
	"context"
	"errors"
	"net/http"

	gmodel "github.com/bayu-aditya/myfacilities-backend/graph/model"
	"github.com/bayu-aditya/myfacilities-backend/lib/middleware"
	"github.com/bayu-aditya/myfacilities-backend/lib/model"
)

// GetByID organization
// TODO only show organization where this user contain in this organization
func GetByID(ctx context.Context, orgID *string) (*gmodel.Organization, error) {
	c, _ := middleware.GinContextFromContext(ctx)
	claims := middleware.GetClaims(ctx)

	user := new(model.User)
	if found := user.FindByID(claims.UserID); found == false {
		c.AbortWithStatus(http.StatusNotFound)
		return nil, errors.New("user not found")
	}

	organization := new(model.Organization)
	if found := organization.FindByID(*orgID); found == false {
		c.AbortWithStatus(http.StatusNotFound)
		return nil, errors.New("organization not found")
	}

	return organization.Convert2GraphModel(), nil
}
