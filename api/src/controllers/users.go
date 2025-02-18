package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositorys"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
)

// The CreateUser is used to create an user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating User!"))
	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare(); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositorys.NewUsersRepository(db)
	user.ID, erro = repository.Create(user)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, user)

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
