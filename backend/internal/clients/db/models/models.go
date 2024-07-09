package models

import (
	"time"
)

type Company struct {
	Id    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"unique;not null"`
}

type Programm struct {
	Id       uint64 `json:"id" gorm:"primaryKey"`
	Title    string `json:"title" gorm:"unique;not null"`
	Duration uint64 `json:"duration" gorm:"not null"`
}

type Role struct {
	Id    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"unique;not null"`
}

type Techonology struct {
	Id    uint64 `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"unique;not null"`
}

type User struct {
	Id           uint          `json:"id" gorm:"primaryKey"`
	Mail         string        `json:"mail" gorm:"unique;not null"`
	Name         string        `json:"name"`
	Fathername   string        `json:"fathername"`
	Surname      string        `json:"surname"`
	AvatarURL    string        `json:"avatar_url"`
	Technologies []Techonology `json:"technologies" gorm:"many2many:user_technologies;"`
	Companies    []Company     `json:"companies" gorm:"many2many:user_companies;"`
	About        string        `json:"about"`

	Student *Student
	Teacher *Teacher
	Admin   *Admin
}

type UserCompanies struct {
	UserID    uint `json:"user_id" gorm:"primaryKey"`
	CompanyID uint `json:"company_id" gorm:"primaryKey"`

	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date" gorm:"not null"`
	EndDate     time.Time `json:"end_date"`
}

type Student struct {
	UserID uint `json:"user_id" gorm:"primaryKey"`
	User   User

	EnterDate     *time.Time `json:"enter_date"`
	Role          *Role      `json:"role"`
	RoleId        *uint      `json:"role_id"`
	CurrentCourse *uint      `json:"current_course"`
	Programm      *Programm  `json:"programm"`
	ProgrammId    *uint      `json:"programm_id"`
}

type Teacher struct {
	UserID uint `json:"user_id" gorm:"primaryKey"`
	User   User

	StudyingYears *uint `json:"studying_years"`
}

type Admin struct {
	UserID uint `json:"user_id" gorm:"primaryKey"`
	User   User
}

type UserHash struct {
	Login        string `gorm:"primaryKey"`
	PasswordHash string
}
