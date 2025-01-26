package repository

import (
	"SISTEM-TICKETING/entity"

	"gorm.io/gorm"
)

type TicketRepository interface {
	FindAll() ([]entity.Ticket, error)
	FindByID(id uint) (entity.Ticket, error)
	Create(ticket *entity.Ticket) error
	Update(ticket *entity.Ticket) error
}

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{db: db}
}

func (r *ticketRepository) FindAll() ([]entity.Ticket, error) {
	var tickets []entity.Ticket
	err := r.db.Preload("Event").Find(&tickets).Error
	return tickets, err
}

func (r *ticketRepository) FindByID(id uint) (entity.Ticket, error) {
	var ticket entity.Ticket
	err := r.db.Preload("Event").First(&ticket, id).Error
	return ticket, err
}

func (r *ticketRepository) Create(ticket *entity.Ticket) error {
	return r.db.Create(ticket).Error
}

func (r *ticketRepository) Update(ticket *entity.Ticket) error {
	return r.db.Save(ticket).Error
}
