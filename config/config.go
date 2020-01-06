package config

import (
	"github.com/Unknwon/goconfig"
	"os"
)

var Config *goconfig.ConfigFile

func init() {
	env := os.Getenv("env")
	if env == "" {
		env = "dev"
	}
	path, _ := os.Getwd()
	configPath := path + "/config/" + env + ".ini"
	cfg, err := goconfig.LoadConfigFile(configPath)
	if err != nil {
		panic("获取应用配置文件错误")
	}
	Config = cfg
}
