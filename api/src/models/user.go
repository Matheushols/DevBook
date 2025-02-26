package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User represent a user using the social media
type User struct {
	ID        uint64    `json:"id, omitempty"`
	Name      string    `json:"name, omitempty"`
	Nick      string    `json:"nick, omitempty"`
	Email     string    `json:"email, omitempty"`
	Password  string    `json:"password, omitempty"`
	CreatedAt time.Time `json:"createdAt, omitempty"`
}

// Prepare going to call the validation and prepare to gived user
func (user *User) Prepare(step string) error {
	if erro := user.validate(step); erro != nil {
		return erro
	}

	if erro := user.format(step); erro != nil {
		return erro
	}
	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "register" {
		passwordWithHash, erro := security.Hash(user.Password)
		if erro != nil {
			return erro
		}
		user.Password = string(passwordWithHash)
	}

	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("The Name is required and must not be blank")
	}

	if user.Nick == "" {
		return errors.New("The Nick is required and must not be blank")
	}

	if user.Email == "" {
		return errors.New("The Email is required and must not be blank")
	}

	if user.Name == "" {
		return errors.New("The name is required and must not be blank")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("The inputted email isn't valid.")
	}

	if step == "register" && user.Password == "" {
		return errors.New("The Password is required and must not be blank")
	}

	return nil
}
