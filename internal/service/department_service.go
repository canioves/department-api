package service

import (
	"department-api/internal/models"
	"department-api/internal/repository"
	"department-api/internal/validation"
	"fmt"
	"strings"
)

type DepartmentService interface {
	CreateDepartment(department *models.Department) (*models.Department, error)
	GetDepartment(id uint, depth int, includeEmployees bool) (*models.Department, error)
}

type departmentService struct {
	departmentRepository repository.DepartmentRepository
	employeeRepository   repository.EmployeeRepository
}

func NewDepartmentService(depRepo repository.DepartmentRepository, emplRepo repository.EmployeeRepository) DepartmentService {
	return &departmentService{
		departmentRepository: depRepo,
		employeeRepository:   emplRepo,
	}
}

func (s *departmentService) CreateDepartment(department *models.Department) (*models.Department, error) {
	if department.ID == *department.ParentID {
		return nil, fmt.Errorf("Can't create department with id = parent_id")
	}
	if ok, err := validation.ValidateMaxLength(department.Name, "name", 200); !ok {
		return nil, err
	}
	if ok, err := validation.ValidateEmpty(department.Name, "name"); !ok {
		return nil, err
	}

	trimmedName := strings.Trim(department.Name, " ")
	department.Name = trimmedName

	siblings, _ := s.departmentRepository.GetSiblingsDepartments(department.ParentID)
	for _, sibling := range siblings {
		if sibling.Name == department.Name {
			return nil, fmt.Errorf("Duplicate name on same level: %s", department.Name)
		}
	}

	result, err := s.departmentRepository.CreateDepartment(department)
	return result, err
}

func (s *departmentService) GetDepartment(id uint, depth int, includeEmployees bool) (*models.Department, error) {
	root, _ := s.departmentRepository.GetDepartmentById(id)
	root.Children = s.buildTree(id, depth, includeEmployees, 0)

	if includeEmployees {
		root.Employees = s.collectEmployees(root)
	}
	return root, nil
}

func (s *departmentService) buildTree(parentID uint, depth int, includeEmployees bool, currentDepth int) []*models.Department {
	if depth > 0 && currentDepth >= depth {
		return nil
	}

	children, _ := s.departmentRepository.GetChildrenDepartments(&parentID)

	var result []*models.Department
	for _, child := range children {
		if includeEmployees {
			child.Employees, _ = s.employeeRepository.GetEmployeesByDepartment(child.ID)
		}
		child.Children = s.buildTree(child.ID, depth, includeEmployees, currentDepth+1)
		result = append(result, child)
	}
	return result
}

func (s *departmentService) collectEmployees(department *models.Department) []*models.Employee {
	var allEmployees []*models.Employee
	allEmployees = append(allEmployees, department.Employees...)

	for _, child := range department.Children {
		childEmployees := s.collectEmployees(child)
		allEmployees = append(allEmployees, childEmployees...)
	}
	return allEmployees
}
