package api

import (
	"encoding/json"
	"io"
	"log"
	"mai-platform/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary post new programm in db
// @Schemes
// @Description post new programm in db
// @Accept json
// @Produce json
// @Success 200 {object} models.Programm
// @Router /api/programm [post]
func (a *App) AddProgramm(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[error] Failed to read body: %v", err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	var p models.Programm
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

	c.JSON(http.StatusOK, p)
}

// @Summary get all programmes in db
// @Schemes
// @Description get all programmes in db
// @Accept json
// @Produce json
// @Success 200 {object} models.Programmes
// @Router /api/programm/list [get]
func (a *App) GetProgrammes(c *gin.Context) {
	p := []models.Programm{
		{Id: 1, Title: "ПМИ", Duration: 4},
		{Id: 2, Title: "ФИИТ", Duration: 4},
		{Id: 3, Title: "ПМ", Duration: 4},
	}

	c.JSON(http.StatusOK, p)
}

// @Summary delete programm in db
// @Schemes
// @Description delete programm in db
// @Accept json
// @Produce json
// @Success 200 {object} models.Programm
// @Router /api/programm [delete]
func (a *App) DeleteProgramm(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[error] Failed to read body: %v", err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	var p models.Programm
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

	c.JSON(http.StatusOK, p)
}
