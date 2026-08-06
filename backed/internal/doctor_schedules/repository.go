package doctor_schedules

import "gorm.io/gorm"

// Repository defines the doctor schedule repository contract.
type Repository interface {
	Create(schedule *DoctorSchedule) error
	GetByID(id uint) (*DoctorSchedule, error)
	GetAll() ([]DoctorSchedule, error)
	GetByDoctorID(doctorID uint) ([]DoctorSchedule, error)
	GetByDay(day Weekday) ([]DoctorSchedule, error)
	Update(schedule *DoctorSchedule) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new doctor schedule repository.
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create stores a doctor schedule.
func (r *repository) Create(schedule *DoctorSchedule) error {
	return r.db.Create(schedule).Error
}

// GetByID retrieves a doctor schedule by ID.
func (r *repository) GetByID(id uint) (*DoctorSchedule, error) {
	var schedule DoctorSchedule

	if err := r.db.First(&schedule, id).Error; err != nil {
		return nil, err
	}

	return &schedule, nil
}

// GetAll retrieves all doctor schedules.
func (r *repository) GetAll() ([]DoctorSchedule, error) {
	var schedules []DoctorSchedule

	if err := r.db.
		Order("doctor_id ASC").
		Order("day ASC").
		Find(&schedules).Error; err != nil {
		return nil, err
	}

	return schedules, nil
}

// GetByDoctorID retrieves all schedules for a doctor.
func (r *repository) GetByDoctorID(doctorID uint) ([]DoctorSchedule, error) {
	var schedules []DoctorSchedule

	if err := r.db.
		Where("doctor_id = ?", doctorID).
		Order("day ASC").
		Find(&schedules).Error; err != nil {
		return nil, err
	}

	return schedules, nil
}

// GetByDay retrieves schedules for a specific weekday.
func (r *repository) GetByDay(day Weekday) ([]DoctorSchedule, error) {
	var schedules []DoctorSchedule

	if err := r.db.
		Where("day = ?", day).
		Order("doctor_id ASC").
		Find(&schedules).Error; err != nil {
		return nil, err
	}

	return schedules, nil
}

// Update updates a doctor schedule.
func (r *repository) Update(schedule *DoctorSchedule) error {
	return r.db.Save(schedule).Error
}

// Delete removes a doctor schedule.
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&DoctorSchedule{}, id).Error
}
