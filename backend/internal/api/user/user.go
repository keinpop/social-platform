package user

import (
	"errors"
	"log"
	"mai-platform/internal/api/company"
	"mai-platform/internal/api/programm"
	"mai-platform/internal/api/role"
	"mai-platform/internal/api/technology"
	"mai-platform/internal/clients/db/models"
	"mai-platform/internal/middleware"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	Mail        string                  `json:"mail"`
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
	EnterDate     *time.Time         `json:"enter_date"`
	Role          *role.Role         `json:"role"`
	CurrentCourse *uint              `json:"current_course"`
	Programm      *programm.Programm `json:"programm"`
}

type Teacher struct {
	StudyingYears *uint `json:"studying_years"`
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

// @Summary post new user in db
// @Schemes
// @Tags User-API
// @Description Usage example: 'curl -X POST -v -H "Content-Type: application/json" -d '{"mail":"test@mail.com", "password":"password123", "is_student":true}' http://localhost:8080/api/user'
// @Accept json
// @Produce json
// @Success 200 {object} User
// @Router /user [post]
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

	a := middleware.GetApp(c)

	err = a.Auth.Register(ur.Mail, ur.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ivalid data",
		})

		return
	}

	res, err := a.DB.AddUser(ur.Mail, ur.IsStudent)
	switch {
	case err == nil:
		c.JSON(http.StatusCreated, convertUserToJson(res, nil))
	case errors.Is(err, gorm.ErrCheckConstraintViolated) || errors.Is(err, gorm.ErrDuplicatedKey):
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record already exists",
		})
	default:
		log.Printf("Failed to create User: %v", err)
		c.JSON(http.StatusInternalServerError, "")
	}
}

func convertUserToJson(u *models.User, uc []models.UserCompanies) *User {
	var tech technology.Technologies
	for i := range u.Technologies {
		tech = append(tech, technology.Techonology(u.Technologies[i]))
	}

	// возможно перемешается
	var wp Workplaces
	for i := range uc {
		wp = append(wp, Workplace{
			Company:     company.Company(u.Companies[i]),
			Description: uc[i].Description,
			StartDate:   uc[i].StartDate,
			EndDate:     uc[i].EndDate,
		})
	}

	res := &User{
		Id:          u.Id,
		Mail:        u.Mail,
		Name:        u.Name,
		Fathername:  u.Fathername,
		Surname:     u.Surname,
		AvatarURL:   u.AvatarURL,
		Techonology: tech,
		Workplaces:  wp,
		About:       u.About,
	}

	if u.Student != nil {
		res.StudentProfile = &Student{
			EnterDate:     u.Student.EnterDate,
			Role:          (*role.Role)(u.Student.Role),
			CurrentCourse: u.Student.CurrentCourse,
			Programm:      (*programm.Programm)(u.Student.Programm),
		}
	}

	if u.Teacher != nil {
		res.TeacherProfile = &Teacher{
			StudyingYears: u.Teacher.StudyingYears,
		}
	}

	if u.Admin != nil {
		res.AdminProfile = &Admin{}
	}

	return res
}

// API-регстрации : (login password flag) -> регистрация в api авторизации +
// добавления юзера в таблицу user + создание профиля учителя/студента
//
// API-редактирования : (id) -> редактирование фио, о себе, роль
// API-добавление технологий : (idUser idTechnology)
// API-добавление компании :   (idUser idCompany)
// API-добавление роль :       (idUser idRole)
// API-добавление курса :      (idUser idCourse)

type GetUserParams struct {
	Id uint `json:"id"`
}

func GetUserData(c *gin.Context) {
	value := c.Param("id")
	id, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id",
		})
		return
	}

	a := middleware.GetApp(c)
	res, us, err := a.DB.GetUser(uint(id))
	switch {
	case err == nil:
		c.JSON(http.StatusOK, convertUserToJson(res, us))
	case errors.Is(err, gorm.ErrCheckConstraintViolated) || errors.Is(err, gorm.ErrDuplicatedKey):
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record already exists",
		})
	default:
		log.Printf("Failed to get User: %v", err)
		c.JSON(http.StatusInternalServerError, "")
	}
}
