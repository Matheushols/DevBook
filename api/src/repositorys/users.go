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

// SerachById get an specific user by the id in database
func (repository users) SearchById(ID uint64) (models.User, error) {
	lines, erro := repository.db.Query(
		"select id, name, nick, email, createdAt from users where id = ?",
		ID,
	)
	if erro != nil {
		return models.User{}, erro
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

// EditUsersRepository get an specific user by the id in database and edit it
func (repository users) EditUsersRepository(ID uint64) (models.User, error) {
	lines, erro := repository.db.Query(
		"update id, name, nick, email, createdAt from users where id = ?",
		ID,
	)
	if erro != nil {
		return models.User{}, erro
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

// Update changes the informations about some user
func (repository users) Update(ID uint64, user models.User) error {
	statment, erro := repository.db.Prepare("update users set name = ?, nick = ?, email = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statment.Close()

	if _, erro = statment.Exec(user.Name, user.Nick, user.Email, ID); erro != nil {
		return erro
	}

	return nil
}

// Delete remove the informations about some user
func (repository users) Delete(ID uint64) error {
	statment, erro := repository.db.Prepare("delete from users where id = ?")
	if erro != nil {
		return erro
	}
	defer statment.Close()

	if _, erro = statment.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// SearchByEmail search an user by email and return his id and password hash
func (repository users) SearchByEmail(email string) (models.User, error) {
	line, erro := repository.db.Query("select id, password from users where email = ?", email)
	if erro != nil {
		return models.User{}, erro
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if erro = line.Scan(&user.ID, &user.Password); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}
