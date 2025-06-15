package routes

import (
	"intern/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}

// RegisterPatient handles patient registration
func RegisterPatient(service models.PatientInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pd models.PatientData

		if err := c.ShouldBindJSON(&pd); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid patient data",
			})
			return
		}

		if err := service.RegisterPatient(&pd); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to register patient or invalid Doctor id",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message":    "Patient registered successfully",
			"patient_id": pd.ID,
		})
	}
}

// GetAllPatients returns all patients
func GetAllPatients(service models.PatientInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		patients, err := service.GetAllPatients()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to fetch patients",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data":    patients,
			"count":   len(patients),
			"message": "Patients retrieved successfully",
		})
	}
}

// GetPatientByID returns a specific patient by ID
func GetPatientByID(service models.PatientInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		pid := c.Query("pid")
		if pid == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "patient ID is required",
			})
			return
		}

		id, err := strconv.Atoi(pid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "invalid patient ID",
				"details": err.Error(),
			})
			return
		}

		patient, err := service.GetPatientByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "patient not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    patient,
			"message": "Patient retrieved successfully",
		})
	}
}

// UpdatePatient handles patient updates
func UpdatePatient(service models.PatientInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pd models.PatientData

		if err := c.ShouldBindJSON(&pd); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "invalid patient data",
				"details": err.Error(),
			})
			return
		}

		if err := service.UpdatePatient(&pd); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "failed to update patient",
				"details": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":    "Patient updated successfully",
			"patient_id": pd.ID,
		})
	}
}

// DeletePatient handles patient deletion
func DeletePatient(service models.PatientInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		pid := c.Query("pid")
		if pid == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "patient ID is required",
			})
			return
		}

		id, err := strconv.Atoi(pid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "invalid patient ID",
				"details": err.Error(),
			})
			return
		}

		if err := service.DeletePatient(uint(id)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "failed to delete patient",
				"details": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Patient deleted successfully",
		})
	}
}
