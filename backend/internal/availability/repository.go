package availability

import (
	"time"

	"gorm.io/gorm"
)

// Repository defines the availability repository contract.
type Repository interface {
	Create(slot *Availability) error
	GetByID(id uint) (*Availability, error)
	GetAll() ([]Availability, error)
	GetByDoctorID(doctorID uint) ([]Availability, error)
	GetByScheduleID(scheduleID uint) ([]Availability, error)
	GetByDate(date time.Time) ([]Availability, error)
	GetByDoctorAndDate(doctorID uint, date time.Time) ([]Availability, error)
	Update(slot *Availability) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new availability repository.
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create stores a new availability slot.
func (r *repository) Create(slot *Availability) error {
	return r.db.Create(slot).Error
}

// GetByID retrieves an availability slot by ID.
func (r *repository) GetByID(id uint) (*Availability, error) {
	var slot Availability

	if err := r.db.First(&slot, id).Error; err != nil {
		return nil, err
	}

	return &slot, nil
}

// GetAll retrieves all availability slots.
func (r *repository) GetAll() ([]Availability, error) {
	var slots []Availability

	if err := r.db.
		Order("date ASC").
		Order("start_time ASC").
		Find(&slots).Error; err != nil {
		return nil, err
	}

	return slots, nil
}

// GetByDoctorID retrieves availability slots for a doctor.
func (r *repository) GetByDoctorID(doctorID uint) ([]Availability, error) {
	var slots []Availability

	if err := r.db.
		Where("doctor_id = ?", doctorID).
		Order("date ASC").
		Order("start_time ASC").
		Find(&slots).Error; err != nil {
		return nil, err
	}

	return slots, nil
}

// GetByScheduleID retrieves availability slots for a schedule.
func (r *repository) GetByScheduleID(scheduleID uint) ([]Availability, error) {
	var slots []Availability

	if err := r.db.
		Where("schedule_id = ?", scheduleID).
		Order("date ASC").
		Order("start_time ASC").
		Find(&slots).Error; err != nil {
		return nil, err
	}

	return slots, nil
}

// GetByDate retrieves availability slots for a specific date.
func (r *repository) GetByDate(date time.Time) ([]Availability, error) {
	var slots []Availability

	if err := r.db.
		Where("date = ?", date).
		Order("start_time ASC").
		Find(&slots).Error; err != nil {
		return nil, err
	}

	return slots, nil
}

// GetByDoctorAndDate retrieves availability slots for a doctor on a specific date.
func (r *repository) GetByDoctorAndDate(doctorID uint, date time.Time) ([]Availability, error) {
	var slots []Availability

	if err := r.db.
		Where("doctor_id = ? AND date = ?", doctorID, date).
		Order("start_time ASC").
		Find(&slots).Error; err != nil {
		return nil, err
	}

	return slots, nil
}

// Update updates an availability slot.
func (r *repository) Update(slot *Availability) error {
	return r.db.Save(slot).Error
}

// Delete removes an availability slot.
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Availability{}, id).Error
}
