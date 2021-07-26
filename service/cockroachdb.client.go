package service

import (
	"github.com/akbar-budiman/personal-playground-2/entity"
)

type CrdbClient interface {
	GetUser(name string) []entity.User
	InsertUser(user *entity.User)
	UpdateUser(name string, user *entity.User)
}

type CrdbClientImpl struct {
}

func (crdb *CrdbClientImpl) GetUser(name string) []entity.User {
	return GetUser(name)
}

func (crdb *CrdbClientImpl) InsertUser(user *entity.User) {
	InsertUser(user)
}

func (crdb *CrdbClientImpl) UpdateUser(name string, user *entity.User) {
	UpdateUser(name, user)
}
