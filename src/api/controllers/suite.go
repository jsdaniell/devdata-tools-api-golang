package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jsdaniell/devdata-tools-api-golang/api/models/request_models"
	"github.com/jsdaniell/devdata-tools-api-golang/api/repository/suite_repository"
	"github.com/jsdaniell/devdata-tools-api-golang/api/repository/user_repository"
	"github.com/jsdaniell/devdata-tools-api-golang/api/responses"
	"io/ioutil"
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

	if len(suites) == 0 {
		suite_repository.CreateSuite(user.Uid, suiteType, "Default")
	} else {
		responses.JSON(w, http.StatusOK, suites)
		return
	}

	suites, err = suite_repository.GetAllSuites(user.Uid, suiteType)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, suites)
	return
}

func CreateNewSuite(w http.ResponseWriter, r *http.Request) {
	var createSuiteRequestModel request_models.CreateSuiteRequestModel

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = json.Unmarshal(bytes, &createSuiteRequestModel)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	auth := r.Header.Get("Authorization")

	user, err := user_repository.GetUserByUid(auth)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, err = suite_repository.CreateSuite(user.Uid, createSuiteRequestModel.Type, createSuiteRequestModel.Title)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}

	suites, err := suite_repository.GetAllSuites(user.Uid, createSuiteRequestModel.Type)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusCreated, suites)
	return
}
