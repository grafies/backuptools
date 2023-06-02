package model

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	SourceDir SourceDir `yaml:"sourceDir"` //源目录
	TargetDir TargetDir `yaml:"targetDir"` //目标目录
	BackName  string    `yaml:"backName"`  //备份名称
	KeepFiles int       `yaml:"keepFiles"` //备份保留数
}

type SourceDir struct {
	Path string `yaml:"path"`
}

type TargetDir struct {
	Path string `yaml:"path"`
}

// ServerConfig 传入配置文件路径
func ServerConfig(configPath string) *Config {
	dataBytes, err := os.ReadFile(configPath)
	if err != nil {
		log.Println("文件读取失败：", err)
	}

	c := &Config{}
	if err = yaml.Unmarshal(dataBytes, &c); err != nil {
		log.Println("解析yaml文件失败：", err)
	}
	return c
}
