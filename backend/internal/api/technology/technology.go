package technology

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Techonology struct {
	Id    uint64 `json:"id" yaml:"id"`
	Title string `json:"title" yaml:"title"`
}

type Technologies []Techonology

// @Summary post new technology in db
// @Schemes
// @Description post new technology in db
// @Accept json
// @Produce json
// @Success 200 {object} Techonology
// @Router /technology [post]
func AddTechnology(c *gin.Context) {
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

	c.JSON(http.StatusOK, t)
}

// @Summary get all technologies in db
// @Schemes
// @Description get all technologies in db
// @Accept json
// @Produce json
// @Success 200 {object} Technologies
// @Router /technology/list [get]
func GetTechnologies(c *gin.Context) {
	t := []Techonology{
		{Id: 1, Title: "C++"},
		{Id: 2, Title: "Golang"},
		{Id: 3, Title: "Python"},
	}

	c.JSON(http.StatusOK, t)
}

// @Summary delete technology in db
// @Schemes
// @Description delete technology in db
// @Accept json
// @Produce json
// @Success 200 {object} Techonology
// @Router /technology [delete]
func DeleteTechnology(c *gin.Context) {
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

	c.JSON(http.StatusOK, t)
}
