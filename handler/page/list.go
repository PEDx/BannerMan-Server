package page

import (
	. "BannerMan-Server/handler"
	"BannerMan-Server/model"
	"BannerMan-Server/pkg/errno"

	"github.com/gin-gonic/gin"
)

type GetListRequest struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}
type GetListResponse struct {
	PageTotal int64             `json:"pageTotal"`
	PageList  []*model.PageInfo `json:"pageList"`
}

// 先查用户所在的所有组,把组内的用户 ID 数组作为 filter 查询满足($in)这个数组的所有页面(创建者);
// 同时可以加入权限筛选 Permission = 1
// 同时加上 Permission = 2 的所有页面
func List(c *gin.Context) {
	var r GetListRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	pageTotal, pageList, err := model.GetPageList(r.PageSize, (r.Page-1)*r.PageSize)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	// Show the user information.
	SendResponse(c, nil, GetListResponse{
		PageTotal: pageTotal,
		PageList:  pageList,
	})
}
