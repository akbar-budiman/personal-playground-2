package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"

	"github.com/akbar-budiman/personal-playground-2/entity"
	"github.com/akbar-budiman/personal-playground-2/graph/generated"
	"github.com/akbar-budiman/personal-playground-2/graph/model"
	"github.com/akbar-budiman/personal-playground-2/service"
)

func (r *mutationResolver) AddUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	userEntity := entity.User{
		Name:       input.Name,
		Age:        input.Age,
		Address:    *input.Address,
		Searchable: *input.Searchable,
	}

	userEntityByte, _ := json.Marshal(userEntity)
	service.ProduceNewUserEvent(userEntityByte)

	response := model.User(input)
	return &response, nil
}

func (r *queryResolver) User(ctx context.Context, name *string) (*model.User, error) {
	userFound := service.GetCertainUserHandler(*name)

	if userFound.Name != "" {
		return userFound.NewModelUser(), nil
	} else {
		return &model.User{}, nil
	}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
