package service

import (
	"department-api/internal/models"
	"department-api/internal/repository"
	"fmt"
	"strings"
)

type DepartmentService interface {
	CreateDepartment(department *models.Department) (*models.Department, error)
}

type departmentService struct {
	repository repository.DepartmentRepository
}

func NewDepartmentService(repository repository.DepartmentRepository) DepartmentService {
	return &departmentService{repository: repository}
}

func (s *departmentService) CreateDepartment(department *models.Department) (*models.Department, error) {
	if department.ID == *department.ParentID {
		return nil, fmt.Errorf("Can't create department with id = parent_id")
	}

	nameLength := len(department.Name)
	if nameLength > 200 {
		return nil, fmt.Errorf("Name too long")
	}
	if department.Name == "" {
		return nil, fmt.Errorf("Name cannot be empty")
	}

	trimmedName := strings.Trim(department.Name, " ")
	department.Name = trimmedName

	siblings, _ := s.repository.GetSiblingsDepartments(department.ParentID)
	for _, sibling := range siblings {
		if sibling.Name == department.Name {
			return nil, fmt.Errorf("Duplicate name on same level: %s", department.Name)
		}
	}

	result, err := s.repository.CreateDepartment(department)
	return result, err
}

// func (s *departmentService) GetDepartment(id uint, depth int, includeEmployee bool) ([]*models.Department, error) {

// }
