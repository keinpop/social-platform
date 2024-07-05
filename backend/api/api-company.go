package api

import (
	"encoding/json"
	"io"
	"log"
	"mai-platform/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary post new company in db
// @Schemes
// @Description post new company in db
// @Accept json
// @Produce json
// @Success 200 {object} models.Company
// @Router /company [post]
func (a *App) AddCompany(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[error] Failed to read body: %v", err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	var comp models.Company
	err = json.Unmarshal(jsonData, &comp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if comp.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty title",
		})
		return
	}

	if result := a.DB.Create(&comp); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error during creation new company",
		})
		return
	}

	c.JSON(http.StatusCreated, comp)
}

// @Summary change company name in db
// @Schemes
// @Description change company name in db
// @Accept json
// @Produce json
// @Success 200 {object} models.Company
// @Router /company [put]
func (a *App) ChangeCompanyName(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[error] Failed to read body: %v", err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	var comp models.Company
	err = json.Unmarshal(jsonData, &comp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if comp.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty title",
		})
		return
	}

	// TODO: change name in db

	c.JSON(http.StatusOK, comp)
}

// @Summary get all companies in db
// @Schemes
// @Description get all companies in db
// @Accept json
// @Produce json
// @Success 200 {object} models.Companies
// @Router /company/list [get]
func (a *App) GetCompanies(c *gin.Context) {
	comp := models.Companies{
		{Id: 1, Title: "Яндекс"},
		{Id: 2, Title: "ВК"},
		{Id: 3, Title: "Тинькофф"},
	}

	c.JSON(http.StatusOK, comp)
}

// @Summary delete company in db
// @Schemes
// @Description delete company in db
// @Accept json
// @Produce json
// @Success 200 {object} models.Company
// @Router /company [delete]
func (a *App) DeleteCompany(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[error] Failed to read body: %v", err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	var comp models.Company
	err = json.Unmarshal(jsonData, &comp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if comp.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty title",
		})
		return
	}

	c.JSON(http.StatusOK, comp)
}
