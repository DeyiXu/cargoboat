package dal

import "github.com/cargoboat/cargoboat/model"

// Modeer 模式
type Modeer interface {
	// Add ...
	Add(mode *model.Mode) (err error)
	// Update ...
	Update(mode *model.Mode) (err error)
	// GetByID ...
	GetByID(id int64) (mode *model.Mode)
	// GetAllByAppID ...
	GetAllByAppID(appID int64) (modes []*model.Mode)
	// IsExistName ...
	IsExistName(name string, appID int64) (exist bool, err error)
}
