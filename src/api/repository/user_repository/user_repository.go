package user_repository

import (
	"api/db"
	"api/models"
	"config/logger"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
)

func loggingAndReturningError(message string) (models.User, error) {
	logger.LogUser.Println(message)
	return models.User{}, fmt.Errorf(message)
}

func GetUserByUid(uid string) (models.User, error) {

	client := db.FirestoreClient()
	defer client.Close()

	users := client.Collection("users")

	user, err := users.Doc(uid).Get(context.Background())
	if err != nil {
		return loggingAndReturningError(`User for uid: "` + uid + `" does not exists.`)
	}

	parsed, err := json.Marshal(user.Data())
	if err != nil {
		return models.User{}, err
	} else {
		var us models.User

		err := json.Unmarshal(parsed, &us)
		if err != nil {
			return models.User{}, err
		}

		logger.LogUser.Println("User logged: " + us.DisplayName + " | " + us.ApiKey + " | " + us.Email)

		return us, nil
	}
}

func CreateNewUserFromLogin(newUser models.User) (models.User, error) {

	client := db.FirestoreClient()
	defer client.Close()

	logger.LogUser.Println("Creating new user from: " + newUser.DisplayName + " | " + newUser.ApiKey + " | " + newUser.Email)
	fmt.Println("Creating new user from: " + newUser.DisplayName + " | " + newUser.ApiKey + " | " + newUser.Email)

	_, err := client.Collection("users").Doc(newUser.Uid).Set(context.Background(), newUser)
	if err != nil {
		return models.User{}, fmt.Errorf("error on save new user")
	} else {

		var newUs models.User

		newUserFromFirestore, err := client.Collection("users").Doc(newUser.Uid).Get(context.Background())
		if err != nil {
			return loggingAndReturningError("error get new saved user: "  + newUser.DisplayName + " | " + newUser.ApiKey + " | " + newUser.Email)
		}

		parsedNewUser, err := json.Marshal(newUserFromFirestore.Data())
		if err != nil {
			return models.User{}, fmt.Errorf("error on marshal new user getted from firestore: " + string(parsedNewUser))
		}

		err = json.Unmarshal(parsedNewUser, &newUs)
		if err != nil {
			return models.User{}, fmt.Errorf("error on unmarshal new user getted from firestore: "+ string(parsedNewUser))
		}

		logger.LogUser.Println("User Created: " + newUs.DisplayName + " | " + newUs.ApiKey + " | " + newUs.Email)

		return newUs, nil

	}
}
