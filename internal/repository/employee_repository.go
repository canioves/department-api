package repository

import (
	"department-api/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	CreateEmployee(employee *models.Employee, departmentId uint) error
	GetEmployeeById(id uint) (*models.Employee, error)
	GetEmployeesByDepartment(departmentId uint) ([]*models.Employee, error)
	DeleteEmployees(departmentId uint) error
	MoveEmployeeToDepartment(id uint, departmentId uint) error
}

type employeeRepository struct {
	database *gorm.DB
}

func NewEmployeeRepository(database *gorm.DB) EmployeeRepository {
	return &employeeRepository{database: database}
}

func (r *employeeRepository) GetEmployeeById(id uint) (*models.Employee, error) {
	var employee *models.Employee
	result := r.database.First(&employee, id)

	if err := result.Error; err != nil {
		return nil, fmt.Errorf("GetEmployee error: %w", err)
	}
	return employee, nil
}

func (r *employeeRepository) CreateEmployee(employee *models.Employee, departmentId uint) error {
	employee.DepartmentID = departmentId
	result := r.database.Create(employee)
	if err := result.Error; err != nil {
		return fmt.Errorf("CreateEmployee error: %w", err)
	}
	return nil
}

func (r *employeeRepository) GetEmployeesByDepartment(departmentId uint) ([]*models.Employee, error) {
	var employees []*models.Employee
	result := r.database.Where("department_id = ?", departmentId).Find(&employees)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("GetEmployeesByDepartment error: %w", err)
	}
	return employees, nil
}

func (r *employeeRepository) DeleteEmployees(departmentId uint) error {
	var employee *models.Employee
	result := r.database.Model(&employee).Where("department_id = ?", departmentId).Delete(&employee)
	if err := result.Error; err != nil {
		return fmt.Errorf("DeleteEmployees error: %w", err)
	}
	return nil
}

func (r *employeeRepository) MoveEmployeeToDepartment(id uint, departmentId uint) error {
	var employee *models.Employee
	return r.database.Model(&employee).Where("id = ?", id).Update("department_id", departmentId).Error
}
