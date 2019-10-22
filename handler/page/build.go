package page

import (
	. "BannerMan-Server/handler"
	"BannerMan-Server/pkg/errno"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Result map[string]interface{}

func Build(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	resp, err := http.Get(viper.GetString("build.api") + id.Hex())
	if err != nil {
		SendResponse(c, errno.ErrPageNotFound, nil)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		SendResponse(c, errno.ErrPageNotFound, nil)
		return
	}
	var res Result
	err = json.Unmarshal(body, &res)
	if err != nil {
		SendResponse(c, errno.ErrPageNotFound, nil)
		return
	}
	SendResponse(c, nil, res)
}
