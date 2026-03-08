package models

import (
	"time"
)

type Department struct {
	ID        uint          `gorm:"primaryKey" json:"id"`
	Name      string        `json:"name"`
	ParentID  *uint         `json:"parent_id,omitempty"`
	CreatedAt time.Time     `gorm:"autoCreateTime" json:"created_at"`
	Parent    *Department   `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Children  []*Department `gorm:"foreignKey:ParentID" json:"children,omitempty"`
	Employees []*Employee   `gorm:"foreignKey:DepartmentID" json:"employees,omitempty"`
}
