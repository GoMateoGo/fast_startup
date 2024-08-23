package conf

import (
	"encoding/json"
	"fmt"
	dbmysql "payconfig/core/db"
	myrd "payconfig/core/redis"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 定义全局变量GlobalCfg，类型为*Config
var (
	GlobalCfg *Config
)

// 定义Config结构体，包含Db、Server、Redis、Sms四个字段
type Config struct {
	Db     dbmysql.MySql // 数据库连接信息
	Server Server        // 服务器信息
	Redis  myrd.RedisCfg // Redis信息
}

// 定义Server结构体，包含AppName、Domain、Port三个字段
type Server struct {
	AppName string // 应用名称
	Domain  string // 域名
	Port    string // 端口号
	Mode    string // 模式
}

// LoadConfig 函数用于加载配置文件
func LoadConfig() {
	viper.SetConfigName("config")                // 配置文件名（不包含扩展名）
	viper.SetConfigType("yaml")                  // 配置文件类型
	viper.AddConfigPath(".")                     // 配置文件所在路径
	viper.WatchConfig()                          // 开始监听配置文件变化
	viper.OnConfigChange(ChangeConfig)           // 配置改变Hook
	if err := viper.ReadInConfig(); err != nil { // 读取配置文件
		panic("加载配置文件失败:" + err.Error())
	}
	// 将读取到的数据反序列化为配置变量
	err := viper.Unmarshal(&GlobalCfg)
	if err != nil {
		panic("反序列化配置文件失败:" + err.Error())
	}

	printConfig()
}

// ChangeConfig 函数用于处理配置文件更改事件
func ChangeConfig(e fsnotify.Event) {
	fmt.Println("配置文件已更改:", e.Name)
	// 在这里重新加载配置
	err := viper.Unmarshal(&GlobalCfg)
	if err != nil {
		fmt.Printf("配置文件解码错误: %s\n", err)
	}

	printConfig()
}

// 打印出格式化的JSON结构
func printConfig() {
	jsonCfg, err := json.MarshalIndent(GlobalCfg, "", "  ")
	if err != nil {
		fmt.Printf("转换配置为JSON格式失败: %v", err)
	}
	fmt.Println(string(jsonCfg))
}
