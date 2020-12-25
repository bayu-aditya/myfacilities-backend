package organization

import (
	"context"
	"errors"
	"net/http"

	gmodel "github.com/bayu-aditya/myfacilities-backend/graph/model"
	"github.com/bayu-aditya/myfacilities-backend/lib/middleware"
	"github.com/bayu-aditya/myfacilities-backend/lib/model"
)

// Get all organization for this user
func Get(ctx context.Context) ([]*gmodel.Organization, error) {
	c, _ := middleware.GinContextFromContext(ctx)
	claims := middleware.GetClaims(ctx)

	user := new(model.User)
	if found := user.FindByID(claims.UserID); found == false {
		c.AbortWithStatus(http.StatusNotFound)
		return nil, errors.New("user not found")
	}

	query := new(model.OrganizationQuery)
	organizations := query.GetOrganizations(user)

	var resp []*gmodel.Organization
	for _, org := range organizations {
		resp = append(resp, org.Convert2GraphModel())
	}
	return resp, nil
}

// GetByID organization
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

	if organization.IsUserContain(user) == false {
		c.AbortWithStatus(http.StatusForbidden)
		return nil, errors.New("you dont have permission to view this organization")
	}

	return organization.Convert2GraphModel(), nil
}
