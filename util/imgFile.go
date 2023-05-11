package util

import (
	"go-file/config"
	"strings"
)

type imgFile struct{}

var ImgUtil = &imgFile{}

func (img imgFile) GetFullPath(path string) string {
	return config.SettingFile.File.ImgPrefixUrl + path
}

func (img imgFile) CheckFileSize(size int64) bool {
	return true
}

func (img imgFile) CheckFileSuffix(suffix string) bool {
	for _, allow := range config.SettingFile.File.ImgAllowExits {
		if strings.ToUpper(allow) == strings.ToUpper(suffix) {
			return true
		}
	}
	return false
}

func (img imgFile) GetSavePath(path string) string {
	return config.SettingFile.File.ImgSavePath + path
}
