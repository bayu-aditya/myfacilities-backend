package organization

import (
	"context"
	"errors"
	"net/http"

	gmodel "github.com/bayu-aditya/myfacilities-backend/graph/model"
	"github.com/bayu-aditya/myfacilities-backend/lib/middleware"
	"github.com/bayu-aditya/myfacilities-backend/lib/model"
)

// InviteUser to the organization
func InviteUser(ctx context.Context, orgID *string, targetUserID *string, role *gmodel.RoleUserInOrganization) (*gmodel.Organization, error) {
	c, _ := middleware.GinContextFromContext(ctx)
	claims := middleware.GetClaims(ctx)

	user := new(model.User)
	if found := user.FindByID(claims.UserID); found == false {
		c.AbortWithStatus(http.StatusNotFound)
		return nil, errors.New("user not found")
	}
	targetUser := new(model.User)
	if found := targetUser.FindByID(*targetUserID); found == false {
		c.AbortWithStatus(http.StatusNotFound)
		return nil, errors.New("target user not found")
	}

	organization := new(model.Organization)
	if found := organization.FindByID(*orgID); found == false {
		c.AbortWithStatus(http.StatusNotFound)
		return nil, errors.New("organization not found")
	}

	if err := organization.AddAdmin(user, targetUser); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return nil, err
	}
	organization.FindByID(*orgID)

	return organization.Convert2GraphModel(), nil
}
