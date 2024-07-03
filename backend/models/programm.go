package models

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

type Programm struct {
	Id       uint64 `json:"id" yaml:"id"`
	Title    string `json:"title" yaml:"title"`
	Duration uint64 `json:"duration" yaml:"duration"`
}

func AddNewProgramm(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[error] Failed to read body: %v", err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	var p Programm
	err = json.Unmarshal(jsonData, &p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if p.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty title",
		})
		return
	}

	if p.Duration == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty duration",
		})
		return
	}

	// TODO: check repeat title

	// TODO: DB

	c.JSON(http.StatusOK, p)
}

func GetAllProgrammes(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[error] Failed to read body: %v", err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	var p []Programm
	err = json.Unmarshal(jsonData, &p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if len(p) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No technologies found",
		})
		return
	}

	c.JSON(http.StatusOK, p)
}
