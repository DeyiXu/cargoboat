package dal

import (
	"github.com/cargoboat/cargoboat/model"
	"github.com/jinzhu/gorm"
	"github.com/nilorg/pkg/db"
	"github.com/nilorg/pkg/logger"
)

// VersionDal 版本
type VersionDal struct {
	db        *db.DataBase
	errFormat string
}

// NewVersionDal ...
func NewVersionDal(db *db.DataBase) *VersionDal {
	return &VersionDal{
		db:        db,
		errFormat: "dal/version/%s:%v",
	}
}

// Add ...
func (v *VersionDal) Add(version *model.Version) (err error) {
	err = v.db.Master().Create(version).Error
	if err != nil {
		logger.Errorf(v.errFormat, "Add", err)
	}
	return

}

// Update ...
func (v *VersionDal) Update(version *model.Version) (err error) {
	err = v.db.Master().Model(&model.Version{}).Updates(&version).Error
	if err != nil {
		logger.Errorf(v.errFormat, "Update", err)
	}
	return
}

// GetByID ...
func (v *VersionDal) GetByID(id int64) (version *model.Version) {
	version = new(model.Version)
	err := v.db.Slave().First(version, id).Error
	if err == gorm.ErrRecordNotFound {
		version = nil
	} else if err != nil {
		logger.Errorf(v.errFormat, "GetByID", err)
		version = nil
	}
	return
}

// GetAllByAppID ...
func (v *VersionDal) GetAllByAppID(appID int64) (versions []*model.Version) {
	if v.db.Slave().Find(&versions, "app_id = ?", appID).RecordNotFound() {
		return nil
	}
	return
}

// IsExistName ...
func (v *VersionDal) IsExistName(name string, appID int64) (exist bool, err error) {
	var count int
	err = v.db.Slave().Model(&model.Version{}).Where("name = ? and app_id = ?", name, appID).Count(&count).Error
	if err != nil {
		logger.Errorf(v.errFormat, "IsExistName", err)
		return
	}
	exist = count > 0
	return
}
