package bll

import (
	"time"

	"github.com/cargoboat/cargoboat/dal"
	"github.com/cargoboat/cargoboat/model"
	cerrors "github.com/cargoboat/cargoboat/module/errors"
	"github.com/cargoboat/cargoboat/module/store"
	"github.com/cargoboat/cargoboat/service"
)

type ConfigBll struct {
	errFormat string
}

func NewConfigBll() *ConfigBll {
	return &ConfigBll{
		errFormat: "bll/config/%s:%v",
	}
}

// Edit 编辑配置
func (c *ConfigBll) Edit(appID int64, name, mode, value string) error {
	conf := dal.Config.GetByLastVersion(name, mode, appID)
	if conf == nil {
		return c.add(appID, name, mode, value)
	}
	return c.update(*conf, value)
}

func (c *ConfigBll) add(appID int64, name, mode, value string) error {
	conf := new(model.Config)
	conf.ID = store.NewSnowflakeID().Int64()
	conf.AppID = appID
	conf.Mode = mode
	conf.Name = name
	conf.Value = value
	conf.Version = 1
	if err := dal.Config.Add(conf); err != nil {
		// 添加配置名称错误
		return cerrors.GetBusinessError(2006)
	}
	// 发布
	service.PubConfigChannel <- *conf
	return nil
}

func (c *ConfigBll) update(srcConf model.Config, value string) error {
	srcConf.CreatedAt = time.Now()
	srcConf.UpdatedAt = srcConf.CreatedAt
	srcConf.DeletedAt = nil
	srcConf.ID = store.NewSnowflakeID().Int64()
	srcConf.Value = value
	srcConf.Version++
	if err := dal.Config.Add(&srcConf); err != nil {
		// 添加配置名称错误
		return cerrors.GetBusinessError(2006)
	}
	// 发布
	service.PubConfigChannel <- srcConf
	return nil
}
