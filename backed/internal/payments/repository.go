package payments

import "gorm.io/gorm"

// Repository defines the payment repository contract.
type Repository interface {
	Create(payment *Payment) error
	GetByID(id uint) (*Payment, error)
	GetAll() ([]Payment, error)

	GetByAppointmentID(appointmentID uint) ([]Payment, error)
	GetByPatientID(patientID uint) ([]Payment, error)
	GetByDoctorID(doctorID uint) ([]Payment, error)
	GetByHospitalID(hospitalID uint) ([]Payment, error)

	Update(payment *Payment) error
	Delete(id uint) error

	GetSummary() (*PaymentSummaryResponse, error)
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new payment repository.
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create stores a payment.
func (r *repository) Create(payment *Payment) error {
	return r.db.Create(payment).Error
}

// GetByID retrieves a payment by ID.
func (r *repository) GetByID(id uint) (*Payment, error) {
	var payment Payment

	if err := r.db.First(&payment, id).Error; err != nil {
		return nil, err
	}

	return &payment, nil
}

// GetAll retrieves all payments.
func (r *repository) GetAll() ([]Payment, error) {
	var payments []Payment

	if err := r.db.Order("created_at DESC").Find(&payments).Error; err != nil {
		return nil, err
	}

	return payments, nil
}

// GetByAppointmentID retrieves payments for an appointment.
func (r *repository) GetByAppointmentID(appointmentID uint) ([]Payment, error) {
	var payments []Payment

	if err := r.db.
		Where("appointment_id = ?", appointmentID).
		Order("created_at DESC").
		Find(&payments).Error; err != nil {
		return nil, err
	}

	return payments, nil
}

// GetByPatientID retrieves payments made by a patient.
func (r *repository) GetByPatientID(patientID uint) ([]Payment, error) {
	var payments []Payment

	if err := r.db.
		Where("patient_id = ?", patientID).
		Order("created_at DESC").
		Find(&payments).Error; err != nil {
		return nil, err
	}

	return payments, nil
}

// GetByDoctorID retrieves payments received by a doctor.
func (r *repository) GetByDoctorID(doctorID uint) ([]Payment, error) {
	var payments []Payment

	if err := r.db.
		Where("doctor_id = ?", doctorID).
		Order("created_at DESC").
		Find(&payments).Error; err != nil {
		return nil, err
	}

	return payments, nil
}

// GetByHospitalID retrieves payments belonging to a hospital.
func (r *repository) GetByHospitalID(hospitalID uint) ([]Payment, error) {
	var payments []Payment

	if err := r.db.
		Where("hospital_id = ?", hospitalID).
		Order("created_at DESC").
		Find(&payments).Error; err != nil {
		return nil, err
	}

	return payments, nil
}

// Update updates a payment.
func (r *repository) Update(payment *Payment) error {
	return r.db.Save(payment).Error
}

// Delete removes a payment.
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Payment{}, id).Error
}

// GetSummary returns payment statistics.
func (r *repository) GetSummary() (*PaymentSummaryResponse, error) {
	summary := &PaymentSummaryResponse{}

	// Total payments
	if err := r.db.Model(&Payment{}).
		Count(&summary.TotalPayments).Error; err != nil {
		return nil, err
	}

	// Total revenue
	r.db.Model(&Payment{}).
		Where("status = ?", StatusPaid).
		Select("COALESCE(SUM(amount),0)").
		Scan(&summary.TotalRevenue)

	// Pending amount
	r.db.Model(&Payment{}).
		Where("status = ?", StatusPending).
		Select("COALESCE(SUM(amount),0)").
		Scan(&summary.PendingAmount)

	// Paid amount
	r.db.Model(&Payment{}).
		Where("status = ?", StatusPaid).
		Select("COALESCE(SUM(amount),0)").
		Scan(&summary.PaidAmount)

	// Refunded amount
	r.db.Model(&Payment{}).
		Where("status = ?", StatusRefunded).
		Select("COALESCE(SUM(amount),0)").
		Scan(&summary.RefundedAmount)

	return summary, nil
}
