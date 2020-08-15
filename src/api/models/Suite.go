package models

type SharedWithModel struct {
	User string `json:"user"`
	Permission string `json:"permission"`
}

type Suite struct {
	Title string `json:"title"`
	SharedWithSlice []SharedWithModel `json:"sharedWith"`
}
