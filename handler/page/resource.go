package page

import (
	. "BannerMan-Server/handler"
	"BannerMan-Server/model"
	"BannerMan-Server/pkg/errno"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PushPageResource(c *gin.Context) {

	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	resource := model.Resource{}
	if err := c.Bind(&resource); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	err := model.PushPageResource(id, &resource)
	if err != nil {
		SendResponse(c, errno.ErrPageNotFound, nil)
		return
	}

	SendResponse(c, nil, nil)
}

type PullRes struct {
	Key string `json:"key"`
}

func PullPageResource(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	r := PullRes{}
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	err := model.PullPageResource(id, r.Key)
	if err != nil {
		SendResponse(c, errno.ErrPageNotFound, nil)
		return
	}

	SendResponse(c, nil, nil)
}
func GetPageResource(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var ret *[]*model.Resource
	err, ret := model.GetPageResource(id)
	if err != nil {
		SendResponse(c, errno.ErrPageNotFound, nil)
		return
	}

	SendResponse(c, nil, ret)
}
