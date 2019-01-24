package model

import (
	"time"

	"github.com/cargoboat/cargoboat/module/store"
)

// BaseModel is base entity
type BaseModel struct {
	ID        int64      `json:"id" gorm:"primary_key;unique_index"`
	CreatedAt time.Time  `json:"created_at" gorm:"not null;type:timestamp"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"not null;type:timestamp"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index" gorm:"type:timestamp"`
}

// AutoMigrate 自动迁移
func AutoMigrate() {
	store.DataBase.Master().AutoMigrate(
		&Application{},
		&Config{},
	)
}
