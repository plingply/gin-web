package file_controller

import (
	"fmt"
	"gin-web/models"
	"gin-web/utils/response"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cast"
)

func UpLoad(c *gin.Context) {

	fmt.Println("OSS Go SDK Version: ", oss.Version)

	var result models.Result
	file, err := c.FormFile("file")
	if err != nil {
		result.Code = -1
		result.Msg = err.Error()
		result.Data = nil
		response.Res(c, result)
		return
	}

	u1 := cast.ToString(uuid.NewV4())
	prefix := "." + strings.Split(file.Filename, ".")[1]
	dst := "../file/" + u1 + prefix
	// gin 简单做了封装,拷贝了文件流
	if err1 := SaveUploadedFile(file, dst); err1 != nil {
		result.Code = -1
		result.Msg = err1.Error()
		result.Data = nil
		response.Res(c, result)
		return
	}

	result.Code = 200
	result.Msg = "上传成功"
	result.Data = dst

	response.Res(c, result)
}

func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	//创建 dst 文件
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	// 拷贝文件
	_, err = io.Copy(out, src)
	return err
}
