// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewUser struct {
	Name       string  `json:"name"`
	Age        int     `json:"age"`
	Address    *string `json:"address"`
	Searchable *string `json:"searchable"`
}

type User struct {
	Name       string  `json:"name"`
	Age        int     `json:"age"`
	Address    *string `json:"address"`
	Searchable *string `json:"searchable"`
}
