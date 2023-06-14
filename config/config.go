package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	App App `yaml:"app" json:"app"`
	Log Log `yaml:"log" json:"log"`
	Db  Db  `yaml:"db" json:"db"`
}

type App struct {
	Host string `yaml:"host" json:"host"`
	Port string `yaml:"port" json:"port"`
}

type Log struct {
	Suffix  string `yaml:"suffix" json:"suffix"`
	MaxSize int    `yaml:"maxSize" json:"maxsize"`
}

type Db struct {
	Mysql string `yaml:"mysql" json:"mysql"`
}

var _config *Config = nil

func GetConfigData() *Config {
	//导入配置文件
	yamlFile, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	//将配置文件读取到结构体中
	err = yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	return _config
}
