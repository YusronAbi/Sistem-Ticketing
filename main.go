package main

import (
	"SISTEM-TICKETING/config"
	"SISTEM-TICKETING/controller"
	"SISTEM-TICKETING/reports"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()
	config.MigrateDB()

	r := gin.Default()

	reportService := reports.NewReportService(config.DB)
	reportController := controller.NewReportController(reportService)

	r.GET("/reports/summary", reportController.SummaryReport)
	r.GET("/reports/event/:id", reportController.EventReport)

	r.Run(":8080")
}
