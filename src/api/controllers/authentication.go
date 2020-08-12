package controllers

import (
	"api/models"
	"api/repository/user_repository"
	"config/logger"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"net/http"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	user, err := user_repository.GetUserByUid(r.Header["Authorization"][0])
	if err != nil {
		logger.LogUser.Println(err)
		fmt.Println(err)
	}

	spew.Dump(user)

	if (models.User{}) == user {
		// TODO: Create a new user
		fmt.Println("NewUser")
	} else {
		js, err := json.Marshal(user)
		if err != nil {
			fmt.Errorf("error on unmarshal user struct json")
		}

		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}



}
