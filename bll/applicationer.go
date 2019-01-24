package bll

import "github.com/cargoboat/cargoboat/model"

// Applicationer ...
type Applicationer interface {
	// GetAll ...
	GetAll() (apps []*model.Application)
	// GetOneByID 获取一个App
	GetOneByID(id int64) (app *model.Application)
	// Add 添加应用
	Add(name string) (appID int64, err error)
	// QueryPage 获取翻页数据
	QueryPage(pageNum, pageSize int, appName string) (data []*ApplicationPagedItem, total int64, err error)
}
