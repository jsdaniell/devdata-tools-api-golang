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
