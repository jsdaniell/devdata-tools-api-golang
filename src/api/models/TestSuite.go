package models

type SharedWithModel struct {
	user string
	permission string
}

type TestSuite struct {
    title string
 	sharedWith []SharedWithModel
}