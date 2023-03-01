package config

import (
	"github.com/spf13/viper"
	"log"
)

type MysqlConfig struct {
	Host     string
	Port     int
	Db       string
	User     string
	Password string
	Timeout  string `default:"10s"`
}

type JwtConfig struct {
	Key        string
	AccessAge  int
	RefreshAge int
}

type Config struct {
	Port    string
	BaseUrl string
	MySql   MysqlConfig
	Jwt     JwtConfig
}

var config *Config

// LoadConfig 加载配置
func LoadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return err
	}
	return nil
}

// GetConfig 获取配置
func GetConfig() *Config {
	return config
}

// InitConfig 初始化配置
func InitConfig(path string) {
	if err := LoadConfig(path); err != nil {
		log.Fatal("加载配置失败: ", err)
	}
}
