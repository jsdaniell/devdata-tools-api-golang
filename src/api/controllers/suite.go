package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jsdaniell/devdata-tools-api-golang/api/repository/suite_repository"
	"github.com/jsdaniell/devdata-tools-api-golang/api/repository/user_repository"
	"github.com/jsdaniell/devdata-tools-api-golang/api/responses"
	"github.com/jsdaniell/devdata-tools-api-golang/api/utils/rules"
	"io/ioutil"
	"net/http"
)

func GetAllSuitesOfAType(w http.ResponseWriter, r *http.Request) {

	suiteType := mux.Vars(r)["type"]

	auth := r.Header.Get("Authorization")

	err := rules.ValidateExistentSuites(suiteType)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user, err := user_repository.GetUserByUid(auth)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	suites, err := suite_repository.GetAllSuites(user.Uid, suiteType)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if len(suites) == 0 {
		_, err := suite_repository.CreateSuite(user.Uid, suiteType, "Default")
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
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

	type CreateSuiteRequestModel struct {
		Title string `json:"title"`
		Type  string `json:"type"`
	}

	var createSuiteRequestModel CreateSuiteRequestModel

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
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	err = rules.ValidateExistentSuites(createSuiteRequestModel.Type)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, err = suite_repository.CreateSuite(user.Uid, createSuiteRequestModel.Type, createSuiteRequestModel.Title)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	suites, err := suite_repository.GetAllSuites(user.Uid, createSuiteRequestModel.Type)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusCreated, suites)
	return
}

func DeleteSuite(w http.ResponseWriter, r *http.Request) {

	suiteType := mux.Vars(r)["type"]
	suiteId := mux.Vars(r)["id"]

	auth := r.Header.Get("Authorization")

	user, err := user_repository.GetUserByUid(auth)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	err = suite_repository.DeleteSuite(user.Uid, suiteType, suiteId)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
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

func AddNewItemOnSuite(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")

	suiteType := mux.Vars(r)["type"]
	suiteId := mux.Vars(r)["id"]

	user, errUser := user_repository.GetUserByUid(auth)
	if errUser != nil {
		responses.ERROR(w, http.StatusUnauthorized, errUser)
		return
	}

	var entity, err = rules.GetInterfaceOfSuite(suiteType)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = json.Unmarshal(bytes, &entity)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = suite_repository.AddNewItemOnSuite(user.Uid, suiteType, suiteId, entity)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	allItems, err := suite_repository.GetItemsFromSuite(user.Uid, suiteType, suiteId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, allItems)
}

func GetAllItemsFromSuite(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")

	suiteType := mux.Vars(r)["type"]
	suiteId := mux.Vars(r)["id"]

	lastDoc := r.URL.Query().Get("lastDoc")
	navigate := r.URL.Query().Get("navigate")



	user, errUser := user_repository.GetUserByUid(auth)
	if errUser != nil {
		responses.ERROR(w, http.StatusUnauthorized, errUser)
		return
	}

	if navigate != "" && lastDoc != "" {

		switch navigate {
		case "next":
			allItems, err := suite_repository.GetItemsFromSuiteNext(user.Uid, suiteType, suiteId, lastDoc)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			 responses.JSON(w, http.StatusOK, allItems)
			return
		case "previous":
			allItems, err := suite_repository.GetItemsFromSuitePrevious(user.Uid, suiteType, suiteId, lastDoc)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			responses.JSON(w, http.StatusOK, allItems)
			return
		default:
			responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("missing valid navigation string on query parameter"))
			return
		}
	} else {
		allItems, err := suite_repository.GetItemsFromSuite(user.Uid, suiteType, suiteId)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		responses.JSON(w, http.StatusOK, allItems)
		return
	}


}

func EditItemFromSuite(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")

	suiteType := mux.Vars(r)["type"]
	suiteId := mux.Vars(r)["id"]
	itemId := mux.Vars(r)["itemId"]

	user, errUser := user_repository.GetUserByUid(auth)
	if errUser != nil {
		responses.ERROR(w, http.StatusUnauthorized, errUser)
		return
	}

	var entity, err = rules.GetInterfaceOfSuite(suiteType)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = json.Unmarshal(bytes, &entity)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = suite_repository.EditItemFromSuite(user.Uid, suiteType, suiteId, itemId, entity)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}

	allItems, err := suite_repository.GetItemsFromSuite(user.Uid, suiteType, suiteId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, allItems)
}

func DeleteItemFromSuite(w http.ResponseWriter, r *http.Request){
	auth := r.Header.Get("Authorization")

	suiteType := mux.Vars(r)["type"]
	suiteId := mux.Vars(r)["id"]
	itemId := mux.Vars(r)["itemId"]

	user, errUser := user_repository.GetUserByUid(auth)
	if errUser != nil {
		responses.ERROR(w, http.StatusUnauthorized, errUser)
		return
	}

	err := suite_repository.DeleteItemFromSuite(user.Uid, suiteType, suiteId, itemId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}

	allItems, err := suite_repository.GetItemsFromSuite(user.Uid, suiteType, suiteId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, allItems)
}
