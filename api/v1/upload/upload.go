package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-file/util"
	"log"
	"strings"
)

// RemoveFile 删除文件接口
func RemoveFile(c *gin.Context) {
	url := c.Param("imgName")
	url = strings.TrimSpace(url)
	// 是否为空或者长度不达标的假路径或无后缀
	suffix := util.GetSuffix(url)

	if url == "" || len(url) < 8 || suffix == "."+url {
		c.JSON(400, gin.H{
			"success": false,
			"msg":     "参数格式有问题",
		})
		return
	}

	strs := strings.Split(url, "/")
	fileName := strs[len(strs)-1]
	file := util.DeleteFile(fileName)

	if !file {
		c.JSON(400, gin.H{
			"success": false,
			"msg":     "文件删除失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"msg":     "文件删除成功",
	})
}

func ImgUpload(c *gin.Context) {
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Printf("fail try get file: %v", err)
	}

	// 检查文件类型
	suffix := util.GetSuffix(header.Filename)
	if !util.ImgUtil.CheckFileSuffix(suffix) {
		c.JSON(400, gin.H{
			"success": false,
			"msg":     "文件不是图片类型",
		})
		return
	}
	// 检查文件大小
	if !util.ImgUtil.CheckFileSize(header.Size) {
		c.JSON(400, gin.H{
			"success": false,
			"msg":     "文件大小超过10MB",
		})
		return
	}

	//生成uuid
	uuid := util.UuidRandom()

	// 合成新filename
	newFileName := uuid + suffix

	err = c.SaveUploadedFile(header, util.ImgUtil.GetSavePath(newFileName))
	if err != nil {
		fmt.Printf("上传失败: %v", err)
		c.JSON(400, gin.H{
			"success": false,
			"msg":     "上传失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"msg":     "上传成功",
		"data":    util.ImgUtil.GetFullPath(newFileName),
	})
}
