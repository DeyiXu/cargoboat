package application

import (
	dalApplication "github.com/cargoboat/cargoboat/dal/application"
	"github.com/cargoboat/cargoboat/model"
	cerrors "github.com/cargoboat/cargoboat/module/errors"
	"github.com/cargoboat/cargoboat/module/store"
)

// GetAll 获取全部App
func GetAll() []*model.Application {
	return dalApplication.GetAll()
}

// GetOneByID 获取一个App
func GetOneByID(id int64) *model.Application {
	return dalApplication.GetByID(id)
}

// Add 添加应用
func Add(name string) (appID int64, err error) {
	b, err := dalApplication.IsExistName(name)
	if err != nil {
		err = cerrors.GetBusinessError(2001)
		return
	}
	if b {
		err = cerrors.GetBusinessError(2002)
		return
	}
	app := &model.Application{
		Name:      name,
		AppSecret: store.NewSnowflakeID().String(),
	}
	app.ID = store.NewSnowflakeID().Int64()
	err = dalApplication.Insert(nil, app)
	if err != nil {
		// 添加应用程序错误
		err = cerrors.GetBusinessError(2003)
		return
	}
	appID = app.ID
	return
}
