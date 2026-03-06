package models

import (
	"time"
)

type Department struct {
	ID        int          `gorm:"primaryKey" json:"id"`
	Name      string       `json:"name"`
	ParentID  *uint        `json:"parent_id"`
	CreatedAt time.Time    `gorm:"autoCreateTime" json:"created_at"`
	Parent    *Department  `gorm:"foreignKey:ParentID" json:"parent"`
	Childern  []Department `gorm:"foreignKey:ParentID" json:"children"`
	Employees []Employee   `gorm:"foreignKey:DepartmentID" json:"employees"`
}
