package api

import (
	"encoding/json"
	"io"
	"log"
	"mai-platform/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) GetUserData(c *gin.Context) {
	// TODO: use db
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[error] Failed to read body: %v", err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	var r models.User
	err = json.Unmarshal(jsonData, &r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	// TODO: понять, что за роль юзера (stud teach admin)

}
