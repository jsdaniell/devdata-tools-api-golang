package models

type UseCase struct {
	Uid         string `json:"uid"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	PhotoUrl    string `json:"photoURL"`
	ApiKey      string `json:"apiKey"`
}