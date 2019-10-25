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

func ProjectDelete(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	resp, err := http.Get(viper.GetString("project.url") + "/delete/" + id.Hex())
	if err != nil {
		SendResponse(c, errno.ErrBuildNetwork, nil)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		SendResponse(c, errno.ErrBuildResponse, nil)
		return
	}
	var res Result
	err = json.Unmarshal(body, &res)
	if err != nil {
		SendResponse(c, errno.ErrBuildResponse, nil)
		return
	}
	SendResponse(c, nil, res)
}
