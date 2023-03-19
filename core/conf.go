package core

import (
	"fmt"
	"gin_vue_project/config"
	"gin_vue_project/global"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

// 读取yaml文件的配置
func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error"))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success.")
	global.Config = c
}
