package service

import (
	"department-api/internal/models"
	"department-api/internal/repository"
	"department-api/internal/validation"
	"fmt"
)

type EmployeeService interface {
	CreateEmployee(employee *models.Employee, departmentId uint) error
}

type employeeService struct {
	departmentRepository repository.DepartmentRepository
	employeeRepository   repository.EmployeeRepository
}

func NewEmployeeService(depRepo repository.DepartmentRepository, emplRepo repository.EmployeeRepository) EmployeeService {
	return &employeeService{
		departmentRepository: depRepo,
		employeeRepository:   emplRepo,
	}
}

func (s *employeeService) CreateEmployee(employee *models.Employee, departmentId uint) error {
	targetDepartment, err := s.departmentRepository.GetDepartmentById(departmentId)
	if err != nil {
		return fmt.Errorf("CreateEmployee service error: %w", err)
	}
	if targetDepartment == nil {
		return fmt.Errorf("Can't add employee to not existing department")
	}

	if ok, err := validation.ValidateMaxLength(employee.FullName, "full_name", 200); !ok {
		return err
	}
	if ok, err := validation.ValidateEmpty(employee.FullName, "full_name"); !ok {
		return err
	}

	if ok, err := validation.ValidateMaxLength(employee.Position, "position", 200); !ok {
		return err
	}
	if ok, err := validation.ValidateEmpty(employee.Position, "position"); !ok {
		return err
	}

	err = s.employeeRepository.CreateEmployee(employee, departmentId)

	return err
}
