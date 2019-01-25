package dal

import (
	"github.com/cargoboat/cargoboat/model"
)

// Configer 配置
type Configer interface {
	// Add ...
	Add(conf *model.Config) (err error)
	// Update ...
	Update(conf *model.Config) (err error)
	// GetByID ...
	GetByID(id int64) (conf *model.Config)
	// IsExistName ...
	IsExistName(name, mode string, appID int64) (exist bool, err error)
	// GetByLastVersion ...
	GetByLastVersion(name, mode string, appID int64) (conf *model.Config)
}
