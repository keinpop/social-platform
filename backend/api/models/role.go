package models

type Role struct {
	Id    uint   `json:"id" yaml:"id"`
	Title string `json:"title" yaml:"title"`
}

type Roles []Role
