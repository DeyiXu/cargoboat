package store

import (
	"log"
	"os"

	"github.com/bwmarrin/snowflake"
	"github.com/jinzhu/gorm"
	"github.com/nilorg/pkg/logger"

	// postgres驱动
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	// DB 数据库连接
	DB *gorm.DB
	// SnowflakeNode 雪花ID节点
	snowflakeNode *snowflake.Node
)

func initDB() {
	// 初始化数据库
	db, err := gorm.Open("postgres", viper.GetString("postgres.address"))
	if err != nil {
		logger.Fatalf(
			"初始化 postgres 连接失败: %s \n",
			errors.Wrap(err, "打开 postgres 连接失败"),
		)
		os.Exit(-1)
	}
	err = db.DB().Ping()
	if err != nil {
		logger.Fatalf(
			"初始化 postgres 连接失败: %s \n",
			errors.Wrap(err, "Ping postgres 失败"),
		)
		os.Exit(-1)
	}

	db.LogMode(viper.GetBool("postgres.log"))

	db.DB().SetMaxOpenConns(viper.GetInt("postgres.max_open"))
	db.DB().SetMaxIdleConns(viper.GetInt("postgres.max_idle"))
	// db.DB().SetConnMaxLifetime(time.Hour)

	DB = db
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
	err := DB.Close()
	if err != nil {
		log.Println(err)
	}
}

// SwitchDB 切换数据库
func SwitchDB(tran *gorm.DB) *gorm.DB {
	if tran != nil {
		return tran
	}
	return DB
}

// NewSnowflakeID 获取雪花ID
func NewSnowflakeID() snowflake.ID {
	return snowflakeNode.Generate()
}
