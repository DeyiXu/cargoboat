package bll

import (
	"github.com/cargoboat/cargoboat/dal"
	"github.com/cargoboat/cargoboat/model"
	cerrors "github.com/cargoboat/cargoboat/module/errors"
	"github.com/cargoboat/cargoboat/module/store"
)

// ApplicationBll ...
type ApplicationBll struct {
	errFormat string
}

func NewApplicationBll() *ApplicationBll {
	return &ApplicationBll{
		errFormat: "bll/application/%s:%v",
	}
}

// GetAll ...
func (a *ApplicationBll) GetAll() (apps []*model.Application) {
	apps = dal.Application.GetAll()
	return
}

// GetOneByID 获取一个App
func (a *ApplicationBll) GetOneByID(id int64) (app *model.Application) {
	app = dal.Application.GetByID(id)
	return
}

// Add 添加应用
func (a *ApplicationBll) Add(name string) (appID int64, err error) {
	b, err := dal.Application.IsExistName(name)
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
	err = dal.Application.Add(app)
	if err != nil {
		// 添加应用程序错误
		err = cerrors.GetBusinessError(2003)
		return
	}
	appID = app.ID
	return
}
