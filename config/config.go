package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	App App `yaml:"app" json:"app"`
	Log Log `yaml:"log" json:"log"`
}

type App struct {
	Host string `yaml:"host" json:"host"`
	Port string `yaml:"port" json:"port"`
}

type Log struct {
	Suffix  string `yaml:"suffix" json:"suffix"`
	MaxSize int    `yaml:"maxSize" json:"maxsize"`
}

func GetConfigData() *Config {
	//导入配置文件
	yamlFile, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	var _config *Config = nil
	//将配置文件读取到结构体中
	err = yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return _config

}
