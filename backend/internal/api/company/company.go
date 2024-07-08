package company

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"log"
	"mai-platform/internal/clients/db/models"
	"mai-platform/internal/middleware"
	"net/http"
)

type Company struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
}

type Companies []Company

// @Summary post new company in db
// @Schemes
// @Tags Company-API
// @Description post new company in db
// @Accept json
// @Produce json
// @Success 200 {object} Company
// @Router /company [post]
func AddCompany(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[error] Failed to read body: %v", err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	var comp Company
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

	a := middleware.GetApp(c)
	res, err := a.DB.AddCompany(comp.Title)
	switch {
	case err == nil:
		c.JSON(http.StatusCreated, Company(*res))
	case errors.Is(err, gorm.ErrCheckConstraintViolated) || errors.Is(err, gorm.ErrDuplicatedKey):
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record already exists",
		})
	default:
		log.Printf("Failed to create company: %v", err)
		c.JSON(http.StatusInternalServerError, "")
	}
}

// @Summary get all companies in db
// @Schemes
// @Tags Company-API
// @Description get all companies in db
// @Accept json
// @Produce json
// @Success 200 {object} Companies
// @Router /company/list [get]
func GetCompanies(c *gin.Context) {
	a := middleware.GetApp(c)
	res, err := a.DB.GetCompanies()
	if err != nil {
		log.Printf("Failed to get companies: %v", err)
		c.JSON(http.StatusInternalServerError, "")
	}

	var ret []Company
	for i := range res {
		ret = append(ret, Company(res[i]))
	}

	c.JSON(http.StatusOK, ret)
}

//Param?????

// @Summary delete company in db
// @Schemes
// @Tags Company-API
// @Description delete company in db
// @Accept json
// @Produce json
// @Success 200 {object} error
// @Router /company [delete]
func DeleteCompany(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[error] Failed to read body: %v", err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	var comp Company
	err = json.Unmarshal(jsonData, &comp)
	if err != nil {
		log.Printf("[error] Failed to unmarshal JSON: %v", err)
		c.JSON(http.StatusBadRequest, "")
		return
	}

	a := middleware.GetApp(c)

	err = a.DB.DeleteCompany(models.Company(comp))
	if err != nil {
		log.Printf("Failed to delete company: %v", err)
		c.JSON(http.StatusInternalServerError, "")
	}

	a := middleware.GetApp(c)
	err = a.DB.DeleteCompanyByID(comp.Id)
	if err != nil {
		log.Printf("[error] Failed to delete company: %v", err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	log.Printf("[info] Company %s deleted successfully", comp.Title)
	c.JSON(http.StatusOK, "")
}
