package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/bayu-aditya/myfacilities-backend/graph/generated"
	"github.com/bayu-aditya/myfacilities-backend/graph/model"
	AuthCtrl "github.com/bayu-aditya/myfacilities-backend/lib/controller/auth"
	UserCtrl "github.com/bayu-aditya/myfacilities-backend/lib/controller/user"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return UserCtrl.Create(ctx, &input)
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	return UserCtrl.Get(ctx)
}

func (r *queryResolver) Login(ctx context.Context, input model.Login) (*model.User, error) {
	return AuthCtrl.Login(ctx, &input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
