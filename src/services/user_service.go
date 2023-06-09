package services

import (
	"database/sql"
	"fmt"

	"github.com/web-service-gin/src/interfaces"
	"github.com/web-service-gin/src/util"
)

type User interfaces.User

func findExistingEmail(email string) (bool, error) {
	var user User
	row := util.DB.QueryRow("SELECT * FROM user WHERE email = ?", email)

	if err := row.Scan(&user.Name, &user.Email, &user.Password, &user.RegisteredAt, &user.LastLogin); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
	}

	return false, fmt.Errorf("user email already in use %v", email)
}

func RegisterUser(user User) (int64, error) {
	emailInUser, err := findExistingEmail(user.Email)
	if err != nil {
		return 0, err
	}

	if !emailInUser {
		return 0, fmt.Errorf("email in use")
	}

	result, err := util.DB.Exec("INSERT INTO user (name, email, password, registeredAt, lastLogin) VALUES (?, ?, ?, ?, ?)", user.Name, user.Email, user.Password, user.RegisteredAt, user.LastLogin)
	if err != nil {
		return 0, fmt.Errorf("failed to insert user")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("add user %v", user)
	}

	return id, nil
}
