package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Global struct {
	File File `mapstructure:"file"`
}

type File struct {
	ImgMaxFileSize   int64    `yaml:"imgMaxFileSize"`
	ImgAllowExits    []string `yaml:"imgAllowExits"`
	ExcelMaxFileSize int64    `yaml:"excelMaxFileSize"`
	ExcelAllowExits  []string `yaml:"excelAllowExits"`
	DocMaxFileSize   int64    `yaml:"docMaxFileSize"`
	DocAllowExits    []string `yaml:"docAllowExits"`

	ImgSavePath   string `yaml:"imgSavePath"`
	DocSavePath   string `yaml:"docSavePath"`
	ExcelSavePath string `yaml:"excelSavePath"`

	ImgPrefixUrl   string `yaml:"imgPrefixUrl"`
	DocPrefixUrl   string `yaml:"docPrefixUrl"`
	ExcelPrefixUrl string `yaml:"excelPrefixUrl"`
}

type Email struct {
}

var SettingFile = &Global{}

func SetUP() {
	fmt.Println("开始读取配置文件")
	viper.SetConfigName("config-prod.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Println(viper.AllSettings())

	err = viper.Unmarshal(SettingFile)
	if err != nil {
		fmt.Println("配置文件反序列化错误")
		return
	}

}
