package models

type Programm struct {
	Id       uint64 `json:"id" yaml:"id"`
	Title    string `json:"title" yaml:"title"`
	Duration uint64 `json:"duration" yaml:"duration"`
}

type Programmes []Programm
