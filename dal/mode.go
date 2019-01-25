package dal

import (
	"github.com/cargoboat/cargoboat/model"
	"github.com/jinzhu/gorm"
	"github.com/nilorg/pkg/db"
	"github.com/nilorg/pkg/logger"
)

// ModeDal 模式
type ModeDal struct {
	db        *db.DataBase
	errFormat string
}

// NewModeDal ...
func NewModeDal(db *db.DataBase) *ModeDal {
	return &ModeDal{
		db:        db,
		errFormat: "dal/mode/%s:%v",
	}
}

// Add ...
func (m *ModeDal) Add(mode *model.Mode) (err error) {
	err = m.db.Master().Create(mode).Error
	if err != nil {
		logger.Errorf(m.errFormat, "Add", err)
	}
	return

}

// Update ...
func (m *ModeDal) Update(mode *model.Mode) (err error) {
	err = m.db.Master().Model(&model.Mode{}).Updates(&mode).Error
	if err != nil {
		logger.Errorf(m.errFormat, "Update", err)
	}
	return
}

// GetByID ...
func (m *ModeDal) GetByID(id int64) (mode *model.Mode) {
	mode = new(model.Mode)
	err := m.db.Slave().First(mode, id).Error
	if err == gorm.ErrRecordNotFound {
		mode = nil
	} else if err != nil {
		logger.Errorf(m.errFormat, "GetByID", err)
		mode = nil
	}
	return
}

// GetAllByAppID ...
func (m *ModeDal) GetAllByAppID(appID int64) (modes []*model.Mode) {
	if m.db.Slave().Find(&modes, "app_id = ?", appID).RecordNotFound() {
		return nil
	}
	return
}

// IsExistName ...
func (m *ModeDal) IsExistName(name string, appID int64) (exist bool, err error) {
	var count int
	err = m.db.Slave().Model(&model.Mode{}).Where("name = ? and app_id = ?", name, appID).Count(&count).Error
	if err != nil {
		logger.Errorf(m.errFormat, "IsExistName", err)
		return
	}
	exist = count > 0
	return
}
