package suite_repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jsdaniell/devdata-tools-api-golang/api/db"
	"github.com/jsdaniell/devdata-tools-api-golang/api/models/database_models"
	"github.com/jsdaniell/devdata-tools-api-golang/api/utils/rules"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	doc, err := groupCollection.Doc(rules.DocNameByTitle(nameSuite)).Get(context.Background())
	if status.Code(err) == codes.NotFound {
		res, err := groupCollection.Doc(rules.DocNameByTitle(nameSuite)).Set(context.Background(), suiteModel)
		if err != nil {
			fmt.Errorf("error on registre new suite")
		}

		return res, nil
	} else {
		if doc.Exists() {
			return nil, fmt.Errorf(`the %q suite already exists`, nameSuite)
		} else {
			return nil, fmt.Errorf("error on create suite")
		}
	}
}

func DeleteSuite(uid string, typeSuite string, nameSuite string) error {
	client := db.FirestoreClient()
	defer client.Close()

	groupCollection := client.Collection("users/" + uid + "/" + typeSuite)

	doc, err := groupCollection.Doc(rules.DocNameByTitle(nameSuite)).Get(context.Background())
	if err != nil {
		return err
	}

	childrenName, err := rules.GetChildrenNameOfSuite(typeSuite)
	if err != nil {
		return err
	}

	if doc.Exists() {

		childrenCollection := client.Collection("users/" + uid + "/" + typeSuite + "/" + nameSuite + "/" + childrenName)

		docs, err := childrenCollection.Documents(context.Background()).GetAll()
		if err != nil {
			return err
		}

		for _, docChildren := range docs {
			_, err := docChildren.Ref.Delete(context.Background())
			if err != nil {
				return err
			}
		}

		doc.Ref.Delete(context.Background())

		return nil
	} else {
		return fmt.Errorf("the suite %q don't exists on the collection %q", nameSuite, typeSuite)
	}




}
