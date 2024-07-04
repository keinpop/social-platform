package models

import "time"

type UserDB struct {
	Id          uint         `json:"id"`
	Name        string       `json:"name"`
	Fathername  string       `json:"fathername"`
	Surname     string       `json:"surname"`
	Mail        string       `json:"mail"`
	Password    string       `json:"password"`
	AvatarURL   string       `json:"avatar_url"`
	Techonology Technologies `json:"technologies"`
	Workplaces  Workplaces   `json:"workplaces"`
	About       string       `json:"about"`
}

type User struct {
	Id          uint         `json:"id"`
	Name        string       `json:"name"`
	Fathername  string       `json:"fathername"`
	Surname     string       `json:"surname"`
	AvatarURL   string       `json:"avatar_url"`
	Techonology Technologies `json:"technologies"`
	Workplaces  Workplaces   `json:"workplaces"`
	About       string       `json:"about"`

	StudentProfile *Student
	TeacherProfule *Teacher
}

type Student struct {
	EnterDate     time.Time `json:"enter_date"`
	Role          Role      `json:"role"`
	CurrentCourse string    `json:"current_course"`
	Programm      Programm  `json:"programm"`
}

type Teacher struct {
	StudyingYears uint `json:"studying_years"`
}

type StudentDB struct {
	UserID        uint      `json:"user_id"`
	EnterDate     time.Time `json:"enter_date"`
	Role          Role      `json:"role"`
	CurrentCourse string    `json:"current_course"`
	Programm      Programm  `json:"programm"`
}

type TeacherDB struct {
	UserID        uint `json:"user_id"`
	StudyingYears uint `json:"studying_years"`
}

type AdminDB struct {
	UserID uint `json:"user_id"`
}
