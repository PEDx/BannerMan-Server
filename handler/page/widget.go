package page

import (
	. "BannerMan-Server/handler"
	"BannerMan-Server/model"
	"BannerMan-Server/pkg/errno"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetWidgetList(c *gin.Context) {

	err, res := GetWidgetsFromNpm()
	if err != nil {
		SendResponse(c, errno.ErrGetWidgetData, nil)
		return
	}
	SendResponse(c, nil, res)
}
func GetPageWidgetsVersion(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	err, ret := model.GetWidgetVersion(id)
	if err != nil {
		SendResponse(c, errno.ErrPageNotFound, nil)
		return
	}

	SendResponse(c, nil, ret)
}

func GetWidgetsFromNpm() (error, []*model.Widgets) {
	resp, err := http.Get(viper.GetString("packages.api"))
	if err != nil {
		return err, nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}
	var res []*model.Widgets
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err, nil
	}
	return nil, res
}
