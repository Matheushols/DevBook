package repositorys

import (
	"api/src/models"
	"database/sql"
	"fmt"
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
func (repository users) List(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	lines, erro := repository.db.Query(
		"select id, name, nick, email, createdAt from users where name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if erro != nil {
		return nil, erro
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}
		users = append(users, user)
	}

	return users, nil
}
