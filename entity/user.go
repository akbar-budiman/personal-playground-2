package entity

import "github.com/akbar-budiman/personal-playground-2/graph/model"

type User struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Address    string `json:"address"`
	Searchable string `json:"searchable"`
}

type EsUser struct {
	Name       string `json:"name"`
	Searchable string `json:"searchable"`
}

func (e *User) NewModelUser() *model.User {
	return &model.User{
		Name:       e.Name,
		Age:        e.Age,
		Address:    &e.Address,
		Searchable: &e.Searchable,
	}
}

func (e *User) NewEsUser() *EsUser {
	return &EsUser{
		Name:       e.Name,
		Searchable: e.Searchable,
	}
}
