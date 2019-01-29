package cargoboat

import "github.com/spf13/viper"

// Cargoboater ...
type Cargoboater interface {
	Start() (err error)
	Stop() (err error)
	AllGroups() []string
	Viper(groupName string) *viper.Viper
}
