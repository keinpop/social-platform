package role

import (
	"encoding/json"
	"io"
	"log"
	"mai-platform/internal/clients/db/models"
	"mai-platform/internal/middleware"
	"net/http"

	"gorm.io/gorm"

	"errors"

	"github.com/gin-gonic/gin"
)

type Role struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
}

type Roles []Role

// @Summary post new role in db
// @Schemes
// @Tags Role-API
// @Description Usage example: 'curl -X POST -v -H "Content-Type: application/json" -d '{"title":"Backend-разработчик"}' http://localhost:8080/api/role'
// @Accept json
// @Produce json
// @Success 200 {object} Role
// @Router /role [post]
func AddRole(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[error] Failed to read body: %v", err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	var role Role
	err = json.Unmarshal(jsonData, &role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if role.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty title",
		})
		return
	}

	a := middleware.GetApp(c)
	res, err := a.DB.AddRole(role.Title)
	switch {
	case err == nil:
		c.JSON(http.StatusCreated, Role(*res))
	case errors.Is(err, gorm.ErrCheckConstraintViolated) || errors.Is(err, gorm.ErrDuplicatedKey):
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record already exists",
		})
	default:
		log.Printf("Failed to create Role: %v", err)
		c.JSON(http.StatusInternalServerError, "")
	}
}

// @Summary get all roles in db
// @Schemes
// @Tags Role-API
// @Description Usage example: 'curl http://localhost:8080/api/role/list'
// @Accept json
// @Produce json
// @Success 200 {object} Roles
// @Router /role/list [get]
func GetRoles(c *gin.Context) {
	a := middleware.GetApp(c)
	res, err := a.DB.GetRoles()
	if err != nil {
		log.Printf("Failed to get companies: %v", err)
		c.JSON(http.StatusInternalServerError, "")
	}

	var ret []Role
	for i := range res {
		ret = append(ret, Role(res[i]))
	}

	c.JSON(http.StatusOK, ret)
}

// @Summary delete role in db
// @Schemes
// @Tags Role-API
// @Description Usage example: 'curl -X DELETE -v -H "Content-Type: application/json" -d '{"title":"Frontend-разработчик"}' http://localhost:8080/api/role/'
// @Accept json
// @Produce json
// @Success 200 {object} Role
// @Router /role [delete]
func DeleteRole(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[error] Failed to read body: %v", err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	var r Role
	err = json.Unmarshal(jsonData, &r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	a := middleware.GetApp(c)

	err = a.DB.DeleteRole(models.Role(r))
	if err != nil {
		log.Printf("Failed to delete programm: %v", err)
		c.JSON(http.StatusInternalServerError, "")
	}

	c.JSON(http.StatusOK, r)
}
