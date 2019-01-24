package model

// ApplicationTableName ...
const ApplicationTableName = "applications"

// Application 应用
type Application struct {
	Name      string `json:"name"`
	AppSecret string `json:"secret" gorm:"column:secret;not null;size:32"`
	BaseModel
}
