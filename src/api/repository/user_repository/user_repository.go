package user_repository

import (
	"encoding/json"
	"fmt"
	"github.com/jsdaniell/devdata-tools-api-golang/api/db"
	"github.com/jsdaniell/devdata-tools-api-golang/api/models/database_models"
	"golang.org/x/net/context"
)

func loggingAndReturningError(message string) (database_models.User, error) {
	//logger.LogUser.Println(message)
	return database_models.User{}, fmt.Errorf(message)
}

func GetUserByUid(uid string) (database_models.User, error) {

	client := db.FirestoreClient()
	defer client.Close()

	users := client.Collection("users")

	user, err := users.Doc(uid).Get(context.Background())
	if err != nil {
		return loggingAndReturningError(err.Error())
	}

	parsed, err := json.Marshal(user.Data())
	if err != nil {
		return loggingAndReturningError(err.Error())
	} else {
		var us database_models.User

		err := json.Unmarshal(parsed, &us)
		if err != nil {
			return loggingAndReturningError(err.Error())
		}

		return us, nil
	}
}

func CreateNewUserFromLogin(newUser database_models.User) (database_models.User, error) {

	client := db.FirestoreClient()
	defer client.Close()

	//logger.LogUser.Println("Creating new user from: " + newUser.DisplayName + " | " + newUser.ApiKey + " | " + newUser.Email)
	fmt.Println("Creating new user from: " + newUser.DisplayName + " | " + newUser.ApiKey + " | " + newUser.Email)

	_, err := client.Collection("users").Doc(newUser.Uid).Set(context.Background(), newUser)
	if err != nil {
		return database_models.User{}, err
	} else {

		var newUs database_models.User

		newUserFromFirestore, err := client.Collection("users").Doc(newUser.Uid).Get(context.Background())
		if err != nil {
			return loggingAndReturningError(err.Error())
		}

		parsedNewUser, err := json.Marshal(newUserFromFirestore.Data())
		if err != nil {
			return database_models.User{}, err
		}

		err = json.Unmarshal(parsedNewUser, &newUs)
		if err != nil {
			return database_models.User{}, err
		}

		return newUs, nil
	}
}
