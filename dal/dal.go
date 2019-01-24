package dal

import "github.com/cargoboat/cargoboat/module/store"

var (
	Application Applicationer = NewApplicationDal(store.DataBase)
	Config      Configer      = NewConfigDal(store.DataBase)
)
