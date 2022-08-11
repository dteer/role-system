package settings

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Mcors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")                                                     // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization") //你想放行的header也可以在后面自行添加
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")                                   //我自己只使用 get post 所以只放行它
		//c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func MstatsdData() gin.HandlerFunc {
	return func(c *gin.Context) {
		start_time := time.Now().UnixMicro()
		c.Next()
		defer func() {
			// 调整了中间件的执行顺序，缺失程序奔溃保护机制，需要加上
			if err := recover(); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		fmt.Println(c.Request.Response.Body)
		end_time := time.Now().UnixMicro()
		interval := end_time - start_time
		statsdObj := StatsdConnect()
		allviews := fmt.Sprintf("%s-%s-%d", Conf.Statsd.Name, "allviews", c.Writer.Status())
		statsdObj.Timing(allviews, interval, 1.0)
		viewstats := fmt.Sprintf("%s-%s-%s", Conf.Statsd.Name, "viewstats", c.HandlerName())
		statsdObj.Timing(viewstats, interval, 1.0)

	}
}
