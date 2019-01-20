package application

import (
	"github.com/cargoboat/cargoboat/model"
	"github.com/cargoboat/cargoboat/module/store"
	"github.com/jinzhu/gorm"
	"github.com/nilorg/pkg/logger"
)

const errFormat = "dal/application/%s:%v"

// Insert ...
func Insert(tran *gorm.DB, e *model.Application) error {
	db := store.SwitchDB(tran)
	if err := db.Create(e).Error; err != nil {
		logger.Errorf(errFormat, "Insert", err)
		return err
	}
	return nil
}

// Update ...
func Update(e *model.Application) {
	err := store.DB.Model(&e).Updates(&e).Error
	if err != nil {
		logger.Errorf(errFormat, "Update", err)
	}
}

// GetByID return Application by ID
func GetByID(id int64) *model.Application {
	app := new(model.Application)
	err := store.DB.First(&app, "id = ?", id).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		logger.Errorf(errFormat, "GetByID", err)
		return nil
	}
	return app
}

// GetAll ...
func GetAll() []*model.Application {
	var apps []*model.Application
	if store.DB.Find(apps).RecordNotFound() {
		return nil
	}
	return apps
}

// IsExistName 判断是否存在
func IsExistName(name string) (b bool, err error) {
	var count int
	err = store.DB.Model(&model.Application{}).Where("name = ?", name).Count(&count).Error
	if err != nil {
		logger.Errorf(errFormat, "IsExistName", err)
		return
	}
	return count > 0, nil
}
