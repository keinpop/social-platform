package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mai-platform/models"
	"net/http"
)

type App struct {
	config *Config
}

func NewApp(cfg *Config) *App {
	return &App{config: cfg}
}

type Techonologies []models.Techonology

// @Summary get list of technologies
// @Schemes
// @Description get list of technologies
// @Accept json
// @Produce json
// @Success 200 {object} Techonologies
// @Router /api/technology/list [get]
func (a *App) GetTechnologies(c *gin.Context) {
	//TODO: DB
	t := []models.Techonology{
		{Id: 1, Title: "C++"},
		{Id: 2, Title: "Golang"},
	}

	c.JSON(http.StatusOK, t)
}

func (a *App) AddTechnology(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[error] Failed to read body: %v", err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	var t models.Techonology
	err = json.Unmarshal(jsonData, &t)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if t.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty title",
		})
		return
	}

	// TODO: check repeat title

	// TODO: DB

	c.JSON(http.StatusOK, t)
}
