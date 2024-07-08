package programm

import (
	"encoding/json"
	"io"
	"log"
	"mai-platform/internal/clients/db/models"
	"mai-platform/internal/middleware"

	"gorm.io/gorm"

	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Programm struct {
	Id       uint64 `json:"id"`
	Title    string `json:"title"`
	Duration uint64 `json:"duration"`
}

type Programmes []Programm

// @Summary post new programm in db
// @Schemes
// @Tags Programm-API
// @Description Usage example: 'curl -X POST -v -H "Content-Type: application/json" -d '{"title":"ПМИ"}' http://localhost:8080/api/programm'
// @Accept json
// @Produce json
// @Success 200 {object} Programm
// @Router /programm [post]
func AddProgramm(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[error] Failed to read body: %v", err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	var prog Programm
	err = json.Unmarshal(jsonData, &prog)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if prog.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty title",
		})
		return
	}

	a := middleware.GetApp(c)
	res, err := a.DB.AddProgramm(prog.Title, prog.Duration)
	switch {
	case err == nil:
		c.JSON(http.StatusCreated, Programm(*res))
	case errors.Is(err, gorm.ErrCheckConstraintViolated) || errors.Is(err, gorm.ErrDuplicatedKey):
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record already exists",
		})
	default:
		log.Printf("Failed to create Programm: %v", err)
		c.JSON(http.StatusInternalServerError, "")
	}
}

// @Summary get all programmes in db
// @Schemes
// @Tags Programm-API
// @Description Usage example: 'curl http://localhost:8080/api/programm/list'
// @Accept json
// @Produce json
// @Success 200 {object} Programmes
// @Router /programm/list [get]
func GetProgrammes(c *gin.Context) {
	a := middleware.GetApp(c)
	res, err := a.DB.GetProgrammes()
	if err != nil {
		log.Printf("Failed to get companies: %v", err)
		c.JSON(http.StatusInternalServerError, "")
	}

	var ret []Programm
	for i := range res {
		ret = append(ret, Programm(res[i]))
	}

	c.JSON(http.StatusOK, ret)
}

// @Summary delete programm in db
// @Schemes
// @Tags Programm-API
// @Description Usage example: 'curl -X DELETE -v -H "Content-Type: application/json" -d '{"title":"ФИИТ"}' http://localhost:8080/api/programm/'
// @Accept json
// @Produce json
// @Success 200 {object} Programm
// @Router /programm [delete]
func DeleteProgramm(c *gin.Context) {
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

	a := middleware.GetApp(c)

	err = a.DB.DeleteProgramm(models.Programm(p))
	if err != nil {
		log.Printf("Failed to delete programm: %v", err)
		c.JSON(http.StatusInternalServerError, "")
	}

	c.JSON(http.StatusOK, p)
}
