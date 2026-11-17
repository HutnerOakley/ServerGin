package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID          uint           `gorm:"primary_key" json:"id"`
	CreatedTime time.Time      `gorm:"column:created_time;autoCreateTime" json:"created_time"`
	UpdatedTime time.Time      `gorm:"column:updated_time;autoUpdateTime" json:"updated_time"`
	IsDeleted   gorm.DeletedAt `gorm:"column:is_deleted;index" json:"-"` // 逻辑删除
}
