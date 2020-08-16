package suite_repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jsdaniell/devdata-tools-api-golang/api/db"
	"github.com/jsdaniell/devdata-tools-api-golang/api/models"
	"github.com/jsdaniell/devdata-tools-api-golang/api/utils/rules"
)

func GetAllSuites(uid string, typeSuite string) ([]models.Suite, error) {

	client := db.FirestoreClient()
	defer client.Close()

	groupCollection := client.Collection("users/" + uid + "/" + typeSuite)

	documents, err := groupCollection.Documents(context.Background()).GetAll()
	if err != nil {
		return nil, err
	}

	var suites []models.Suite

	for _, doc := range documents {
		var suite models.Suite

		jsonString, _ := json.Marshal(doc.Data())

		err := json.Unmarshal(jsonString, &suite)
		if err != nil {
			return nil, err
		}

		suites = append(suites, suite)
	}

	return suites, nil
}

func CreateSuite(uid string, typeSuite string, nameSuite string) (*firestore.WriteResult, error) {

	client := db.FirestoreClient()
	defer client.Close()

	var suiteModel models.Suite

	suiteModel.Title = nameSuite
	suiteModel.SharedWithSlice = make([]models.SharedWithModel, 0)

	groupCollection := client.Collection("users/" + uid + "/" + typeSuite)

	res, err := groupCollection.Doc(rules.DocNameByTitle(nameSuite)).Set(context.Background(), suiteModel)
	if err != nil {
		fmt.Errorf("error on registre new suite")
	}

	return res, nil

}
