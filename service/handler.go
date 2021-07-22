package service

import (
	"encoding/json"

	"github.com/akbar-budiman/personal-playground-2/entity"
	"github.com/akbar-budiman/personal-playground-2/es"
)

func AddOrReplaceUser(user *entity.User) {
	addUserToRedis(user)
	addOrReplaceUserToCockroachDb(user)
	addOrReplaceUserToElasticSearch(user)
}

func addUserToRedis(user *entity.User) {
	uByte, _ := json.Marshal(user)
	err := SetValue(user.Name, uByte)
	if err != nil {
		panic(err)
	}
}

func addOrReplaceUserToCockroachDb(user *entity.User) {
	userFound := getCertainUserFromCockroachDb(user.Name)
	if userFound.Name != "" {
		replaceUserInCockroachDb(userFound.Name, user)
	} else {
		addUserToCockroachDb(user)
	}
}

func addOrReplaceUserToElasticSearch(user *entity.User) {
	esUser := user.NewEsUser()
	es.InsertData(esUser)
}

func replaceUserInCockroachDb(name string, user *entity.User) {
	UpdateUser(name, user)
}

func addUserToCockroachDb(user *entity.User) {
	InsertUser(user)
}

func GetCertainUserHandler(name string) *entity.User {
	foundUserRedis := getCertainUserFromRedis(name)
	if foundUserRedis != nil {
		var foundUser entity.User
		errJson := json.Unmarshal(foundUserRedis, &foundUser)
		if errJson != nil {
			panic(errJson)
		}
		return &foundUser
	}

	foundUser := getCertainUserFromCockroachDb(name)

	if foundUser.Name != "" {
		addUserToRedis(foundUser)
	}

	return foundUser
}

func getCertainUserFromRedis(name string) []byte {
	data, _ := GetValue(name)
	if len(data) > 0 {
		return data
	} else {
		return nil
	}
}

func getCertainUserFromCockroachDb(name string) *entity.User {
	usersFound := GetUser(name)
	if len(usersFound) > 0 {
		return &usersFound[0]
	} else {
		return &entity.User{}
	}
}

func GetCertainUsersBySearchKey(searchKey *string) []*entity.User {
	esUsers := es.FindDataBySearchKey(*searchKey)

	var users = []*entity.User{}
	for _, esUser := range esUsers {
		user := GetCertainUserHandler(esUser.Name)
		users = append(users, user)
	}

	return users
}
