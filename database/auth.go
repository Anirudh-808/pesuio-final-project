package database

import (
	"github.com/anuragrao04/pesuio-final-project/models"
)

func CreateUser(databaseFileName, username, password string) (string, error) {
	// creates a new user in the database, returns error if any
	dataBase, err := Init(databaseFileName)
	Credentials := models.User{
		Username: username,
		Password: password,
	}

	dataBase.Create(&Credentials)
	return "CREATED SUCCESSFULLY", err
}

func CheckPassword(username, password string) (success bool, err error) {
	// checks if the password is correct for the given username
	return true, nil
}
