package user

import (
	. "BannerMan-Server/handler"
	"BannerMan-Server/model"
	"BannerMan-Server/pkg/errno"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := (&model.User{
		Username: r.Username,
		Password: r.Password,
	}).New()
	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	if err := u.GetUserByUsername(r.Username); err != nil {
		SendResponse(c, errno.ErrUserAlreadyExisted, nil)
		return
	}
	if err := u.CreateUser(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
