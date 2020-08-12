package controllers

import (
	"api/models"
	"api/repository"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"net/http"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	user, err := repository.GetUserByUid(r.Header["Authorization"][0])
	if err != nil {
		fmt.Println(err)
	} else {
		spew.Dump(user)
	}



	w.Write([]byte("Login User Controller"))
}
