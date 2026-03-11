package models

import "time"

type Employee struct {
	ID           uint       `gorm:"primaryKey"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	Department   Department `gorm:"foreignKey:DepartmentID"`
	FullName     string
	Position     string
	HiredAt      *time.Time
	DepartmentID uint
}
