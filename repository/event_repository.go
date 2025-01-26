package repository

import (
	"SISTEM-TICKETING/entity"

	"gorm.io/gorm"
)

type EventRepository interface {
	FindAll() ([]entity.Event, error)
	FindByID(id uint) (entity.Event, error)
	Create(event *entity.Event) error
	Update(event *entity.Event) error
	Delete(event *entity.Event) error
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db: db}
}

func (r *eventRepository) FindAll() ([]entity.Event, error) {
	var events []entity.Event
	err := r.db.Find(&events).Error
	return events, err
}

func (r *eventRepository) FindByID(id uint) (entity.Event, error) {
	var event entity.Event
	err := r.db.First(&event, id).Error
	return event, err
}

func (r *eventRepository) Create(event *entity.Event) error {
	return r.db.Create(event).Error
}

func (r *eventRepository) Update(event *entity.Event) error {
	return r.db.Save(event).Error
}

func (r *eventRepository) Delete(event *entity.Event) error {
	return r.db.Delete(event).Error
}
