package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	baiduClientId     string
	baiduClientSecret string
	baiduCallbackUri  = "https://alist.nn.ci/tool/baidu/callback"
)

func baiduToken(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		ErrorStr(c, "no code")
		return
	}
	res, err := RestyClient.R().
		Get(fmt.Sprintf(
			"https://openapi.baidu.com/oauth/2.0/token?grant_type=authorization_code&code=%s&client_id=%s&client_secret=%s&redirect_uri=%s",
			code, baiduClientId, baiduClientSecret, baiduCallbackUri))
	if err != nil {
		Error(c, err)
		return
	}
	JsonBytes(c, res.Bytes())
}
