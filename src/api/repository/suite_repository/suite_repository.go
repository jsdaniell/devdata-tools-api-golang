package suite_repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jsdaniell/devdata-tools-api-golang/api/db"
	"github.com/jsdaniell/devdata-tools-api-golang/api/models"
	"github.com/jsdaniell/devdata-tools-api-golang/api/utils/json_utility"
	"github.com/jsdaniell/devdata-tools-api-golang/api/utils/rules"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const LimitOfDocsPerPage = 7

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

		docWithId := doc.Data()

		docWithId["docId"] = doc.Ref.ID

		jsonString, _ := json.Marshal(docWithId)

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
	suiteModel.SharedWith = make([]models.SharedWithModel, 1)
	suiteModel.SharedWith = suiteModel.SharedWith[:len(suiteModel.SharedWith)-1]

	groupCollection := client.Collection("users/" + uid + "/" + typeSuite)

	// TODO: Change Keys of Suite Model to lowerCase
	

	lowerCaseJson, err := json_utility.StructToLowerCaseJson(suiteModel)
	if err != nil {
		return nil, err
	}


	doc, err := groupCollection.Doc(rules.DocNameByTitle(nameSuite)).Get(context.Background())
	if status.Code(err) == codes.NotFound {
		res, err := groupCollection.Doc(rules.DocNameByTitle(nameSuite)).Set(context.Background(), lowerCaseJson)
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

func DeleteSuite(uid string, typeSuite string, suiteId string) error {
	client := db.FirestoreClient()
	defer client.Close()

	groupCollection := client.Collection("users/" + uid + "/" + typeSuite)

	doc, err := groupCollection.Doc(suiteId).Get(context.Background())
	if err != nil {
		return fmt.Errorf("the suite %q don't exists on the collection %q", suiteId, typeSuite)
	}

	childrenName, err := rules.GetChildrenNameOfSuite(typeSuite)
	if err != nil {
		return err
	}

	if doc.Exists() {

		childrenCollection := client.Collection("users/" + uid + "/" + typeSuite + "/" + suiteId + "/" + childrenName)

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
		return fmt.Errorf("the suite %q don't exists on the collection %q", suiteId, typeSuite)
	}
}

// Suite Items Repository Transactions

func AddNewItemOnSuite(uid string, typeSuite string, idSuite string, item interface{}) error {
	client := db.FirestoreClient()
	defer client.Close()

	groupCollection := client.Collection("users/" + uid + "/" + typeSuite)

	doc, err := groupCollection.Doc(idSuite).Get(context.Background())
	if err != nil {
		return fmt.Errorf("the suite %q don't exists on the collection %q", idSuite, typeSuite)
	}

	childrenName, err := rules.GetChildrenNameOfSuite(typeSuite)
	if err != nil {
		return err
	}

	if doc.Exists() {

		childrenCollection := client.Collection("users/" + uid + "/" + typeSuite + "/" + idSuite + "/" + childrenName)

		var marshaled, err = json.Marshal(item)
		if err != nil {
			return err
		}

		type TitleModel struct {
			Title string `json:"title"`
		}

		titleModel := TitleModel{}

		err = json.Unmarshal(marshaled, &titleModel)
		if err != nil {
			return err
		}

		docExists, _ := childrenCollection.Doc(rules.DocNameByTitle(titleModel.Title)).Get(context.Background())

		if docExists.Exists() {
			return fmt.Errorf("already Exists a document with this title")
		}

		_, err = childrenCollection.Doc(rules.DocNameByTitle(titleModel.Title)).Set(context.Background(), item)
		if err != nil {
			return err
		}

		return nil
	} else {
		return fmt.Errorf("the suite %q don't exists on the collection %q", idSuite, typeSuite)
	}
}

func GetItemsFromSuite(uid string, typeSuite string, idSuite string) ([]interface{}, error) {
	client := db.FirestoreClient()
	defer client.Close()

	groupCollection := client.Collection("users/" + uid + "/" + typeSuite)

	doc, err := groupCollection.Doc(idSuite).Get(context.Background())
	if err != nil {
		return nil, fmt.Errorf("the suite %q don't exists on the collection %q", idSuite, typeSuite)
	}

	childrenName, err := rules.GetChildrenNameOfSuite(typeSuite)
	if err != nil {
		return nil, err
	}

	if doc.Exists() {

		childrenCollection, _ := client.Collection("users/" + uid + "/" + typeSuite + "/" + idSuite + "/" + childrenName).OrderBy("title", firestore.Asc).Limit(LimitOfDocsPerPage).Documents(context.Background()).GetAll()

		var items []interface{}

		for _, doc := range childrenCollection {
			var item interface{}

			docWithId := doc.Data()

			docWithId["docId"] = doc.Ref.ID

			jsonString, _ := json.Marshal(docWithId)

			err := json.Unmarshal(jsonString, &item)
			if err != nil {
				return nil, err
			}

			items = append(items, item)
		}

		return items, nil
	} else {
		return nil, fmt.Errorf("the suite %q don't exists on the collection %q", idSuite, typeSuite)
	}

}

func GetItemsFromSuiteNext(uid string, typeSuite string, idSuite string, lastDoc string) ([]interface{}, error) {
	client := db.FirestoreClient()
	defer client.Close()

	groupCollection := client.Collection("users/" + uid + "/" + typeSuite)

	doc, err := groupCollection.Doc(idSuite).Get(context.Background())
	if err != nil {
		return nil, fmt.Errorf("the suite %q don't exists on the collection %q", idSuite, typeSuite)
	}

	childrenName, err := rules.GetChildrenNameOfSuite(typeSuite)
	if err != nil {
		return nil, err
	}

	if doc.Exists() {

		childrenCollection, err := client.Collection("users/" + uid + "/" + typeSuite + "/" + idSuite + "/" + childrenName).OrderBy("title", firestore.Asc).StartAfter(lastDoc).Limit(LimitOfDocsPerPage).Documents(context.Background()).GetAll()
		if err != nil {
			return nil, err
		}

		var items []interface{}

		for _, doc := range childrenCollection {
			var item interface{}

			docWithId := doc.Data()

			docWithId["docId"] = doc.Ref.ID

			jsonString, _ := json.Marshal(docWithId)

			err := json.Unmarshal(jsonString, &item)
			if err != nil {
				return nil, err
			}

			items = append(items, item)
		}

		return items, nil
	} else {
		return nil, fmt.Errorf("the suite %q don't exists on the collection %q", idSuite, typeSuite)
	}

}

func GetItemsFromSuitePrevious(uid string, typeSuite string, idSuite string, lastDoc string) ([]interface{}, error) {
	client := db.FirestoreClient()
	defer client.Close()

	groupCollection := client.Collection("users/" + uid + "/" + typeSuite)

	doc, err := groupCollection.Doc(idSuite).Get(context.Background())
	if err != nil {
		return nil, fmt.Errorf("the suite %q don't exists on the collection %q", idSuite, typeSuite)
	}

	childrenName, err := rules.GetChildrenNameOfSuite(typeSuite)
	if err != nil {
		return nil, err
	}

	if doc.Exists() {

		childrenCollection, _ := client.Collection("users/" + uid + "/" + typeSuite + "/" + idSuite + "/" + childrenName).OrderBy("title", firestore.Asc).EndBefore(lastDoc).Limit(LimitOfDocsPerPage).Documents(context.Background()).GetAll()

		var items []interface{}

		for _, doc := range childrenCollection {
			var item interface{}

			docWithId := doc.Data()

			docWithId["docId"] = doc.Ref.ID

			jsonString, _ := json.Marshal(docWithId)

			err := json.Unmarshal(jsonString, &item)
			if err != nil {
				return nil, err
			}

			items = append(items, item)
		}

		return items, nil
	} else {
		return nil, fmt.Errorf("the suite %q don't exists on the collection %q", idSuite, typeSuite)
	}

}

func EditItemFromSuite(uid string, typeSuite string, idSuite string, idItem string, jsonToUpdate interface{}) error {
	client := db.FirestoreClient()
	defer client.Close()

	groupCollection := client.Collection("users/" + uid + "/" + typeSuite)

	doc, err := groupCollection.Doc(idSuite).Get(context.Background())
	if err != nil {
		return fmt.Errorf("the suite %q don't exists on the collection %q", idSuite, typeSuite)
	}

	childrenName, err := rules.GetChildrenNameOfSuite(typeSuite)
	if err != nil {
		return err
	}

	if doc.Exists() {

		docFromCollection, _ := client.Collection("users/" + uid + "/" + typeSuite + "/" + idSuite + "/" + childrenName).Doc(idItem).Get(context.Background())

		if docFromCollection.Exists() {

			_, err := docFromCollection.Ref.Set(context.Background(), jsonToUpdate)
			if err != nil {
				return err
			}

		} else {
			return fmt.Errorf(`the document id %q don't exists`, idItem)
		}

		return nil
	} else {
		return fmt.Errorf("the suite %q don't exists on the collection %q", idSuite, typeSuite)
	}
}

func DeleteItemFromSuite(uid string, typeSuite string, idSuite string, idItem string) error {
	client := db.FirestoreClient()
	defer client.Close()

	groupCollection := client.Collection("users/" + uid + "/" + typeSuite)

	doc, err := groupCollection.Doc(idSuite).Get(context.Background())
	if err != nil {
		return fmt.Errorf("the suite %q don't exists on the collection %q", idSuite, typeSuite)
	}

	childrenName, err := rules.GetChildrenNameOfSuite(typeSuite)
	if err != nil {
		return err
	}

	if doc.Exists() {

		_, err :=  client.Collection("users/" + uid + "/" + typeSuite + "/" + idSuite + "/" + childrenName).Doc(idItem).Delete(context.Background())
		if err != nil {
			return err
		}

		return nil
	} else {
		return fmt.Errorf("the suite %q don't exists on the collection %q", idSuite, typeSuite)
	}
}
