package programm

import (
	"encoding/json"
	"gorm.io/gorm"
	"io"
	"log"
	"mai-platform/internal/middleware"

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
// @Description post new programm in db
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
// @Description get all programmes in db
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
// @Description delete programm in db
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
