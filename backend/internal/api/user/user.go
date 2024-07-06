package api

import (
	"encoding/json"
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
	TeacherProfule *Teacher
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
