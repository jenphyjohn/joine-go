package config

import (
	"github.com/kylelemons/go-gypsy/yaml"
	"log"
)

var (
	File = New()
)

/**
 * 返回单例实例
 * @method New
 */
func New() *yaml.File {
	file, err := yaml.ReadFile("resources/application.yml")
	if err != nil {
		log.Fatal("读取配置文件失败：", err)
	}
	return file
}

func GetValue(key string) string {
	value, err := File.Get(key)
	if err != nil {
		log.Printf(key + " 属性不存在")
	}
	return value
}
