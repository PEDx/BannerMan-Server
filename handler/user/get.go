package user

import (
	. "BannerMan-Server/handler"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {

	// Show the user information.
	SendResponse(c, nil, nil)
}
