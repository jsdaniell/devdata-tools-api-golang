package controllers

import (
	"encoding/json"
	"github.com/jsdaniell/devdata-tools-api-golang/api/models"
	"github.com/jsdaniell/devdata-tools-api-golang/api/repository/user_repository"
	"github.com/jsdaniell/devdata-tools-api-golang/api/responses"
	"io/ioutil"
	"net/http"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// TODO: Compare received json and return bad request if isn't

	var userReceived models.User

	var user models.User

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = json.Unmarshal(bytes, &userReceived)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user, err = user_repository.GetUserByUid(userReceived.Uid)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
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
