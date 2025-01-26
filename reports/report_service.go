package reports

import (
	"SISTEM-TICKETING/entity"

	"gorm.io/gorm"
)

type ReportService interface {
	GetSummaryReport() (map[string]interface{}, error)
	GetEventReport(eventID uint) (map[string]interface{}, error)
}

type reportService struct {
	db *gorm.DB
}

func NewReportService(db *gorm.DB) ReportService {
	return &reportService{db: db}
}

func (s *reportService) GetSummaryReport() (map[string]interface{}, error) {
	var totalTickets int64
	var totalRevenue float64

	err := s.db.Model(&entity.Ticket{}).Where("status = ?", "sold").Count(&totalTickets).Error
	if err != nil {
		return nil, err
	}

	err = s.db.Model(&entity.Ticket{}).Where("status = ?", "sold").Select("SUM(price)").Scan(&totalRevenue).Error
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"total_tickets_sold": totalTickets,
		"total_revenue":      totalRevenue,
	}, nil
}

func (s *reportService) GetEventReport(eventID uint) (map[string]interface{}, error) {
	var totalTickets int64
	var totalRevenue float64

	err := s.db.Model(&entity.Ticket{}).Where("event_id = ? AND status = ?", eventID, "sold").Count(&totalTickets).Error
	if err != nil {
		return nil, err
	}

	err = s.db.Model(&entity.Ticket{}).Where("event_id = ? AND status = ?", eventID, "sold").Select("SUM(price)").Scan(&totalRevenue).Error
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"event_id":           eventID,
		"total_tickets_sold": totalTickets,
		"total_revenue":      totalRevenue,
	}, nil
}
