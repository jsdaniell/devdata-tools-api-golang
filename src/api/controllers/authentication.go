package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/jsdaniell/devdata-tools-api-golang/api/models"
	"github.com/jsdaniell/devdata-tools-api-golang/api/repository/user_repository"
	"github.com/jsdaniell/devdata-tools-api-golang/api/responses"
	"github.com/jsdaniell/devdata-tools-api-golang/api/utils/cors"
	"io/ioutil"
	"net/http"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {

	cors.EnableCors(&w)

	cors.SetupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	var userReceived models.User

	var user models.User

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = json.Unmarshal(bytes, &userReceived)
	if err != nil {
		fmt.Println("Error when unmarshal userBody")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err = user_repository.GetUserByUid(userReceived.Uid)
	if err != nil {
		fmt.Println(err)
	}

	if (models.User{}) == user {
		newUser, err := user_repository.CreateNewUserFromLogin(userReceived)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}

		w.WriteHeader(200)
		w.Header().Set("Content-Type", "text/plain")

		w.Write([]byte(newUser.Uid))

	} else {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(user.Uid))
	}
}
