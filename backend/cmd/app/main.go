package main

import (
	"log"
	"mai-platform/docs"
	"mai-platform/internal/api/company"
	"mai-platform/internal/api/programm"
	"mai-platform/internal/api/role"
	"mai-platform/internal/api/technology"
	"mai-platform/internal/app"
	"mai-platform/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api

const CONFIG_PATH = "./config/config.yaml"

func main() {
	c, err := app.NewConfig(CONFIG_PATH)
	if err != nil {
		log.Fatalf("Failed to get config: %v", err)
	}

	a := app.NewApp(c)
	if err := a.Init(); err != nil {
		log.Fatalf("Failed to init app: %v", err)
	}

	r := gin.Default()
	r.Use(middleware.WithApp(a))

	InitMetrics(r)

	docs.SwaggerInfo.BasePath = "/api"

	r.POST("/api/progrmm", programm.AddProgramm)
	r.GET("/api/programm/list", programm.GetProgrammes)
	r.DELETE("/api/programm", programm.DeleteProgramm)

	r.POST("/api/role", role.AddRole)
	r.GET("/api/role/list", role.GetRoles)
	r.DELETE("/api/role", role.DeleteRole)

	r.POST("/api/technology", technology.AddTechnology)
	r.GET("/api/technology/list", technology.GetTechnologies)
	r.DELETE("/api/technology", technology.DeleteTechnology)

	r.POST("/api/company", company.AddCompany)
	r.GET("/api/company/list", company.GetCompanies)
	r.DELETE("/api/company/", company.DeleteCompany)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func InitMetrics(r *gin.Engine) {
	m := ginmetrics.GetMonitor()

	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	// set middleware for gin
	m.Use(r)
}
