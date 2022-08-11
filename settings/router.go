package settings

import (
	"fmt"
	"net/http"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(Conf.RunMode)
	Engine = gin.New()
	middlewareHandle(Engine)
	baseRouter := Engine.Group("/role-system")
	routerHandle(baseRouter)
	return Engine
}

// middlewareHandle 中间件注册
func middlewareHandle(engine *gin.Engine) {
	engine.Use(gin.Logger(), MstatsdData(), gin.Recovery(), Mcors())
	engine.Use(sentrygin.New(sentrygin.Options{Repanic: true}))
}

// routerHandle 业务路由
func routerHandle(router *gin.RouterGroup) {
	// 心跳检测
	router.GET("/health_check", func(c *gin.Context) {
		fmt.Println(c.Request.Response.Body)
		c.String(http.StatusOK, "success")
	})

	// 业务部分

}
