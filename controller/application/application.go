package application

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/nilorg/sdk/convert"

	"github.com/cargoboat/cargoboat/errors"
	"github.com/cargoboat/cargoboat/model"
	"github.com/cargoboat/cargoboat/module/store"
	ngin "github.com/nilorg/pkg/gin"
)

// Get 获取
func Get(ctx *ngin.WebAPIContext) {
	var apps []model.Application
	if store.DB.Find(apps).RecordNotFound() {
		ctx.ResultError(errors.ErrApplicationNotFound)
	} else {
		ctx.JSON(200, map[string]interface{}{
			"data":  apps,
			"total": len(apps),
		})
	}
}

// GetOne 获取一个
func GetOne(w *ngin.WebAPIContext) {
	var config model.Config
	if store.DB.First(&config, w.Param("id")).RecordNotFound() {
		w.Status(http.StatusNotFound)
	} else {
		w.JSON(200, config)
	}
}

// GetConfigs 获取配置文件
func GetConfigs(w *ngin.WebAPIContext) {
	appID := w.Param("id")
	name, nameExist := w.GetQuery("name")
	query := "app_id = ?"
	values := []interface{}{
		appID,
	}
	if nameExist {
		query += " and name = ?"
		values = append(values, name)
	}

	var configs []model.Config
	if store.DB.Where(query, values...).Find(&configs).RecordNotFound() {
		w.ResultError(errors.ErrApplicationNotFound)
	} else {
		w.JSON(200, map[string]interface{}{
			"data":  configs,
			"total": len(configs),
		})
	}
}

// Post 创建
func Post(w *ngin.WebAPIContext) {
	app := new(model.Application)
	h := md5.New()
	h.Write([]byte(convert.ToString(time.Now().Unix())))
	app.AppSecret = hex.EncodeToString(h.Sum(nil))

	if err := store.DB.Create(app).Error; err != nil {
		w.ResultError(errors.ErrApplicationCreateFailure)
	} else {
		w.JSON(200, app)
	}
}

// Delete 删除
func Delete(w *ngin.WebAPIContext) {
	if err := store.DB.Where("id = ?", "").Delete(&model.Application{}).Error; err != nil {
		w.ResultError(errors.ErrApplicationDelFailure)
	} else {
		w.Status(http.StatusNoContent)
	}
}
