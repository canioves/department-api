package repository

import (
	"department-api/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type DepartmentRepository interface {
	CreateDepartment(department *models.Department) error
	GetAllDepartments() ([]*models.Department, error)
	GetChildrenDepartments(parentID *uint) ([]*models.Department, error)
	GetDepartmentById(id uint) (*models.Department, error)
	GetSiblingsDepartments(id *uint) ([]*models.Department, error)
	UpdateDepartment(department *models.Department) error
	DeleteDepartment(id uint) error
}

type departmentRepository struct {
	database *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentRepository{database: db}
}

func (r *departmentRepository) CreateDepartment(department *models.Department) error {
	result := r.database.Create(department)
	if err := result.Error; err != nil {
		return fmt.Errorf("CreateDepartment error: %w", err)
	}
	return nil
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
		return nil, fmt.Errorf("GetChildrenDepartments error: %w", err)
	}
	return children, nil
}

func (r *departmentRepository) GetDepartmentById(id uint) (*models.Department, error) {
	var department *models.Department
	result := r.database.Where("id = ?", id).First(&department)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("GetDepartmentById error: %w", err)
	}
	return department, nil
}

func (r *departmentRepository) GetSiblingsDepartments(id *uint) ([]*models.Department, error) {
	var siblings []*models.Department
	var whereResult *gorm.DB

	if id != nil {
		whereResult = r.database.Where("parent_id = ?", id)
	} else {
		whereResult = r.database.Where("parent_id is null")
	}

	result := whereResult.Find(&siblings)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("GetSiblingsDepartments error: %w", err)
	}
	return siblings, nil
}

func (r *departmentRepository) UpdateDepartment(department *models.Department) error {
	result := r.database.Model(department).Select("name", "parent_id").Updates(department)
	if err := result.Error; err != nil {
		return fmt.Errorf("UpdateDepartment error: %w", err)
	}
	return nil
}

func (r *departmentRepository) DeleteDepartment(id uint) error {
	var department *models.Department
	result := r.database.Delete(&department, id)
	if err := result.Error; err != nil {
		return fmt.Errorf("DeleteDepartment error: %w", err)
	}
	return nil
}
