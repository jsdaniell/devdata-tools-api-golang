package repository

import (
	"api/db"
	"api/models"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
)

func GetUserByUid(uid string) (models.User, error) {

	client := db.FirestoreClient()

	users := client.Collection("users")

	user, err := users.Doc(uid).Get(context.Background())
	if err != nil {
		fmt.Println(err)
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

		return us, nil
	}
}
