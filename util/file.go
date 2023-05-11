package util

import (
	uuid "github.com/satori/go.uuid"
	"os"
	"strings"
)

type FileUtil interface {
	GetFullPath(path string) string
	CheckFileSize(size int64) bool
	CheckFileSuffix(suffix string) bool
	GetSavePath(path string) string
}

func GetSuffix(name string) string {
	split := strings.Split(name, ".")
	return "." + split[len(split)-1]
}

func UuidRandom() string {
	return uuid.NewV4().String()
}

func DeleteFile(name string) bool {
	path := ImgUtil.GetSavePath(name)
	err := os.RemoveAll(path)
	if err != nil {
		return false
	}

	return true
}
