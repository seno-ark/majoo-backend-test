package repository

import (
	"log"
	"majoo-backend-test/model"
)

func (r *Repository) GetAccount(userName string) (*model.Account, error) {

	var account = &model.Account{}

	query := "SELECT id, name, user_name, password FROM Users WHERE user_name = ?"
	err := r.DB.Get(account, query, userName)
	if err != nil {
		log.Println(err.Error())
		return account, err
	}

	return account, nil
}

func (r *Repository) GetUser(userID int) (*model.User, error) {

	var user = &model.User{}

	query := "SELECT id, name, user_name FROM Users WHERE id = ?"
	err := r.DB.Get(user, query, userID)
	if err != nil {
		log.Println(err.Error())
		return user, err
	}

	return user, nil
}
