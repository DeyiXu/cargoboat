package dal

import (
	"github.com/jinzhu/gorm"
	"github.com/nilorg/pkg/logger"

	"github.com/cargoboat/cargoboat/model"
	"github.com/nilorg/pkg/db"
)

// ApplicationDal ...
type ApplicationDal struct {
	db        *db.DataBase
	errFormat string
}

// NewApplicationDal ...
func NewApplicationDal(db *db.DataBase) *ApplicationDal {
	return &ApplicationDal{
		db:        db,
		errFormat: "dal/application/%s:%v",
	}
}

// GetPaged 获取翻页数据
func (a *ApplicationDal) QueryPage(pageNum, pageSize int, appName string) (data []*model.Application, total int64, err error) {
	sqlCmd := "deleted_at IS NULL"
	sqlValues := make([]interface{}, 0)
	if appName != "" {
		sqlCmd += " and name like ?"
		sqlValues = append(sqlValues, "%"+appName+"%")
	}
	total, err = db.SelectPageData(a.db.Slave(), &data, model.ApplicationTableName, "id", pageNum, pageSize, sqlCmd, sqlValues)
	if err != nil {
		logger.Errorf(a.errFormat, "QueryPage", err)
	}
	return
}

// GetAll ...
func (a *ApplicationDal) GetAll() (apps []*model.Application) {
	if a.db.Slave().Find(apps).RecordNotFound() {
		return nil
	}
	return
}

// Add ...
func (a *ApplicationDal) Add(app *model.Application) (err error) {
	err = a.db.Master().Create(app).Error
	if err != nil {
		logger.Errorf(a.errFormat, "Add", err)
	}
	return
}

// Update ...
func (a *ApplicationDal) Update(app *model.Application) (err error) {
	err = a.db.Master().Model(&app).Updates(&app).Error
	if err != nil {
		logger.Errorf(a.errFormat, "Update", err)
	}
	return
}

// GetByID return Application by ID
func (a *ApplicationDal) GetByID(id int64) (app *model.Application) {
	err := a.db.Slave().First(&app, "id = ?", id).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		logger.Errorf(a.errFormat, "GetByID", err)
		return nil
	}
	return
}

// IsExistName 判断是否存在
func (a *ApplicationDal) IsExistName(name string) (b bool, err error) {
	var count int
	err = a.db.Slave().Model(&model.Application{}).Where("name = ?", name).Count(&count).Error
	if err != nil {
		logger.Errorf(a.errFormat, "IsExistName", err)
		return
	}
	b = count > 0
	return
}
