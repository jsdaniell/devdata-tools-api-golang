package user_repository

import (
	"encoding/json"
	"fmt"
	"github.com/jsdaniell/devdata-tools-api-golang/api/db"
	"github.com/jsdaniell/devdata-tools-api-golang/api/models"
	"github.com/jsdaniell/devdata-tools-api-golang/api/utils/json_utility"
	"golang.org/x/net/context"
)

func loggingAndReturningError(message string) (models.User, error) {
	//logger.LogUser.Println(message)
	return models.User{}, fmt.Errorf(message)
}

func GetUserByUid(uid string) (models.User, error) {

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
		var us models.User

		err := json.Unmarshal(parsed, &us)
		if err != nil {
			return loggingAndReturningError(err.Error())
		}

		return us, nil
	}
}

func CreateNewUserFromLogin(newUser models.User) (models.User, error) {

	client := db.FirestoreClient()
	defer client.Close()

	//logger.LogUser.Println("Creating new user from: " + newUser.DisplayName + " | " + newUser.ApiKey + " | " + newUser.Email)
	fmt.Println("Creating new user from: " + newUser.DisplayName + " | " + newUser.ApiKey + " | " + newUser.Email)

	jsonLowerCase, err := json_utility.StructToLowerCaseJson(newUser)
	if err != nil {
		return models.User{}, err
	}

	_, err = client.Collection("users").Doc(newUser.Uid).Set(context.Background(), jsonLowerCase)
	if err != nil {
		return models.User{}, err
	} else {

		var newUs models.User

		newUserFromFirestore, err := client.Collection("users").Doc(newUser.Uid).Get(context.Background())
		if err != nil {
			return loggingAndReturningError(err.Error())
		}

		parsedNewUser, err := json.Marshal(newUserFromFirestore.Data())
		if err != nil {
			return models.User{}, err
		}

		err = json.Unmarshal(parsedNewUser, &newUs)
		if err != nil {
			return models.User{}, err
		}

		return newUs, nil
	}
}
