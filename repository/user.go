package repository

import (
	"database/sql"
	"fmt"
	"github.com/model"
)

type User struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *User {
	return &User{db}
}

func (r User) Create(user model.User) (uint64, error) {
	statement, error := r.db.Prepare("insert into user (name, nick, email, password) values (?, ?, ?, ?)")
	if error != nil {
		return 0, error
	}
	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if error != nil {
		return 0, error
	}

	lastInsertId, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}

	return uint64(lastInsertId), nil
}

func (r User) Update(user model.User) error {
	statement, error := r.db.Prepare("update user set name = ?, nick = ?, email = ?, password = ? where id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.Password, user.Id)
	if error != nil {
		return error
	}

	if _, error := result.RowsAffected(); error != nil {
		return error
	}

	return nil
}

func (r User) FindAll(nameOrNick string) ([]model.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	rows, error := r.db.Query(
		"select id, name, nick, email, password, created_at from user where name like ? or nick like ?",
		nameOrNick,
		nameOrNick,
	)

	if error != nil {
		return nil, error
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var user model.User

		if error = scan(rows, &user); error != nil {
			return nil, error
		}

		users = append(users, user)
	}

	return users, nil
}

func (r User) FindById(id uint64) (model.User, error) {
	rows, error := r.db.Query("select id, name, nick, email, password, created_at from user where id = ?", id)
	if error != nil {
		return model.User{}, error
	}
	defer rows.Close()

	var user model.User

	if rows.Next() {
		if error = scan(rows, &user); error != nil {
			return model.User{}, error
		}
	}

	return user, nil
}

func (r User) DeleteById(id uint64) error {
	statement, error := r.db.Prepare("delete from user where id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error := statement.Exec(id); error != nil {
		return error
	}

	return nil
}

func (r User) FindByEmail(email string) (model.User, error) {
	rows, error := r.db.Query("select id, password from user where email = ?", email)
	if error != nil {
		return model.User{}, error
	}
	defer rows.Close()

	var user model.User

	if rows.Next() {
		if error = rows.Scan(&user.Id, &user.Password); error != nil {
			return model.User{}, error
		}
	}

	return user, nil
}

func scan(rows *sql.Rows, user *model.User) error {
	return rows.Scan(&user.Id, &user.Name, &user.Nick, &user.Email, &user.Password, &user.CreatedAt)
}
