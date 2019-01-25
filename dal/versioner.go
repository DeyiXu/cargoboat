package dal

import "github.com/cargoboat/cargoboat/model"

// Versioner 版本
type Versioner interface {
	// Add ...
	Add(version *model.Version) (err error)
	// Update ...
	Update(version *model.Version) (err error)
	// GetByID ...
	GetByID(id int64) (version *model.Version)
	// GetAllByAppID ...
	GetAllByAppID(appID int64) (versions []*model.Version)
	// IsExistName ...
	IsExistName(name string, appID int64) (exist bool, err error)
}
