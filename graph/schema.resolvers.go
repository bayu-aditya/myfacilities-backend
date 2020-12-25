package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/bayu-aditya/myfacilities-backend/graph/generated"
	"github.com/bayu-aditya/myfacilities-backend/graph/model"
	AuthCtrl "github.com/bayu-aditya/myfacilities-backend/lib/controller/auth"
	OrganizationCtrl "github.com/bayu-aditya/myfacilities-backend/lib/controller/organization"
	UserCtrl "github.com/bayu-aditya/myfacilities-backend/lib/controller/user"
	"github.com/bayu-aditya/myfacilities-backend/lib/middleware"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return UserCtrl.Create(ctx, &input)
}

func (r *mutationResolver) CreateOrganization(ctx context.Context, input model.NewOrganization) (*model.Organization, error) {
	if err := middleware.AuthorizationRequired(ctx); err != nil {
		return nil, err
	}

	return OrganizationCtrl.Create(ctx, &input)
}

func (r *mutationResolver) EditOrganization(ctx context.Context, input model.EditOrganization) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteOrganization(ctx context.Context, id string) (*model.ResponseCode, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InviteUserOrganization(ctx context.Context, id string, userID string, role model.RoleUserInOrganization) (*model.Organization, error) {
	if err := middleware.AuthorizationRequired(ctx); err != nil {
		return nil, err
	}

	return OrganizationCtrl.InviteUser(ctx, &id, &userID, &role)
}

func (r *mutationResolver) ChangeRoleUserOrganization(ctx context.Context, id string, userID string, role model.RoleUserInOrganization) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Login(ctx context.Context, input model.Login) (*model.LoginResult, error) {
	return AuthCtrl.Login(ctx, &input)
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	if err := middleware.AuthorizationRequired(ctx); err != nil {
		return nil, err
	}

	return UserCtrl.Get(ctx)
}

func (r *queryResolver) UserByEmail(ctx context.Context, email string) (*model.User, error) {
	if err := middleware.AuthorizationRequired(ctx); err != nil {
		return nil, err
	}

	return UserCtrl.GetByEmail(ctx, &email)
}

func (r *queryResolver) Organizations(ctx context.Context) ([]*model.Organization, error) {
	if err := middleware.AuthorizationRequired(ctx); err != nil {
		return nil, err
	}

	return OrganizationCtrl.Get(ctx)
}

func (r *queryResolver) Organization(ctx context.Context, id string) (*model.Organization, error) {
	if err := middleware.AuthorizationRequired(ctx); err != nil {
		return nil, err
	}

	return OrganizationCtrl.GetByID(ctx, &id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
