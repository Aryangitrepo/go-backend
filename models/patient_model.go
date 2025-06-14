package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type PatientData struct {
	gorm.Model
	UserID    int    `json:"userid"`
	Pname     string `gorm:"not null" json:"pname"`
	Page      int    `json:"page"`
	Problem   string `gorm:"not null" json:"problem"`
	Condition string `gorm:"not null" json:"condition"`
	User      User   `gorm:"foreignKey:UserID;references:ID" json:"-"`
}
type PatientInterface interface {
	RegisterPatient(*PatientData) error
	GetAllPatients() ([]PatientData, error)
	GetPatientByID(uint) (PatientData, error)
	UpdatePatient(*PatientData) error
	DeletePatient(uint) error
}

// Patient operations
func (s *PatientData) RegisterPatient(patient *PatientData) error {
	if err := db.Create(patient).Error; err != nil {
		return fmt.Errorf("registering patient: %w", err)
	}
	return nil
}

// returns all patients
func (s *PatientData) GetAllPatients() ([]PatientData, error) {
	var patients []PatientData
	if err := db.Preload("User").Find(&patients).Error; err != nil {
		return nil, fmt.Errorf("fetching patients: %w", err)
	}
	return patients, nil
}

// returns patients using ID
func (s *PatientData) GetPatientByID(pid uint) (PatientData, error) {
	var patient PatientData
	err := db.Preload("User").Where("id = ?", pid).First(&patient).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return PatientData{}, fmt.Errorf("patient with ID %d not found", pid)
		}
		return PatientData{}, fmt.Errorf("fetching patient: %w", err)
	}
	return patient, nil
}

// returns error if didnt update
func (s *PatientData) UpdatePatient(patient *PatientData) error {
	result := db.Where("ID=?", patient.ID).Updates(patient)
	if result.Error != nil {
		return fmt.Errorf("updating patient: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no patient found with ID %d", patient.ID)
	}
	return nil
}

// returns error if didnt delete
func (s *PatientData) DeletePatient(id uint) error {
	result := db.Where("id = ?", id).Delete(&PatientData{})
	if result.Error != nil {
		return fmt.Errorf("deleting patient: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no patient found with ID %d", id)
	}
	return nil
}
