package service

import (
	"department-api/internal/models"
	"department-api/internal/repository"
	"fmt"
)

type EmployeeService interface {
}

type employeeService struct {
	departmentRepository repository.DepartmentRepository
	employeeRepository   repository.EmployeeRepository
}

func NewEmployeeRepository(depRepo repository.DepartmentRepository, emplRepo repository.EmployeeRepository) EmployeeService {
	return &employeeService{
		departmentRepository: depRepo,
		employeeRepository:   emplRepo,
	}
}

func (s *employeeService) CreateEmployee(employee *models.Employee) (*models.Employee, error) {
	targetDepartment, err := s.departmentRepository.GetDepartmentById(&employee.DepartmentID)
	if err != nil {
		return nil, fmt.Errorf("CreateEmployee service error: %w", err)
	}
	if targetDepartment == nil {
		return nil, fmt.Errorf("Can't add employee to not existing department")
	}

	return employee, nil
}
