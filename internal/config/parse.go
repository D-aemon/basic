package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	DB struct {
		User   string `yaml:"user"`
		Passwd string `yaml:"passwd"`
		IP     string `yaml:"ip"`
		Port   string `yaml:"port"`
		DBName string `yaml:"dbName"`
	} `yaml:"db"`
	Log struct {
		Level   int    `yaml:"level"`
		Kinds   string `yaml:"kinds"`
		Project string `yaml:"project"`
	} `yaml:"log"`
	Port struct {
		GrpcPort string `yaml:"grpc"`
		HttpPort string `yaml:"http"`
	} `yaml:"port"`
}

var Cfg Config

func init() {
	file, err := os.Open("/app/config.yml")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// 解析 YAML 文件
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&Cfg)
	if err != nil {
		fmt.Printf("Error decoding YAML: %v\n", err)
		return
	}
}
