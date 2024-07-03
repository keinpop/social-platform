package models

type Company struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
}

type Companies []Company
