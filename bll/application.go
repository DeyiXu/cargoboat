package bll

import (
	"time"

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

// ApplicationPagedItem 应用
type ApplicationPagedItem struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	AppSecret string     `json:"secret"`
	Version   string     `json:"version"`
	ModeCount int        `json:"mode_count"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// QueryPage 获取翻页数据
func (a *ApplicationBll) QueryPage(pageNum, pageSize int, appName string) (data []*ApplicationPagedItem, total int64, err error) {
	var apps []*model.Application
	apps, total, err = dal.Application.QueryPage(pageNum, pageSize, appName)
	if err != nil {
		err = cerrors.GetBusinessError(2007)
		return
	}

	for _, app := range apps {
		data = append(data, &ApplicationPagedItem{
			Name:      app.Name,
			AppSecret: app.AppSecret,
			ID:        app.ID,
			CreatedAt: app.CreatedAt,
			UpdatedAt: app.UpdatedAt,
			DeletedAt: app.DeletedAt,
		})
	}

	return
}
