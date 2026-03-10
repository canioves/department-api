package dto

import (
	"department-api/internal/models"
	"time"
)

type CreateDepartmentRequest struct {
	Name     string `json:"name" binding:"required"`
	ParentID *uint  `json:"parent_id"`
}

type UpdateDepartmentRequest struct {
	Name     *string `json:"name"`
	ParentId *uint   `json:"parent_id"`
}

type DepartmentResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	ParentID  *uint     `json:"parent_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

func ToDepartmentResponse(department *models.Department) *DepartmentResponse {
	if department == nil {
		return nil
	}

	return &DepartmentResponse{
		ID:        department.ID,
		Name:      department.Name,
		ParentID:  department.ParentID,
		CreatedAt: department.CreatedAt,
	}
}
