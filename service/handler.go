package service

import (
	"encoding/json"
	"fmt"

	"github.com/akbar-budiman/personal-playground-2/entity"
)

func AddOrReplaceUser(user *entity.User) {
	addUserToRedis(user)
	addOrReplaceUserToCockroachDb(user)
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
	fmt.Println("userFound:", userFound)
	if userFound.Name != "" {
		replaceUserInCockroachDb(userFound.Name, user)
	} else {
		addUserToCockroachDb(user)
	}
}

func replaceUserInCockroachDb(name string, user *entity.User) {
	UpdateUser(name, user)
}

func addUserToCockroachDb(user *entity.User) {
	InsertUser(user)
}

func GetCertainUserHandler(name string) entity.User {
	foundUserRedis := getCertainUserFromRedis(name)
	if foundUserRedis != nil {
		var foundUser entity.User
		errJson := json.Unmarshal(foundUserRedis, &foundUser)
		if errJson != nil {
			panic(errJson)
		}
		return foundUser
	}

	foundUser := getCertainUserFromCockroachDb(name)

	if foundUser.Name != "" {
		addUserToRedis(&foundUser)
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

func getCertainUserFromCockroachDb(name string) entity.User {
	usersFound := GetUser(name)
	fmt.Println("len(rowsFound):", len(usersFound))
	if len(usersFound) > 0 {
		return usersFound[0]
	} else {
		return entity.User{}
	}
}
