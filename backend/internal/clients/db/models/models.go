package models

import (
	"time"
)

type Company struct {
	Id    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"unique"`
}

type Programm struct {
	Id       uint64 `json:"id" gorm:"primaryKey"`
	Title    string `json:"title" gorm:"unique"`
	Duration uint64 `json:"duration"`
}

type Role struct {
	Id    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"unique"`
}

type Techonology struct {
	Id    uint64 `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"unique"`
}

type User struct {
	Id            uint          `json:"id" gorm:"primaryKey"`
	Name          string        `json:"name"`
	Fathername    string        `json:"fathername"`
	Surname       string        `json:"surname"`
	Mail          string        `json:"mail" gorm:"unique"`
	Password      string        `json:"password"`
	AvatarURL     string        `json:"avatar_url"`
	Techonologies []Techonology `json:"technologies" gorm:"many2many:user_technologies;"`
	Companies     []Company     `json:"companies" gorm:"many2many:user_companies;"`
	About         string        `json:"about"`
}

type UserCompanies struct {
	UserID    uint `json:"user_id" gorm:"primaryKey"`
	CompanyID uint `json:"company_id" gorm:"primaryKey"`

	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

type Student struct {
	UserID uint `json:"user_id" gorm:"primaryKey"`
	User   User

	EnterDate     time.Time `json:"enter_date"`
	Role          Role      `json:"role" gorm:"many2many:student_roles;"`
	CurrentCourse string    `json:"current_course"`
	Programm      Programm  `json:"programm" gorm:"many2many:student_programms;"`
}

type Teacher struct {
	UserID uint `json:"user_id" gorm:"primaryKey"`
	User   User

	StudyingYears uint `json:"studying_years"`
}

type Admin struct {
	UserID uint `json:"user_id" gorm:"primaryKey"`
	User   User
}
