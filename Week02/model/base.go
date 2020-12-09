package model

import "time"

type BaseModel struct {
	ID        int64      `gorm:"column:id;primary_key" json:"id,string"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"-"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index" json:"-"`
}
