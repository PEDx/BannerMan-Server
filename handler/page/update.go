package page

import (
	. "BannerMan-Server/handler"
	"BannerMan-Server/model"
	"BannerMan-Server/pkg/errno"
	"BannerMan-Server/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Update(c *gin.Context) {
	patchMap := map[string]interface{}{}
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	if err := c.Bind(&patchMap); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	tags := utils.GetAllTagValue(&model.PgaeUpdateData{}, "json")
	for k := range patchMap {
		ok, _ := utils.InArray(k, tags)
		if !ok {
			delete(patchMap, k)
		}
	}

	p := &model.PgaeUpdateInfo{
		ID:         id,
		EditorID:   id,
		EditorName: "test",
	}
	if err := model.UpdatePage(p, &patchMap); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	// Show the user information.
	SendResponse(c, nil, nil)
}
