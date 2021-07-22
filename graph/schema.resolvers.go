package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"

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

func (r *queryResolver) UserBySearchKey(ctx context.Context, searchKey *string) ([]*model.User, error) {
	users := service.GetCertainUsersBySearchKey(searchKey)

	var response = []*model.User{}
	for _, user := range users {
		response = append(response, user.NewModelUser())
	}
	return response, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) UserBySearchable(ctx context.Context, searchable *string) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
