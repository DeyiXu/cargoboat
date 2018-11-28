package config

import (
	"net/http"

	"github.com/cargoboat/cargoboat/model"
	"github.com/cargoboat/cargoboat/module/store"
	"github.com/dkeng/pkg/context/gin"
)

// Get 获取
func Get(w *gin.WrapContext) {

}

type confModel struct {
	Name  string `json:"name" form:"name" binding:"required"`
	Value string `json:"value" form:"value" binding:"required"`
	// 模式
	Mode byte `json:"mode" form:"mode" binding:"required"`
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
func Post(w *gin.WrapContext) {
	cm := new(confModel)
	if !w.BindValidation(cm,nil) {
		return
	}

	conf := new(model.Config)
	conf.AppID = cm.AppID
	conf.Mode = cm.Mode
	conf.Name = cm.Name
	conf.Value = cm.Value
	conf.Version = cm.Version
	if err := store.DB.Create(conf).Error; err != nil {
		w.ErrorJSON(errConfigAdd)
	} else {
		w.Status(http.StatusCreated)
	}
}

// Delete 删除
func Delete(w *gin.WrapContext) {
	store.DB.Where("id = ?", w.Param("id")).Delete(&model.Config{})
}

// Put 修改
func Put(w *gin.WrapContext) {
	conf := new(model.Config)
	store.DB.Update(conf)
}
