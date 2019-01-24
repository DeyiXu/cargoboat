package dal

import (
	"github.com/cargoboat/cargoboat/module/store"
)

var (
	Application Applicationer
	Config      Configer
)

func Init() {
	Application = NewApplicationDal(store.DataBase)
	Config = NewConfigDal(store.DataBase)
}
