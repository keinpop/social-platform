package models

type Company struct {
	Id    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
}

type Companies []Company
