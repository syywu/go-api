package models

type Wine struct {
	Id          int    `json:"id"`
	Producer    string `json:"producer"`
	Name        string `json:"name"`
	Region      string `json:"region"`
	Rating      int    `json:"rating"`
	Description string `json:"description"`
}
