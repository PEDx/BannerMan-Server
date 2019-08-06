package user

import (
	. "BannerMan-Server/handler"
	"BannerMan-Server/model"
	"BannerMan-Server/pkg/errno"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	userList, err := model.GetUserList(10, 0)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	// Show the user information.
	SendResponse(c, nil, userList)
}
