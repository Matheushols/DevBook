package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositorys"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// The CreateUser is used to create an user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating User!"))
	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User
	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := database.Connect()
	if erro != nil {
		log.Fatal(erro)
	}

	repository := repositorys.NewUsersRepository(db)
	userId, erro := repository.Create(user)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("Id inserted: %d", userId)))

}

// The ListUsers is used to list all users
func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching all Users!"))
}

// The SearchUser is used to get an specific user
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching for a specific User!"))
}

// The EditUser is used to edit an specific user
func EditUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Editing an specific User!"))
}

// The DeleteUser is used to delete an specific user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting an specific User!"))
}
