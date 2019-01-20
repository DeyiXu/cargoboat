package config

import (
	"time"

	dalConfig "github.com/cargoboat/cargoboat/dal/config"
	"github.com/cargoboat/cargoboat/model"
	cerrors "github.com/cargoboat/cargoboat/module/errors"
	"github.com/cargoboat/cargoboat/module/store"
	"github.com/cargoboat/cargoboat/service"
)

// Edit 编辑配置
func Edit(appID int64, name, mode, value string) error {
	conf := dalConfig.GetByLastVersion(name, mode, appID)
	if conf == nil {
		return add(appID, name, mode, value)
	}
	return update(*conf, value)
}

func add(appID int64, name, mode, value string) error {
	conf := new(model.Config)
	conf.ID = store.NewSnowflakeID().Int64()
	conf.AppID = appID
	conf.Mode = mode
	conf.Name = name
	conf.Value = value
	conf.Version = 1
	if err := dalConfig.Insert(nil, conf); err != nil {
		// 添加配置名称错误
		return cerrors.GetBusinessError(2006)
	}
	// 发布
	service.PubConfigChannel <- *conf
	return nil
}

func update(srcConf model.Config, value string) error {
	srcConf.CreatedAt = time.Now()
	srcConf.UpdatedAt = srcConf.CreatedAt
	srcConf.DeletedAt = nil
	srcConf.ID = store.NewSnowflakeID().Int64()
	srcConf.Value = value
	srcConf.Version++
	if err := dalConfig.Insert(nil, &srcConf); err != nil {
		// 添加配置名称错误
		return cerrors.GetBusinessError(2006)
	}
	// 发布
	service.PubConfigChannel <- srcConf
	return nil
}
