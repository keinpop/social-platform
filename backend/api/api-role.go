package api

import (
	"encoding/json"
	"io"
	"log"
	"mai-platform/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary post new role in db
// @Schemes
// @Description post new role in db
// @Accept json
// @Produce json
// @Success 200 {object} models.Role
// @Router /role [post]
func (a *App) AddRole(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[error] Failed to read body: %v", err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	var r models.Role
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

// @Summary get all roles in db
// @Schemes
// @Description get all roles in db
// @Accept json
// @Produce json
// @Success 200 {object} models.Roles
// @Router /role/list [get]
func (a *App) GetRoles(c *gin.Context) {
	r := []models.Role{
		{Id: 1, Title: "Teamlead"},
		{Id: 2, Title: "Frontend-разработчик"},
		{Id: 3, Title: "Backend-разработчик"},
		{Id: 4, Title: "ML-инженер"},
	}

	c.JSON(http.StatusOK, r)
}

// @Summary delete role in db
// @Schemes
// @Description delete role in db
// @Accept json
// @Produce json
// @Success 200 {object} models.Role
// @Router /role [delete]
func (a *App) DeleteRole(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[error] Failed to read body: %v", err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	var r models.Role
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
