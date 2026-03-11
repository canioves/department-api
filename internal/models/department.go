package models

import (
	"time"
)

type Department struct {
	ID        uint          `gorm:"primaryKey"`
	CreatedAt time.Time     `gorm:"autoCreateTime"`
	Parent    *Department   `gorm:"foreignKey:ParentID"`
	Children  []*Department `gorm:"foreignKey:ParentID"`
	Employees []*Employee   `gorm:"foreignKey:DepartmentID"`
	Name      string
	ParentID  *uint
}
