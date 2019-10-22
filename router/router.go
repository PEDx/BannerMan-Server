package router

import (
	"net/http"

	"BannerMan-Server/handler/page"
	"BannerMan-Server/handler/sd"
	"BannerMan-Server/handler/user"
	"BannerMan-Server/router/middleware"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
	apiv1 := g.Group("/api/v1")
	{
		apiv1.POST("user", user.Create)       // 创建用户
		apiv1.DELETE("user/:id", user.Delete) // 删除用户
		apiv1.PUT("user/:id", user.Update)    // 更新用户
		apiv1.GET("user", user.List)          // 用户列表
		apiv1.GET("user/:username", user.Get) // 获取指定用户的详细信息

		apiv1.POST("page", page.Create)    // 创建页面
		apiv1.PUT("page/:id", page.Update) // 更新页面
		apiv1.GET("page/:id", page.Get)    // 获取指定用户的详细信息
		apiv1.GET("build/:id", page.Build) // 构建页面
	}
	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
