package bll

// Configer ...
type Configer interface {
	// Edit 编辑配置
	Edit(appID int64, name, mode, value string) (err error)
}
