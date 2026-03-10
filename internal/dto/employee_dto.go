package dto

import (
	"department-api/internal/models"
	"time"
)

type EmployeeRequest struct {
	FullName string  `json:"full_name"`
	Position string  `json:"position"`
	HiredAt  *string `json:"hired_at,omitempty"`
}

type EmployeeResponse struct {
	ID           uint       `json:"id"`
	DepartmentID uint       `json:"department_id"`
	FullName     string     `json:"full_name"`
	Position     string     `json:"position"`
	HiredAt      *time.Time `json:"hired_at,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
}

func ToEmployeeResponse(employee *models.Employee) *EmployeeResponse {
	if employee == nil {
		return nil
	}

	return &EmployeeResponse{
		ID:           employee.ID,
		DepartmentID: employee.DepartmentID,
		FullName:     employee.FullName,
		Position:     employee.Position,
		HiredAt:      employee.HiredAt,
		CreatedAt:    employee.CreatedAt,
	}
}

func (er *EmployeeRequest) ParseHiredAt() (*time.Time, error) {
	if er.HiredAt == nil {
		return nil, nil
	}

	parsedTime, err := time.Parse("2006-01-02", *er.HiredAt)
	if err != nil {
		return nil, err
	}
	return &parsedTime, nil
}
