package models

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

type UserJSON struct {
	Id          uint         `json:"id"`
	Name        string       `json:"name"`
	Fathername  string       `json:"fathername"`
	Surname     string       `json:"surname"`
	AvatarURL   string       `json:"avatar_url"`
	Techonology Technologies `json:"technologies"`
	Workplaces  Workplaces   `json:"workplaces"`
	About       string       `json:"about"`
}

type Student struct {
	EnterDate     string   `json:"enter_date"`
	Role          Role     `json:"role"`
	CurrentCourse string   `json:"current_course"`
	Programm      Programm `json:"programm"`
}

type Teacher struct {
	StudyingYears uint `json:"studying_years"`
}
