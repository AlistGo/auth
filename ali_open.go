package alist

import (
	"fmt"
	"strings"

	"api.nn.ci/apps/common"
	"api.nn.ci/utils"
	"github.com/gin-gonic/gin"
)

var (
	aliClientID     string
	aliClientSecret string
	aliMinutes      int
	aliMax          int
)

type AliAccessTokenReq struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
	RefreshToken string `json:"refresh_token"`
}

type AliAccessTokenErr struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func aliAccessToken(c *gin.Context) {
	var req AliAccessTokenReq
	err := c.ShouldBind(&req)
	if err != nil {
		common.ErrorJson(c, AliAccessTokenErr{
			Code:    "InternalError",
			Message: err.Error(),
			Error:   err.Error(),
		})
		return
	}
	if req.ClientID == "" {
		req.ClientID = aliClientID
		req.ClientSecret = aliClientSecret
	}
	if req.GrantType != "authorization_code" && req.GrantType != "refresh_token" {
		common.ErrorJson(c, AliAccessTokenErr{
			Code:    "Invalid request",
			Message: "Incorrect GrantType",
			Error:   "Incorrect GrantType",
		}, 400)
		return
	}
	if len(req.RefreshToken) == 32 {
		common.ErrorJson(c, AliAccessTokenErr{
			Code:    "Invalid request",
			Message: "You should use the token that request with aliyundrive open insted of aliyundrive",
			Error:   "You should use the token that request with aliyundrive open insted of aliyundrive",
		}, 400)
		return
	}
	if req.GrantType == "authorization_code" && req.Code == "" {
		common.ErrorJson(c, AliAccessTokenErr{
			Code:    "Invalid request",
			Message: "Code missed",
			Error:   "Code missed",
		}, 400)
		return
	}
	if req.GrantType == "refresh_token" && strings.Count(req.RefreshToken, ".") != 2 {
		common.ErrorJson(c, AliAccessTokenErr{
			Code:    "Invalid request",
			Message: "Incorrect refresh_token or missed",
			Error:   "Incorrect refresh_token or missed",
		}, 400)
		return
	}
	var e AliAccessTokenErr
	res, err := utils.RestyClient.R().SetBody(req).SetError(&e).Post("https://openapi.aliyundrive.com/oauth/access_token")
	if err != nil {
		common.ErrorJson(c, AliAccessTokenErr{
			Code:    "InternalError",
			Message: err.Error(),
			Error:   err.Error(),
		})
		return
	}
	if e.Code != "" {
		e.Error = fmt.Sprintf("%s: %s", e.Code, e.Message)
		common.ErrorJson(c, e, res.StatusCode())
		return
	}
	common.JsonBytes(c, res.Body())
}

type aliQrcodeReq struct {
	ClientID     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	Scopes       []string `json:"scopes"`
}

func aliQrcode(c *gin.Context) {
	var req aliQrcodeReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(500, AliAccessTokenErr{
			Code:    "InternalError",
			Message: err.Error(),
			Error:   err.Error(),
		})
		return
	}
	if req.ClientID == "" {
		req.ClientID = aliClientID
		req.ClientSecret = aliClientSecret
	}
	if req.Scopes == nil || len(req.Scopes) == 0 {
		req.Scopes = []string{"user:base", "file:all:read", "file:all:write"}
	}
	var e AliAccessTokenErr
	res, err := utils.RestyClient.R().SetBody(req).SetError(&e).Post("https://openapi.aliyundrive.com/oauth/authorize/qrcode")
	if err != nil {
		c.JSON(500, AliAccessTokenErr{
			Code:    "InternalError",
			Message: err.Error(),
			Error:   err.Error(),
		})
		return
	}
	if e.Code != "" {
		e.Error = fmt.Sprintf("%s: %s", e.Code, e.Message)
		c.JSON(res.StatusCode(), e)
		return
	}
	common.JsonBytes(c, res.Body())
}
