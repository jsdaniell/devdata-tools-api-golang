package suite_repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jsdaniell/devdata-tools-api-golang/api/db"
	"github.com/jsdaniell/devdata-tools-api-golang/api/models/database_models"
	"github.com/jsdaniell/devdata-tools-api-golang/api/utils/rules"
)

func GetAllSuites(uid string, typeSuite string) ([]database_models.Suite, error) {

	client := db.FirestoreClient()
	defer client.Close()

	groupCollection := client.Collection("users/" + uid + "/" + typeSuite)

	documents, err := groupCollection.Documents(context.Background()).GetAll()
	if err != nil {
		return nil, err
	}

	var suites []database_models.Suite

	for _, doc := range documents {
		var suite database_models.Suite

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

	var suiteModel database_models.Suite

	suiteModel.Title = nameSuite
	suiteModel.SharedWith = make([]database_models.SharedWithModel, 1)
	suiteModel.SharedWith = suiteModel.SharedWith[:len(suiteModel.SharedWith)-1]

	groupCollection := client.Collection("users/" + uid + "/" + typeSuite)

	res, err := groupCollection.Doc(rules.DocNameByTitle(nameSuite)).Set(context.Background(), suiteModel)
	if err != nil {
		fmt.Errorf("error on registre new suite")
	}

	return res, nil

}
