package config

import (
	"github.com/cargoboat/cargoboat/model"
	"github.com/cargoboat/cargoboat/module/store"
	"github.com/jinzhu/gorm"
	"github.com/nilorg/pkg/logger"
)

const errFormat = "dal/config/%s:%v"

// Insert 添加
func Insert(tran *gorm.DB, e *model.Config) error {
	db := store.SwitchDB(tran)
	if err := db.Create(e).Error; err != nil {
		logger.Errorf(errFormat, "Insert", err)
		return err
	}
	return nil
}

// Update ...
func Update(e *model.Config) {
	err := store.DB.Model(&e).Updates(&e).Error
	if err != nil {
		logger.Errorf(errFormat, "Update", err)
	}
}

// GetByID return user by ID
func GetByID(id int64) *model.Config {
	conf := new(model.Config)
	err := store.DB.First(&conf, "id = ?", id).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		logger.Errorf(errFormat, "GetByID", err)
		return nil
	}
	return conf
}

// IsExistName 判断是否存在
func IsExistName(name, mode string, appID int64) (b bool, err error) {
	var count int
	err = store.DB.Model(&model.Config{}).Where("name = ? and mode = ? and app_id = ?", name, mode, appID).Count(&count).Error
	if err != nil {
		logger.Errorf(errFormat, "IsExistName", err)
		return
	}
	return count > 0, nil
}

// GetByLastVersion 获取最后一个版本的配置
func GetByLastVersion(name, mode string, appID int64) *model.Config {
	conf := new(model.Config)
	err := store.DB.Last(&conf, "name = ? and mode = ? and app_id = ?", name, mode, appID).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		logger.Errorf(errFormat, "GetByLastVersion", err)
		return nil
	}
	return conf
}
