package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	config *Config
}

func NewApp(cfg *Config) *App {
	return &App{config: cfg}
}

type Techonology struct {
	Id    uint64 `json:"id"`
	Title string `json:"title"`
}

type Techonologies []Techonology

// @Summary get list of technologies
// @Schemes
// @Description get list of technologies
// @Accept json
// @Produce json
// @Success 200 {object} Techonologies
// @Router /api/technology/list [get]
func (a *App) GetTechnologies(c *gin.Context) {
	//TODO: DB
	t := []Techonology{
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

	var t Techonology
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
