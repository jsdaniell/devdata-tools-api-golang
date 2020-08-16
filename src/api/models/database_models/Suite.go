package database_models

type SharedWithModel struct {
	User string `json:"user"`
	Permission string `json:"permission"`
}

type Suite struct {
	Title string `json:"title"`
	SharedWith []SharedWithModel `json:"sharedWith"`
}
