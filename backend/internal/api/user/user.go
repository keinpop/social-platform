package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mai-platform/internal/api/company"
	"mai-platform/internal/api/programm"
	"mai-platform/internal/api/role"
	"mai-platform/internal/api/technology"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Workplace struct {
	Company     company.Company `json:"company_name"`
	Description string          `json:"description"`
	StartDate   time.Time       `json:"start_date"`
	EndDate     time.Time       `json:"end_date"`
}

type Workplaces []Workplace

type CreateUserParams struct {
	isStudent bool
}

type User struct {
	Id          uint                    `json:"id"`
	Name        string                  `json:"name"`
	Fathername  string                  `json:"fathername"`
	Surname     string                  `json:"surname"`
	AvatarURL   string                  `json:"avatar_url"`
	Techonology technology.Technologies `json:"technologies"`
	Workplaces  Workplaces              `json:"workplaces"`
	About       string                  `json:"about"`

	StudentProfile *Student
	TeacherProfile *Teacher
	AdminProfile   *Admin
}

type Student struct {
	EnterDate     time.Time         `json:"enter_date"`
	Role          role.Role         `json:"role"`
	CurrentCourse string            `json:"current_course"`
	Programm      programm.Programm `json:"programm"`
}

type Teacher struct {
	StudyingYears uint `json:"studying_years"`
}

type Admin struct {
}

type UserRequest struct {
	Mail      string `json:"mail"`
	Password  string `json:"password"`
	IsStudent bool   `json:"is_student"`
}

// type LoginPassword struct {
// 	Login    string `form:"login" json:"login" binding:"required"`
// 	Password string `form:"password" json:"password" binding:"required"`
// }

func AddUser(c *gin.Context) {
	// идет в авторизацию
	// пытается создать
	// если ок - добавляем данные по юзеру в бд
	// иначе возвращаем
	var ur UserRequest

	err := c.ShouldBind(&ur)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	postBody, _ := json.Marshal(map[string]string{
		"login":    ur.Mail,
		"password": ur.Password,
	})

	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("http://backend-auth:8090/register", "application/json", responseBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	c.JSON(http.StatusCreated, "OK")
}

// API-регстрации : (login password flag) -> регистрация в api авторизации +
// добавления юзера в таблицу user + создание профиля учителя/студента
//
// API-редактирования : (id) -> редактирование фио, о себе, роль
// API-добавление технологий : (idUser idTechnology)
// API-добавление компании :   (idUser idCompany)
// API-добавление роль :       (idUser idRole)
// API-добавление курса :      (idUser idCourse)

// Доделать
func GetUserData(c *gin.Context) {
	// TODO: use db
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[error] Failed to read body: %v", err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	var r User
	err = json.Unmarshal(jsonData, &r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	// TODO: понять, что за роль юзера (stud teach admin)

}
