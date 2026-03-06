package models

import (
	"time"
)

type Department struct {
	ID        int          `gorm:"primaryKey" json:"id"`
	Name      string       `json:"name"`
	ParentID  *uint        `json:"parent_id"`
	CreatedAt time.Time    `json:"created_at"`
	Parent    *Department  `gorm:"foreignKey:ParentID" json:"parent"`
	Childern  []Department `gorm:"foreignKey:ParentID" json:"child"`
}
