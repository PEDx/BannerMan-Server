package page

import (
	. "BannerMan-Server/handler"
	"BannerMan-Server/model"
	"BannerMan-Server/pkg/errno"
	"BannerMan-Server/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetWidgetList(c *gin.Context) {

	err, res := service.GetWidgetsFromNpm()
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
