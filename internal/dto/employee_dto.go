package dto

import (
	"department-api/internal/models"
	"time"
)

type EmployeeRequest struct {
	FullName string     `json:"full_name"`
	Position string     `json:"position"`
	HiredAt  *time.Time `json:"hired_at,omitempty"`
}

type EmployeeResponse struct {
	ID           uint    `json:"id"`
	DepartmentID uint    `json:"department_id"`
	FullName     string  `json:"full_name"`
	Position     string  `json:"position"`
	HiredAt      *string `json:"hired_at,omitempty"`
	CreatedAt    string  `json:"created_at"`
}

func ToEmployeeResponce(employee *models.Employee) *EmployeeResponse {
	if employee == nil {
		return nil
	}

	var hiredAt *string
	if employee.HiredAt != nil {
		formatted := employee.HiredAt.Format(time.RFC3339)
		hiredAt = &formatted
	}

	return &EmployeeResponse{
		ID:           employee.ID,
		DepartmentID: employee.DepartmentID,
		FullName:     employee.FullName,
		Position:     employee.Position,
		HiredAt:      hiredAt,
		CreatedAt:    employee.CreatedAt.Format(time.RFC3339),
	}
}
