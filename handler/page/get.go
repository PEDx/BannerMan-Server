package page

import (
	. "BannerMan-Server/handler"
	"BannerMan-Server/model"
	"BannerMan-Server/pkg/errno"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Get(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	var p *model.Page

	pageInfo, err := p.GetPageInfoByID((id))
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	// Show the user information.
	SendResponse(c, nil, pageInfo)
}
