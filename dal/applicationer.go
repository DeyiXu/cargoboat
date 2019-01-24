package dal

import "github.com/cargoboat/cargoboat/model"

// Applicationer ...
type Applicationer interface {
	// GetPaged 获取翻页数据
	QueryPage(pageNum, pageSize int, appName string) (data []*model.Application, total int64, err error)
	// GetAll ...
	GetAll() (apps []*model.Application)
	// Add ...
	Add(app *model.Application) (err error)
	// Update ...
	Update(app *model.Application) (err error)
	// GetByID return Application by ID
	GetByID(id int64) (app *model.Application)
	// IsExistName ...
	IsExistName(name string) (b bool, err error)
}
