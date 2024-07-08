package technology

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

type Techonology struct {
	Id    uint64 `json:"id"`
	Title string `json:"title"`
}

type Technologies []Techonology

// @Summary post new technology in db
// @Schemes
// @Tags Techonology-API
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

	var tech Techonology
	err = json.Unmarshal(jsonData, &tech)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if tech.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty title",
		})
		return
	}

	a := middleware.GetApp(c)
	res, err := a.DB.AddTechonology(tech.Title)
	switch {
	case err == nil:
		c.JSON(http.StatusCreated, Techonology(*res))
	case errors.Is(err, gorm.ErrCheckConstraintViolated) || errors.Is(err, gorm.ErrDuplicatedKey):
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record already exists",
		})
	default:
		log.Printf("Failed to create Technology: %v", err)
		c.JSON(http.StatusInternalServerError, "")
	}
}

// @Summary get all technologies in db
// @Schemes
// @Tags Techonology-API
// @Description get all technologies in db
// @Accept json
// @Produce json
// @Success 200 {object} Technologies
// @Router /technology/list [get]
func GetTechnologies(c *gin.Context) {
	a := middleware.GetApp(c)
	res, err := a.DB.GetTechonologies()
	if err != nil {
		log.Printf("Failed to get companies: %v", err)
		c.JSON(http.StatusInternalServerError, "")
	}

	var ret []Techonology
	for i := range res {
		ret = append(ret, Techonology(res[i]))
	}

	c.JSON(http.StatusOK, ret)
}

// @Summary delete technology in db
// @Schemes
// @Tags Techonology-API
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

	a := middleware.GetApp(c)

	err = a.DB.DeleteTechnology(models.Techonology(t))
	if err != nil {
		log.Printf("Failed to delete technology: %v", err)
		c.JSON(http.StatusInternalServerError, "")
	}

	c.JSON(http.StatusOK, t)
}
