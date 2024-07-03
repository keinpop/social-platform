package models

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

type Techonology struct {
	Id    uint64 `json:"id" yaml:"id"`
	Title string `json:"title" yaml:"title"`
}

func AddNewTechonology(c *gin.Context) {
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

func GetAllTechonologies(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[error] Failed to read body: %v", err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	var t []Techonology
	err = json.Unmarshal(jsonData, &t)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if len(t) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No technologies found",
		})
		return
	}

	c.JSON(http.StatusOK, t)
}
