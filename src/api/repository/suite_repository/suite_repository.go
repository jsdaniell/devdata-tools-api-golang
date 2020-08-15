package suite_repository

import (
	"context"
	"encoding/json"
	"github.com/jsdaniell/devdata-tools-api-golang/api/db"
	"github.com/jsdaniell/devdata-tools-api-golang/api/models"
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
