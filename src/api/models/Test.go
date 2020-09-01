package models

type Test struct {
	DocId         string   `json:"docId,omitempty"`
	Title         string   `json:"title"`
	Id            string   `json:"id"`
	Environment   string   `json:"environment"`
	Priority      string   `json:"priority"`
	Name          string   `json:"name"`
	Actor         string   `json:"actor"`
	Preconditions []string `json:"preconditions"`
	Procedures    []string `json:"procedures"`
	PostCondition string   `json:"postCondition"`
}
