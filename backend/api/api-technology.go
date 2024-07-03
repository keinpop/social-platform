package api

import (
	"encoding/json"
	"io"
	"log"
	"mai-platform/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	c.JSON(http.StatusOK, t)
}

func (a *App) GetTechnologies(c *gin.Context) {
	t := []models.Techonology{
		{Id: 1, Title: "C++"},
		{Id: 2, Title: "Golang"},
		{Id: 3, Title: "Python"},
	}

	c.JSON(http.StatusOK, t)
}

func (a *App) DeleteTechnology(c *gin.Context) {
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

	c.JSON(http.StatusOK, t)
}
