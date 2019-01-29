package cargoboat

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"

	"github.com/nilorg/pkg/logger"

	"github.com/fsnotify/fsnotify"
)

type GroupEventFunc func(groupName string)

// Cargoboat ...
type Cargoboat struct {
	redisClient *redis.Client
	fileWatcher *fsnotify.Watcher
	config      map[string]*viper.Viper
	configRWM   sync.RWMutex
	DirName     string
}

// newViper ...
func newViper(filename string) *viper.Viper {
	v := viper.New()
	v.SetConfigFile(filename)
	if err := v.ReadInConfig(); err != nil {
		return nil
	} else {
		v.WatchConfig()
	}
	return v
}

// NewCargoboat ...
func NewCargoboat(dirname string, redisClient *redis.Client) (watcher *Cargoboat, err error) {
	var fileWatcher *fsnotify.Watcher
	fileWatcher, err = fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	return &Cargoboat{
		redisClient: redisClient,
		fileWatcher: fileWatcher,
		DirName:     dirname,
		config:      make(map[string]*viper.Viper),
	}, nil
}

func (c *Cargoboat) initConfig() (err error) {
	var fileInfos []os.FileInfo
	fileInfos, err = ioutil.ReadDir(c.DirName)
	if err != nil {
		return
	}
	// 配置文件初始化
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}
		c.createViper(fileInfo.Name(), newViper(path.Join(c.DirName, fileInfo.Name())))
	}
	// 初始化Redis
	return c.initRedis()
}

func (c *Cargoboat) initRedis() (err error) {
	keysCmd := c.redisClient.Keys("cargoboat.*")
	keys := keysCmd.Val()
	c.redisClient.Del(keys...)

	for groupKey, groupValue := range c.config {
		addGroupResult := c.redisClient.SAdd(RedisKeyGroupList, fmt.Sprintf(RedisKeyGroupFormat, groupKey))
		if addGroupResult.Err() != nil {
			err = addGroupResult.Err()
			return
		}
		groupValueKeys := groupValue.AllKeys()
		for i := 0; i < len(groupValueKeys); i++ {
			confKey := groupValueKeys[i]
			c.redisClient.HSet(fmt.Sprintf(RedisKeyGroupFormat, groupKey), confKey, groupValue.Get(confKey))
		}
	}
	return
}

// Start 启动
func (c *Cargoboat) Start() (err error) {
	err = c.initConfig()
	if err != nil {
		return
	}
	go func() {
		for {
			select {
			case event := <-c.fileWatcher.Events:
				if !strings.HasSuffix(event.Name, "___jb_tmp___") && !strings.HasSuffix(event.Name, "___jb_old___") {
					groupName := filepath.Base(event.Name)
					switch event.Op {
					case fsnotify.Create:
						logger.Infof("create:%s", event.Name)
						c.createViper(groupName, newViper(event.Name))
					case fsnotify.Remove:
						logger.Infof("remove:%s", event.Name)
						c.removeViper(groupName)
					}
				}
			case err := <-c.fileWatcher.Errors:
				logger.Errorln("error:1111111", err)
			}
		}
	}()
	err = c.fileWatcher.Add(c.DirName)
	return
}

// Stop 停止
func (c *Cargoboat) Stop() error {
	return c.fileWatcher.Close()
}

// AllGroups 查询所有配置分组
func (c *Cargoboat) AllGroups() []string {
	var groups []string
	for key := range c.config {
		groups = append(groups, key)
	}
	return groups
}

// Viper 获取配置集合
func (c *Cargoboat) Viper(groupName string) *viper.Viper {
	c.configRWM.RLock()
	defer c.configRWM.RUnlock()
	return c.config[groupName]
}

// createViper 创建配置集合
func (c *Cargoboat) createViper(groupName string, v *viper.Viper) {
	if v == nil {
		return
	}
	c.configRWM.Lock()
	defer c.configRWM.Unlock()
	logger.Debugf("groupName Check Is Exist:%s", groupName)
	_, exist := c.config[groupName]
	if exist {
		logger.Debugf("groupName Is Exist:%s", groupName)
		return
	}
	v.OnConfigChange(func(e fsnotify.Event) {
		c.onGroupChange(path.Base(e.Name))
	})
	c.config[groupName] = v
}

// removeViper 移除配置集合
func (c *Cargoboat) removeViper(groupName string) {
	c.configRWM.Lock()
	defer c.configRWM.Unlock()
	delete(c.config, groupName)
	c.onGroupDelete(groupName)
}

func (c *Cargoboat) onGroupChange(groupName string) {
	v := c.Viper(groupName)
	groupValueKeys := v.AllKeys()
	c.redisClient.HDel(fmt.Sprintf(RedisKeyGroupFormat, groupName))
	for i := 0; i < len(groupValueKeys); i++ {
		confKey := groupValueKeys[i]
		c.redisClient.HSet(fmt.Sprintf(RedisKeyGroupFormat, groupName), confKey, v.Get(confKey))
	}
	// 发布更改消息
	c.redisClient.Publish(RedisKeyChangeChannel, groupName)
}
func (c *Cargoboat) onGroupDelete(groupName string) {
	// 删除 group set集合
	c.redisClient.Del(RedisKeyGroupList)
	for _, gkey := range c.AllGroups() {
		c.redisClient.SAdd(RedisKeyGroupList, gkey)
	}
	c.redisClient.Del(fmt.Sprintf(RedisKeyGroupFormat, groupName))
	// 发布删除消息
	c.redisClient.Publish(RedisKeyDeleteChannel, groupName)
}
