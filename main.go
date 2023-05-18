package main

import (
	"github.com/gin-gonic/gin"
	"go-file/api/v1/upload"
	"go-file/config"
	"go-file/util"
	"net/http"
)

// go-file
func main() {
	// 初始化配置文件
	config.SetUP()

	r := gin.Default()
	r.Use(config.Cors())
	r.StaticFS("/imgs/", http.Dir(util.ImgUtil.GetSavePath("")))
	r.GET("/remove/:imgName", upload.RemoveFile)
	r.POST("/imgUpload", upload.ImgUpload)
	r.Run(":8080")
}
