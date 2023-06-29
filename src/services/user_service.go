package services

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/web-service-gin/src/interfaces"
	"github.com/web-service-gin/src/util"
)

func findExistingEmail(email string) (bool, error) {
	var user interfaces.User
	row := util.DB.QueryRow("SELECT * FROM user WHERE email = ?", email)

	if err := row.Scan(&user.Name, &user.Email, &user.Password, &user.RegisteredAt, &user.LastLogin); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
	}

	return false, fmt.Errorf("user email already in use %v", email)
}

func RegisterUser(user interfaces.User) (int64, error) {
	emailInUser, err := findExistingEmail(user.Email)
	if err != nil {
		return 0, err
	}

	if emailInUser {
		return 0, fmt.Errorf("email in use")
	}
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return 0, fmt.Errorf("could not hash password")
	}
	result, err := util.DB.Exec("INSERT INTO user (name, email, password, registeredAt, lastLogin) VALUES (?, ?, ?, NOW(), NOW())", user.Name, user.Email, hashedPassword)
	if err != nil {
		return 0, fmt.Errorf("failed to insert user")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("add user %v", user)
	}

	return id, nil
}

func LoginUser(email string, password string) (*interfaces.User, error) {
	var user interfaces.User
	row := util.DB.QueryRow("SELECT * FROM user WHERE email = ?", email)

	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RegisteredAt, &user.LastLogin); err != nil {
		if err == sql.ErrNoRows {
			return &user, fmt.Errorf("user with email %s does not exist, %v", email, err)
		}
		return &user, fmt.Errorf("user with email %s, %v", email, err)
	}

	err := isValidHash(user.Password, password)
	if err != nil {
		return &interfaces.User{}, fmt.Errorf("incorrect password: %v", err)
	}

	res := &interfaces.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		LastLogin: user.LastLogin,
	}

	return res, nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hash), err
}

func isValidHash(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
