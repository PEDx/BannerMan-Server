package page

import (
	. "BannerMan-Server/handler"
	"BannerMan-Server/model"
	"BannerMan-Server/pkg/errno"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PageWidgetUpdata struct {
	ID             primitive.ObjectID `json:"id"`
	WidgetNameList []string           `json:"widgetNameList"`
}

// 更新机制:
// 更新全部.包括添加最新和删除不存在了的组件
// 更新单个或多个组件版本

func Version(c *gin.Context) {
	patchMap := model.WidgetsVersionMap{}
	var w PageWidgetUpdata
	if err := c.Bind(&w); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	// 更新组件版本信息
	err, res := GetWidgetsFromNpm()
	if err != nil {
		SendResponse(c, errno.ErrGetWidgetData, nil)
		return
	}

	for _, widgetName := range w.WidgetNameList {
		for _, widget := range res {
			if widget.Name == widgetName {
				patchMap[widgetName] = widget.Version
			}
		}
	}
	// 为空就全量更新
	if len(w.WidgetNameList) == 0 {
		for _, widget := range res {
			patchMap[widget.Name] = widget.Version
		}
	}
	if err := model.UpdateWidgetVersion(w.ID, &patchMap); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}
