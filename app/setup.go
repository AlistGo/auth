package app

import (
	"os"

	"github.com/gin-gonic/gin"
)

func initVar() {
	// client
	aliOpenClientID = os.Getenv("ALI_OPEN_CLIENT_ID")
	aliOpenClientSecret = os.Getenv("ALI_OPEN_CLIENT_SECRET")
	baiduClientId = os.Getenv("BAIDU_CLIENT_ID")
	baiduClientSecret = os.Getenv("BAIDU_CLIENT_SECRET")
	dropboxClientID = os.Getenv("DROPBOX_CLIENT_ID")
	dropboxClientSecret = os.Getenv("DROPBOX_CLIENT_SECRET")
	appID115 = os.Getenv("APP_ID_115")
}

func Setup(g *gin.RouterGroup) {
	initVar()
	g.GET("/ali/qr", Qr)
	g.POST("/ali/ck", Ck)
	g.POST("/onedrive/get_refresh_token", onedriveToken)
	g.POST("/onedrive/get_site_id", spSiteID)
	g.GET("/baidu/get_refresh_token", baiduToken)
	g.POST("/wopan/login", wopanLogin)
	g.POST("/wopan/verify_code", wopanVerifyCode)
	g.POST("/dropbox/token", dropboxToken)
	aliOpen := g.Group("/ali_open")
	aliOpen.Any("/token", aliAccessToken)
	aliOpen.Any("/refresh", aliAccessToken)
	aliOpen.Any("/code", aliAccessToken)
	aliOpen.Any("/qr", aliQrcode)
	_115 := g.Group("/115")
	_115.Any("/auth_device_code", authDeviceCode115)
	_115.Any("/get_status", getStatus115)
	_115.Any("/get_token", getToken115)
}
