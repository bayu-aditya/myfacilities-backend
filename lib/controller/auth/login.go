package auth

import (
	"context"
	"errors"
	"net/http"

	gmodel "github.com/bayu-aditya/myfacilities-backend/graph/model"
	"github.com/bayu-aditya/myfacilities-backend/lib/middleware"
	"github.com/bayu-aditya/myfacilities-backend/lib/model"
	"github.com/bayu-aditya/myfacilities-backend/lib/tools"
)

// Login Controller
func Login(ctx context.Context, input *gmodel.Login) (*gmodel.LoginResult, error) {
	c, _ := middleware.GinContextFromContext(ctx)

	user := new(model.User)
	if found := user.FindByEmail(input.Email); found == false {
		c.AbortWithStatus(http.StatusNotFound)
		return nil, errors.New("user not found")
	}

	if user.Password != tools.PasswordFromPlainText(input.Password) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return nil, errors.New("Invalid Password")
	}

	return &gmodel.LoginResult{
		Token: middleware.GenerateToken(user),
		User:  user.Convert2GraphModel(),
	}, nil
}
