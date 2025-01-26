package controller

import (
	"SISTEM-TICKETING/reports"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReportController struct {
	service reports.ReportService
}

func NewReportController(service reports.ReportService) *ReportController {
	return &ReportController{service: service}
}

func (rc *ReportController) SummaryReport(c *gin.Context) {
	report, err := rc.service.GetSummaryReport()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, report)
}

func (rc *ReportController) EventReport(c *gin.Context) {
	eventIDStr := c.Param("id")

	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	report, err := rc.service.GetEventReport(uint(eventID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, report)
}
