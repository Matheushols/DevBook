package repositorys

import (
	"api/src/models"
	"database/sql"
	"log"
)

// Users represent a repository of users
type users struct {
	db *sql.DB
}

// NewUsersRepository create a repository of users
func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

// Create input an user on the database
func (repository users) Create(user models.User) (uint64, error) {
	statement, erro := repository.db.Prepare("insert into users (name, nick, email, password) values (?, ?, ?, ?)")

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}

	lastidInserted, erro := result.LastInsertId()
	if erro != nil {
		log.Fatal(erro)
	}

	return uint64(lastidInserted), nil

}

// List is used to list the users
/*func (repository users) List(user models.User) (string, error) {
	request, erro := repository.db.Query("select * from users")

	if erro != nil {
		return "0", erro
	}
	defer request.Close()

	return request
}
*/
