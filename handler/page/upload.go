package page

import (
	. "BannerMan-Server/handler"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

var accessKey string
var secretKey string

func UploadToken(c *gin.Context) {
	bucket := "bannerman"
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"name":"$(fname)"}`,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	SendResponse(c, nil, upToken)
}

func init() {
	b, e := ioutil.ReadFile("./secrets/accessKey.txt")
	if e != nil {
		fmt.Println("read accessKey file error")
		return
	}
	accessKey = string(b)
	c, e := ioutil.ReadFile("./secrets/secretKey.txt")
	if e != nil {
		fmt.Println("read accessKey file error")
		return
	}
	secretKey = string(c)
}
