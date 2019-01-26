package model

// Mode 模式
type Mode struct {
	Name  string `json:"name" gorm:"not null"`
	AppID int64  `json:"app_id" gorm:"not null"` // 应用ID
	BaseModel
}