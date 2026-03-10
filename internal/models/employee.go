package models

import "time"

type Employee struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	FullName     string     `json:"full_name"`
	Position     string     `json:"position"`
	HiredAt      *time.Time `json:"hired_at"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	DepartmentID uint       `json:"department_id"`
	Department   Department `gorm:"foreignKey:DepartmentID" json:"department"`
}
