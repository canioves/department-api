package service

import (
	"department-api/internal/models"
	"department-api/internal/repository"
	"department-api/internal/validation"
	"errors"
	"fmt"
	"strings"
)

var ErrCycle = errors.New("a cycle detected")

type DepartmentService interface {
	CreateDepartment(department *models.Department) error
	GetDepartment(id uint, depth int, includeEmployees bool) (*models.Department, error)
	UpdateDepartment(id uint, department *models.Department) (*models.Department, error)
	DeleteDepartment(id uint, mode string, reassignId uint) error
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

func (s *departmentService) CreateDepartment(department *models.Department) error {
	if department.ParentID != nil {
		if *department.ParentID == 0 {
			return fmt.Errorf("parent ID cannot be 0")
		}

		parentDepartment, err := s.departmentRepository.GetDepartmentById(*department.ParentID)
		if parentDepartment == nil {
			return fmt.Errorf("parent department is not exist")
		}
		if err != nil {
			return err
		}
	}

	trimmedName := strings.Trim(department.Name, " ")
	department.Name = trimmedName

	if ok, err := validation.ValidateMaxLength(department.Name, "name", 200); !ok {
		return err
	}
	if ok, err := validation.ValidateEmpty(department.Name, "name"); !ok {
		return err
	}

	siblings, _ := s.departmentRepository.GetSiblingsDepartments(department.ParentID)
	for _, sibling := range siblings {
		if sibling.Name == department.Name {
			return fmt.Errorf("Duplicate name on same level: %s", department.Name)
		}
	}

	err := s.departmentRepository.CreateDepartment(department)
	return err
}

func (s *departmentService) GetDepartment(id uint, depth int, includeEmployees bool) (*models.Department, error) {
	root, err := s.departmentRepository.GetDepartmentById(id)
	if err != nil {
		return nil, err
	}
	root.Children = s.buildTree(id, depth, includeEmployees, 0)

	if includeEmployees {
		root.Employees, err = s.employeeRepository.GetEmployeesByDepartment(root.ID)
		if err != nil {
			return nil, err
		}
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
			retrivedEmployees, err := s.employeeRepository.GetEmployeesByDepartment(child.ID)
			if err != nil {
				return nil
			}
			child.Employees = retrivedEmployees
		}
		child.Children = s.buildTree(child.ID, depth, includeEmployees, currentDepth+1)
		result = append(result, child)
	}
	return result
}

func (s *departmentService) UpdateDepartment(id uint, department *models.Department) (*models.Department, error) {
	existing, err := s.departmentRepository.GetDepartmentById(id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, fmt.Errorf("department not found")
	}

	if department.Name != "" {
		trimmedName := strings.Trim(department.Name, " ")
		if ok, err := validation.ValidateEmpty(trimmedName, "name"); !ok {
			return nil, err
		}
		if ok, err := validation.ValidateMaxLength(trimmedName, "name", 200); !ok {
			return nil, err
		}
		parentID := existing.ParentID
		if department.ParentID != nil {
			parentID = department.ParentID
		}
		siblings, _ := s.departmentRepository.GetSiblingsDepartments(parentID)
		for _, sib := range siblings {
			if sib.ID != existing.ID && sib.Name == trimmedName {
				return nil, fmt.Errorf("duplicate name on same level: %s", trimmedName)
			}
		}
		existing.Name = trimmedName
	}

	if department.ParentID != nil {
		if *department.ParentID == existing.ID {
			return nil, fmt.Errorf("cannot set department as its own parent")
		}
		parentDepartment, err := s.departmentRepository.GetDepartmentById(*department.ParentID)
		if err != nil {
			return nil, err
		}
		if parentDepartment == nil {
			return nil, fmt.Errorf("parent department does not exist")
		}
		cur := parentDepartment
		for cur != nil {
			if cur.ID == existing.ID {
				return nil, ErrCycle
			}
			if cur.ParentID == nil {
				break
			}
			cur, _ = s.departmentRepository.GetDepartmentById(*cur.ParentID)
		}
		existing.ParentID = department.ParentID
	}

	err = s.departmentRepository.UpdateDepartment(existing)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *departmentService) DeleteDepartment(id uint, mode string, reassignId uint) error {
	switch mode {
	case "reassign":
		employees, _ := s.recursiveCollectEmployees(id)
		for _, empl := range employees {
			s.employeeRepository.MoveEmployeeToDepartment(empl.ID, reassignId)
		}
		return s.departmentRepository.DeleteDepartment(id)
	case "cascade":
		s.recursiveDeleteEmployees(id)
		return s.departmentRepository.DeleteDepartment(id)
	default:
		return fmt.Errorf("Mode does not support: %s", mode)
	}
}

func (s *departmentService) recursiveDeleteEmployees(id uint) error {
	children, _ := s.departmentRepository.GetChildrenDepartments(&id)
	for _, child := range children {
		if err := s.employeeRepository.DeleteEmployees(child.ID); err != nil {
			return err
		}
		if err := s.recursiveDeleteEmployees(child.ID); err != nil {
			return err
		}
	}
	return s.employeeRepository.DeleteEmployees(id)
}

func (s *departmentService) recursiveCollectEmployees(id uint) ([]*models.Employee, error) {
	var allEmployees []*models.Employee
	employees, _ := s.employeeRepository.GetEmployeesByDepartment(id)
	children, _ := s.departmentRepository.GetChildrenDepartments(&id)
	allEmployees = append(allEmployees, employees...)

	for _, child := range children {
		childEmployees, _ := s.recursiveCollectEmployees(child.ID)
		allEmployees = append(allEmployees, childEmployees...)
	}
	return allEmployees, nil
}
