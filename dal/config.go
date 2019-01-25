package dal

import (
	"github.com/cargoboat/cargoboat/model"
	"github.com/jinzhu/gorm"
	"github.com/nilorg/pkg/db"
	"github.com/nilorg/pkg/logger"
)

// ConfigDal ...
type ConfigDal struct {
	db        *db.DataBase
	errFormat string
}

// NewConfigDal ...
func NewConfigDal(db *db.DataBase) *ConfigDal {
	return &ConfigDal{
		db:        db,
		errFormat: "dal/config/%s:%v",
	}
}

// Add 添加
func (c *ConfigDal) Add(conf *model.Config) (err error) {
	err = c.db.Master().Create(conf).Error
	if err != nil {
		logger.Errorf(c.errFormat, "Add", err)
	}
	return
}

// Update ...
func (c *ConfigDal) Update(e *model.Config) (err error) {
	err = c.db.Master().Model(&model.Config{}).Updates(&e).Error
	if err != nil {
		logger.Errorf(c.errFormat, "Update", err)
	}
	return
}

// GetByID return user by ID
func (c *ConfigDal) GetByID(id int64) (conf *model.Config) {
	conf = new(model.Config)
	err := c.db.Slave().First(conf, id).Error
	if err == gorm.ErrRecordNotFound {
		conf = nil
	} else if err != nil {
		logger.Errorf(c.errFormat, "GetByID", err)
		conf = nil
	}
	return
}

// IsExistName 判断是否存在
func (c *ConfigDal) IsExistName(name, mode string, appID int64) (exist bool, err error) {
	var count int
	err = c.db.Slave().Model(&model.Config{}).Where("name = ? and mode = ? and app_id = ?", name, mode, appID).Count(&count).Error
	if err != nil {
		logger.Errorf(c.errFormat, "IsExistName", err)
		return
	}
	exist = count > 0
	return
}

// GetByLastVersion 获取最后一个版本的配置
func (c *ConfigDal) GetByLastVersion(name, mode string, appID int64) (conf *model.Config) {
	err := c.db.Slave().Last(&conf, "name = ? and mode = ? and app_id = ?", name, mode, appID).Error
	if err == gorm.ErrRecordNotFound {
		conf = nil
	} else if err != nil {
		logger.Errorf(c.errFormat, "GetByLastVersion", err)
		conf = nil
	}
	return
}
