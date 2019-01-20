package config

import (
	"net/http"

	"github.com/cargoboat/cargoboat/model"
	"github.com/cargoboat/cargoboat/module/store"
	ngin "github.com/nilorg/pkg/gin"
)

// Get 获取
func Get(ctx *ngin.WebAPIContext) {

}

type confModel struct {
	Name  string `json:"name" form:"name" binding:"required"`
	Value string `json:"value" form:"value" binding:"required"`
	// 模式
	Mode string `json:"mode" form:"mode" binding:"required"`
	// 版本
	Version float64 `json:"version" form:"version" binding:"required"`
	// 应用ID
	AppID int64 `json:"app_id" form:"app_id" binding:"required"`
}

// Validation 验证
func (c *confModel) Validation(obj interface{}) error {
	var count int64
	if err := store.DB.Model(&model.Config{}).Where(c).Count(&count).Error; err != nil {
		return errConfigAdd
	}
	if count > 0 {
		return errConfigExist
	}
	return nil
}

// Post 新增
func Post(w *ngin.WebAPIContext) {
	cm := new(confModel)

	if err := w.Bind(&cm); err != nil {
		w.ResultError(err)
		return
	}

	conf := new(model.Config)
	conf.AppID = cm.AppID
	conf.Mode = cm.Mode
	conf.Name = cm.Name
	conf.Value = cm.Value
	conf.Version = cm.Version
	if err := store.DB.Create(conf).Error; err != nil {
		w.ResultError(errConfigAdd)
	} else {
		w.Status(http.StatusCreated)
	}
}

// Delete 删除
func Delete(w *ngin.WebAPIContext) {
	store.DB.Where("id = ?", w.Param("id")).Delete(&model.Config{})
}

// Put 修改
func Put(w *ngin.WebAPIContext) {
	conf := new(model.Config)
	store.DB.Update(conf)
}
