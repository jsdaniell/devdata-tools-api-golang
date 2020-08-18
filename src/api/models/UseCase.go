package models

type Procedure struct {
	Id      string   `json:"id"`
	Content string   `json:"content"`
	Sublist []string `json:"sublist"`
}

type UseCase struct {
	DocId          string      `json:"docId,omitempty"`
	Title          string      `json:"title"`
	Actor          string      `json:"actor"`
	Id             string      `json:"id"`
	Scenario       string      `json:"scenario"`
	ListProcedures []Procedure `json:"listProcedures"`
	Preconditions  []string    `json:"preconditions"`
	PostCondition  string      `json:"postCondition"`
}
