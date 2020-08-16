package database_models

type User struct {
	Uid         string `json:"uid"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	PhotoUrl    string `json:"photoURL"`
	ApiKey      string `json:"apiKey"`
}
