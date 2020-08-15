package controllers

import (
	"github.com/gorilla/mux"
	"github.com/jsdaniell/devdata-tools-api-golang/api/repository/suite_repository"
	"github.com/jsdaniell/devdata-tools-api-golang/api/repository/user_repository"
	"github.com/jsdaniell/devdata-tools-api-golang/api/responses"
	"net/http"
)

func GetAllSuitesOfAType(w http.ResponseWriter, r *http.Request) {

	suiteType := mux.Vars(r)["type"]

	auth := r.Header.Get("Authorization")

	user, err := user_repository.GetUserByUid(auth)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	suites, err := suite_repository.GetAllSuites(user.Uid, suiteType)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, suites)
	return
}

func CreateNewSuite(w http.ResponseWriter, r *http.Request) {
	//suiteType := mux.Vars(r)["type"]
	//
	//auth := r.Header.Get("Authorization")



	w.Write([]byte("Create New Suite"))
}

