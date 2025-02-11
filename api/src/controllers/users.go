package controllers

import "net/http"

// The CreateUser is used to create an user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating User!"))
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
