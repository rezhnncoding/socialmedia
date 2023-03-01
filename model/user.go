package model

import (
	"fmt"
	"github.com/badoux/checkmail"
	"github.com/security"
	"strings"
	"time"
)

type User struct {
	Id        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) Prepare() error {
	if error := user.validate(); error != nil {
		return error
	}

	if error := user.sanitize(); error != nil {
		return error
	}

	return nil
}

func (user *User) validate() error {
	var validations []string

	if strings.TrimSpace(user.Name) == "" {
		validations = append(validations, "User name cannot be blank")
	}

	if strings.TrimSpace(user.Nick) == "" {
		validations = append(validations, "User nick cannot be blank")
	}

	if strings.TrimSpace(user.Email) == "" {
		validations = append(validations, "User email cannot be blank")
	}

	if error := checkmail.ValidateFormat(user.Email); error != nil {
		validations = append(validations, "User email cannot have invalid format")
	}

	if strings.TrimSpace(user.Password) == "" {
		validations = append(validations, "User password cannot be blank")
	}

	if len(validations) == 0 {
		return nil
	}

	return fmt.Errorf("constraint violations on user: [%s]", strings.Join(validations, ","))
}

func (user *User) sanitize() error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	passwordHash, error := security.Hash(user.Password)
	if error != nil {
		return error
	}

	user.Password = string(passwordHash)

	return nil
}
