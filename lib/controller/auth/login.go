package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/bayu-aditya/myfacilities-backend/lib/middleware"

	"github.com/bayu-aditya/myfacilities-backend/lib/tools"

	gmodel "github.com/bayu-aditya/myfacilities-backend/graph/model"
	"github.com/bayu-aditya/myfacilities-backend/lib/model"
)

// Login Controller
func Login(ctx context.Context, input *gmodel.Login) (*gmodel.User, error) {
	c, _ := middleware.GinContextFromContext(ctx)

	user := new(model.User)
	if err := user.FindByEmail(input.Email); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return nil, err
	}

	if user.Password != tools.PasswordFromPlainText(input.Password) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return nil, errors.New("Invalid Password")
	}
	return user.Convert2GraphModel(), nil
}
