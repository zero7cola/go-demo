package bootstrap

import (
	"github.com/gin-gonic/gin"
	"go1/app/http/middlewares"
	"go1/routes"
	"net/http"
	"strings"
)

func SetupRoute(router *gin.Engine) {

	// 注册全局中间件
	registerGlobalMiddleWare(router)

	// 注册路由
	routes.RegisterAPIRoutes(router)

	// 配置 404 路由
	setup404Handler(router)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(middlewares.Logger(), middlewares.Recovery())
}

func setup404Handler(router *gin.Engine) {

	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "页面 404")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"error_code": 404,
				"error_msg":  "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
