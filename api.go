package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterUserRouter() {
	fmt.Println("Registering router")

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/users", addUserEvent).Methods("POST")
	myRouter.HandleFunc("/users/direct", addUserDirect).Methods("POST")
	myRouter.HandleFunc("/users/{name}", getCertainUser).Methods("GET")

	myRouter.Use(printUri)
	fmt.Println("Router registered.")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func addUserEvent(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	ProduceNewUserEvent(reqBody)
	w.WriteHeader(http.StatusCreated)
}

func addUserDirect(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)

	AddOrReplaceUser(&user)
	w.WriteHeader(http.StatusCreated)
}

func getCertainUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	userFound := GetCertainUserHandler(name)
	if userFound.Name != "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(userFound)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func printUri(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method + " " + r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
