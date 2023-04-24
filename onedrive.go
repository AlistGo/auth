package alist

import (
	"encoding/base64"
	"net/url"
	"strings"

	"api.nn.ci/apps/common"
	"api.nn.ci/utils"
	"github.com/gin-gonic/gin"
)

type Zone struct {
	Oauth string
	Api   string
}

var zones = map[string]Zone{
	"global": {
		Oauth: "https://login.microsoftonline.com",
		Api:   "https://graph.microsoft.com",
	},
	"cn": {
		Oauth: "https://login.chinacloudapi.cn",
		Api:   "https://microsoftgraph.chinacloudapi.cn",
	},
	"us": {
		Oauth: "https://login.microsoftonline.us",
		Api:   "https://graph.microsoft.us",
	},
	"de": {
		Oauth: "https://login.microsoftonline.de",
		Api:   "https://graph.microsoft.de",
	},
}

func onedriveToken(c *gin.Context) {
	req := struct {
		Code   string `json:"code"`
		Client string `json:"client"`
	}{}
	err := c.ShouldBind(&req)
	if err != nil {
		common.Error(c, err)
		return
	}
	data, err := base64.StdEncoding.DecodeString(req.Client)
	if err != nil {
		common.Error(c, err)
		return
	}
	clients := strings.Split(string(data), "::")
	if len(clients) < 3 {
		common.ErrorStr(c, "client error")
		return
	}
	if zone, ok := zones[clients[2]]; ok {
		res, err := utils.RestyClient.R().
			SetFormData(map[string]string{
				"client_id":     clients[0],
				"client_secret": clients[1],
				"code":          req.Code,
				"grant_type":    "authorization_code",
				"redirect_uri":  "https://alist.nn.ci/tool/onedrive/callback",
			}).
			Post(zone.Oauth + "/common/oauth2/v2.0/token")
		if err != nil {
			common.Error(c, err)
			return
		}
		common.JsonBytes(c, res.Body())
		return
	}
	common.ErrorStr(c, "zone doesn't exist")
	return
}

func spSiteID(c *gin.Context) {
	req := struct {
		AccessToken string `json:"access_token"`
		SiteUrl     string `json:"site_url"`
		Zone        string `json:"zone"`
	}{}
	err := c.ShouldBind(&req)
	if err != nil {
		common.Error(c, err)
		return
	}
	u, err := url.Parse(req.SiteUrl)
	if err != nil {
		common.Error(c, err)
		return
	}
	siteName := u.Path
	if zone, ok := zones[req.Zone]; ok {
		res, err := utils.RestyClient.R().
			SetHeader("Authorization", "Bearer "+req.AccessToken).
			Get(zone.Api + "/v1.0/sites/root:/" + siteName)
		if err != nil {
			common.Error(c, err)
			return
		}
		common.JsonBytes(c, res.Body())
		return
	}
	common.ErrorStr(c, "zone doesn't exist")
	return
}
