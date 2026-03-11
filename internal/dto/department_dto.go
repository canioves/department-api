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

type DepartmentDetailResponse struct {
	ID        uint                        `json:"id"`
	Name      string                      `json:"name"`
	CreatedAt time.Time                   `json:"created_at"`
	Children  []*DepartmentDetailResponse `json:"children,omitempty"`
	Employees []*EmployeeResponse         `json:"employees,omitempty"`
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

func ToDepartmentDetailResponse(department *models.Department) *DepartmentDetailResponse {
	if department == nil {
		return nil
	}

	response := &DepartmentDetailResponse{
		ID:        department.ID,
		Name:      department.Name,
		CreatedAt: department.CreatedAt,
		Children:  make([]*DepartmentDetailResponse, len(department.Children)),
		Employees: make([]*EmployeeResponse, len(department.Employees)),
	}

	for i, empl := range department.Employees {
		response.Employees[i] = ToEmployeeResponse(empl)
	}

	for i, dep := range department.Children {
		response.Children[i] = ToDepartmentDetailResponse(dep)
	}

	return response
}
