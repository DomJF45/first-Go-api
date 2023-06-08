package main

import (
	"database/sql"
	"fmt"
)

func findExistingEmail(email string) (bool, error) {
	var user User
	row := db.QueryRow("SELECT * FROM user WHERE email = ?", email)

	if err := row.Scan(&user.Name, &user.Email, &user.Password, &user.RegisteredAt, &user.LastLogin); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
	}

	return false, fmt.Errorf("User email already in use %d", email)
}

func registerUser(user User) (int64, error) {
	emailInUser, err := findExistingEmail(user.Email)
	if err != nil {
		return 0, err
	}

	if !emailInUser {
		return 0, fmt.Errorf("Email in use")
	}

	result, err := db.Exec("INSERT INTO user (name, email, password, registeredAt, lastLogin) VALUES (?, ?, ?, ?, ?)", user.Name, user.Email, user.Password, user.RegisteredAt, user.LastLogin)
	if err != nil {
		return 0, fmt.Errorf("Failed to insert user")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("add user %d", user)
	}

	return id, nil
}
