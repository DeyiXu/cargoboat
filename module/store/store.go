package store

import (
	"github.com/nilorg/pkg/db"

	"github.com/bwmarrin/snowflake"
	"github.com/nilorg/pkg/logger"

	// postgres驱动
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var (
	// DataBase 数据库
	DataBase *db.DataBase
	// SnowflakeNode 雪花ID节点
	snowflakeNode *snowflake.Node
)

func initDB() {
	dbConf := db.DataBaseConfig{
		DBType:        "postgres",
		MasterAddress: viper.GetString("postgres.address"),
		LogFlag:       viper.GetBool("postgres.log"),
		MaxOpen:       viper.GetInt("postgres.max_open"),
		MaxIdle:       viper.GetInt("postgres.max_idle"),
		SlaveAddress: []string{
			viper.GetString("postgres.address"),
		},
	}
	DataBase = db.NewDataBase(dbConf)
}

func initID() {
	// 设置雪花ID节点
	node, err := snowflake.NewNode(viper.GetInt64("system.snowflake_node"))
	if err != nil {
		logger.Fatalf("snowflake:%v\n", err)
	}
	snowflakeNode = node
}

// Start 启动存储
func Start() {
	initDB()
	initID()
}

// Close 关闭
func Close() {
	DataBase.Close()
}

// NewSnowflakeID 获取雪花ID
func NewSnowflakeID() snowflake.ID {
	return snowflakeNode.Generate()
}
