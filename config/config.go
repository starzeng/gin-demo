package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Name         string `mapstructure:"name"`
		Mode         string `mapstructure:"mode"`
		Host         string `mapstructure:"host"`
		Port         string `mapstructure:"port"`
		RelativePath string `mapstructure:"relative_path"`
	} `mapstructure:"server"`

	MySQL struct {
		Host         string `mapstructure:"host"`
		Port         int    `mapstructure:"port"`
		User         string `mapstructure:"user"`
		Password     string `mapstructure:"password"`
		DBName       string `mapstructure:"dbname"`
		MaxIdleConns int    `mapstructure:"max_idle_conns"`
		MaxOpenConns int    `mapstructure:"max_open_conns"`
		LogDebug     bool   `mapstructure:"log_debug"`
	} `mapstructure:"mysql"`

	Redis struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
	} `mapstructure:"redis"`
}

var AppConfig *Config

func LoadConfig() {
	viper.SetConfigName("config") // 文件名
	viper.SetConfigType("yaml")   // 文件类型
	viper.AddConfigPath("./conf") // 配置路径

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}

	AppConfig = &Config{}
	err = viper.Unmarshal(AppConfig)
	if err != nil {
		log.Fatalf("解析配置失败: %v", err)
	}

	fmt.Println("配置加载成功")
}

// GetServerAddr 获取server地址
func GetServerAddr() string {
	return AppConfig.Server.Host + ":" + AppConfig.Server.Port
}
