package role

import (
	"encoding/json"
	"gorm.io/gorm"
	"io"
	"log"
	"mai-platform/internal/middleware"
	"net/http"

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
// @Description post new role in db
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
// @Description get all roles in db
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
// @Description delete role in db
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

	if r.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty title",
		})
		return
	}

	c.JSON(http.StatusOK, r)
}
