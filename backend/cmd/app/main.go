package main

import (
	"log"
	"mai-platform/api"
	"mai-platform/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1

const CONFIG_PATH = "./etc/config.yaml"

func main() {
	c, err := api.NewConfig(CONFIG_PATH)
	if err != nil {
		log.Fatalf("Failed to get config: %v", err)
	}

	a := api.NewApp(c)

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	r.POST("/api/progrmm", a.AddProgramm)
	r.GET("/api/programm/list", a.GetProgrammes)
	r.DELETE("/api/programm", a.DeleteProgramm)

	r.POST("/api/role", a.AddRole)
	r.GET("/api/role/list", a.GetRoles)
	r.DELETE("/api/role", a.DeleteRole)

	r.POST("/api/technology", a.AddTechnology)
	r.GET("/api/technology/list", a.GetTechnologies)
	r.DELETE("/api/technology", a.DeleteTechnology)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
