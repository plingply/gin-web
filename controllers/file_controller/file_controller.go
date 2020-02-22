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
	_ "github.com/joho/godotenv/autoload"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cast"
)

var endpoint = os.Getenv("endpoint")
var accessKeyID = os.Getenv("accessKeyId")
var accessKeySecret = os.Getenv("accessKeySecret")
var bucketName = os.Getenv("bucketName")

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

func OSSUpLoad(c *gin.Context) {

	var result models.Result
	file, err := c.FormFile("file")
	dirname := c.PostForm("dirname")
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

	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		result.Code = -1
		result.Msg = err.Error()
		result.Data = ""
		response.Res(c, result)
		os.Exit(-1)
		return
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		result.Code = -1
		result.Msg = err.Error()
		result.Data = ""
		response.Res(c, result)
		os.Exit(-1)
		return
	}

	// 读取本地文件。
	fd, err := os.Open(dst)
	if err != nil {
		result.Code = -1
		result.Msg = err.Error()
		result.Data = ""
		response.Res(c, result)
		os.Exit(-1)
		return
	}

	defer fd.Close()

	url := u1 + prefix
	if dirname != "" {
		url = dirname + url
	}
	// 上传文件流。
	err = bucket.PutObject(url, fd)
	if err != nil {
		result.Code = -1
		result.Msg = err.Error()
		result.Data = ""
		response.Res(c, result)
		os.Exit(-1)
		return
	}

	err = os.Remove(dst)

	if err != nil {
		result.Code = 200
		result.Msg = "上传成功，临时文件删除失败" + err.Error()
		result.Data = os.Getenv("fileEndpoint") + "/" + url
		response.Res(c, result)
		os.Exit(-1)
		return
	}

	result.Code = 200
	result.Msg = "上传成功"
	result.Data = os.Getenv("fileEndpoint") + "/" + url
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
