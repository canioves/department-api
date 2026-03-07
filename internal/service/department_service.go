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
	nameLength := len(department.Name)
	if nameLength < 1 {
		return nil, fmt.Errorf("Name too short")
	}
	if nameLength > 200 {
		return nil, fmt.Errorf("Name too long")
	}
	if department.Name == "" {
		return nil, fmt.Errorf("Name cannot be empty")
	}

	siblings := s.repository.
	trimmedName := strings.Trim(department.Name, " ")
	department.Name = trimmedName

	result, err := s.repository.CreateDepartment(department)
	return result, err
}
