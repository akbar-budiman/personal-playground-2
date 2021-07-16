package main

import (
	"encoding/json"
	"fmt"
)

func AddOrReplaceUser(user *User) {
	addUserToRedis(user)
	addOrReplaceUserToCockroachDb(user)
}

func addUserToRedis(user *User) {
	uByte, _ := json.Marshal(user)
	err := SetValue(user.Name, uByte)
	if err != nil {
		panic(err)
	}
}

func addOrReplaceUserToCockroachDb(user *User) {
	userFound := getCertainUserFromCockroachDb(user.Name)
	fmt.Println("userFound:", userFound)
	if userFound.Name != "" {
		replaceUserInCockroachDb(userFound.Name, user)
	} else {
		addUserToCockroachDb(user)
	}
}

func replaceUserInCockroachDb(name string, user *User) {
	UpdateUser(name, user)
}

func addUserToCockroachDb(user *User) {
	InsertUser(user)
}

func GetCertainUserHandler(name string) User {
	foundUserRedis := getCertainUserFromRedis(name)
	if foundUserRedis != nil {
		var foundUser User
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

func getCertainUserFromCockroachDb(name string) User {
	usersFound := GetUser(name)
	fmt.Println("len(rowsFound):", len(usersFound))
	if len(usersFound) > 0 {
		return usersFound[0]
	} else {
		return User{}
	}
}
