package programm

import (
	"encoding/json"
	"io"
	"log"

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

// @Summary get all programmes in db
// @Schemes
// @Description get all programmes in db
// @Accept json
// @Produce json
// @Success 200 {object} Programmes
// @Router /programm/list [get]
func GetProgrammes(c *gin.Context) {
	p := []Programm{
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
