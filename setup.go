package alist

import (
	"os"
	"strconv"
	"time"

	"github.com/axiaoxin-com/ratelimiter"
	"github.com/gin-gonic/gin"
)

func initVar() {
	s := os.Getenv("ALI_LIMIT_MINUTES")
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	aliMinutes = v
	s = os.Getenv("ALI_LIMIT_MAX")
	v, err = strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	aliMax = v
	// client
	aliClientID = os.Getenv("ALI_DRIVE_CLIENT_ID")
	aliClientSecret = os.Getenv("ALI_DRIVE_CLIENT_SECRET")
	baiduClientId = os.Getenv("BAIDU_CLIENT_ID")
	baiduClientSecret = os.Getenv("BAIDU_CLIENT_SECRET")
}

func Setup(g *gin.RouterGroup) {
	initVar()
	g.GET("/ali/qr", Qr)
	g.POST("/ali/ck", Ck)
	g.POST("/onedrive/get_refresh_token", onedriveToken)
	g.POST("/onedrive/get_site_id", spSiteID)
	g.GET("/baidu/get_refresh_token", baiduToken)
	aliOpen := g.Group("/ali_open")
	aliOpen.Any("/limit", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"minutes": aliMinutes,
			"max":     aliMax,
		})
	})
	aliOpenLimit := aliOpen.Group("")
	aliOpenLimit.Use(ratelimiter.GinMemRatelimiter(ratelimiter.GinRatelimiterConfig{
		LimitKey: func(ctx *gin.Context) string {
			return ctx.ClientIP()
		},
		LimitedHandler: func(ctx *gin.Context) {
			ctx.JSON(429, gin.H{
				"code":    "Too Many Requests",
				"message": "Too Many Requests",
				"error":   "Too Many Requests",
			})
			ctx.Abort()
		},
		TokenBucketConfig: func(context *gin.Context) (time.Duration, int) {
			return time.Duration(aliMinutes) * time.Minute, aliMax
		},
	}))
	aliOpenLimit.Any("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ClientIP": c.ClientIP(),
			"RemoteIP": c.RemoteIP(),
		})
	})
	aliOpenLimit.Any("/token", aliAccessToken)
	aliOpenLimit.Any("/refresh", aliAccessToken)
	aliOpenLimit.Any("/code", aliAccessToken)
	aliOpenLimit.Any("/qr", aliQrcode)
}
