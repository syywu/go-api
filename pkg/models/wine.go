package models

type Wine struct {
	// gorm.Model
	Id          int    `json:"id" gorm:"primaryKey"`
	Producer    string `json:"producer"`
	Name        string `json:"name"`
	Region      string `json:"region"`
	Rating      int    `json:"rating"`
	Description string `json:"description"`
}
