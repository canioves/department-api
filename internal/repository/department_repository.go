package repository

import (
	"department-api/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type DepartmentRepository interface {
	CreateDepartment(department *models.Department) (*models.Department, error)
	GetAllDepartments() ([]*models.Department, error)
	GetChildrenDepartments(parentID *uint) ([]*models.Department, error)
	GetDepartmentById(id *uint) (*models.Department, error)
	GetSiblingsDepartments()
}

type departmentRepository struct {
	database *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentRepository{database: db}
}

func (r *departmentRepository) CreateDepartment(department *models.Department) (*models.Department, error) {
	result := r.database.Create(department)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("CreateDepartment error: %w", err)
	}
	return department, nil
}

func (r *departmentRepository) GetAllDepartments() ([]*models.Department, error) {
	var departments []*models.Department
	result := r.database.Find(&departments)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("GetAllDepartments error: %w", err)
	}
	return departments, nil
}

func (r *departmentRepository) GetChildrenDepartments(parentID *uint) ([]*models.Department, error) {
	var children []*models.Department
	result := r.database.Where("parent_id = ?", parentID).Find(&children)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("GetChildren error: %w", err)
	}
	return children, nil
}

func (r *departmentRepository) GetDepartmentById(id *uint) (*models.Department, error) {
	var department *models.Department
	result := r.database.Where("id = ?", id).First(&department)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("GetDepartmentById error: %w", err)
	}
	return department, nil

}
