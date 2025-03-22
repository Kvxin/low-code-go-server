package service

import (
	"go_server/internal/model"
	"go_server/internal/repository/mongodb"
	"go_server/internal/util"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func CreateUser(email, password string) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email:    email,
		Username: email[:strings.Index(email, "@")],
		Password: string(hashedPassword),
		IsActive: true,
		Role:     "user",
	}

	if err := mongodb.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func AuthenticateUser(email, password string) (*model.User, error) {
	user, err := mongodb.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
} 