package model

// Version 版本
type Version struct {
	Name  string `json:"name" gorm:"not null"`
	AppID int64  `json:"app_id" gorm:"not null"` // 应用ID
	BaseModel
}
