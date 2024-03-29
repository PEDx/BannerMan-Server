package page

import (
	. "BannerMan-Server/handler"
	"BannerMan-Server/model"
	"BannerMan-Server/pkg/errno"
	"BannerMan-Server/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateResponse struct {
	ID   primitive.ObjectID `json:"id"`
	Name string             `json:"name"`
}

func Create(c *gin.Context) {
	widgetMap := map[string]string{}
	var r model.PageInfo
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	// 创建页面时凝固组件版本信息
	err, widgetList := service.GetWidgetsFromNpm()
	if err != nil {
		SendResponse(c, errno.ErrGetWidgetData, nil)
		return
	}
	for _, widget := range widgetList {
		widgetMap[widget.Name] = widget.Version
	}
	p := (&model.Page{
		Name:           r.Name,
		Creater:        r.Creater, // 从 JWT 中获取
		CreaterName:    r.CreaterName,
		ExpiryStart:    r.ExpiryStart.UTC(),
		ExpiryEnd:      r.ExpiryEnd.UTC(),
		WidgetsVersion: widgetMap,
		Permission:     r.Permission,
	}).New()

	if err := p.CreatePage(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	rsp := CreateResponse{
		ID:   p.ID,
		Name: p.Name,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
