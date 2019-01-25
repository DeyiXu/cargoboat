package dal

import (
	"github.com/cargoboat/cargoboat/module/store"
)

var (
	Application Applicationer
	Config      Configer
	Mode        Modeer
	Version     Versioner
)

func Init() {
	Application = NewApplicationDal(store.DataBase)
	Config = NewConfigDal(store.DataBase)
	Mode = NewModeDal(store.DataBase)
	Version = NewVersionDal(store.DataBase)
}
