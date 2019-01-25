package model

// Config 配置
type Config struct {
	Name    string  `json:"name" gorm:"not null;unique_index"`
	Value   string  `json:"value" gorm:"not null"`
	Mode    string  `json:"mode" gorm:"not null"`    // 模式
	ModeID  int64   `json:"mode_id" gorm:"not null"` // 模式ID
	Version float64 `json:"version" gorm:"not null"` // 版本
	AppID   int64   `json:"app_id" gorm:"not null"`  // 应用ID
	BaseModel
}
