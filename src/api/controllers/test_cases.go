package controllers

import "net/http"

func GetAllTestsSuites(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetAllTestsSuites"))
}

func DeleteTestSuite(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteTestSuite"))
}
func CreateTestSuite(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateTestSuite"))
}
func UpdateTestSuite(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateTestSuite"))
}

func GetTestsFromATestSuite(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetTestsFromATestSuite"))
}

func DeleteOneTestFromTestSuite(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteOneTestFromTestSuite"))
}
func UpdateOneTestFromTestSuite(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateOneTestFromTestSuite"))
}
func CreateOneTestOnTestSuite(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateOneTestSuite"))
}
