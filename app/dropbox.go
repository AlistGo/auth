package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	dropboxClientID     string
	dropboxClientSecret string
)

type dropboxTokenReq struct {
	ClientID     string `json:"client_id" form:"client_id"`
	ClientSecret string `json:"client_secret" form:"client_secret"`
	GrantType    string `json:"grant_type" form:"grant_type"`
	RedirectUri  string `json:"redirect_uri" form:"redirect_uri"`
	Code         string `json:"code" form:"code"`
	RefreshToken string `json:"refresh_token" form:"refresh_token"`
}

func dropboxToken(c *gin.Context) {
	var req dropboxTokenReq
	err := c.ShouldBind(&req)
	if err != nil {
		Error(c, err)
		return
	}
	if req.ClientID == "" || req.ClientID == dropboxClientID {
		req.ClientID = dropboxClientID
		req.ClientSecret = dropboxClientSecret
	}
	form := map[string]string{
		"grant_type":    req.GrantType,
		"client_id":     req.ClientID,
		"client_secret": req.ClientSecret,
		//"redirect_uri":  req.RedirectUri,
		//"code":          req.Code,
		//"refresh_token": req.RefreshToken,
	}
	if req.RedirectUri != "" {
		form["redirect_uri"] = req.RedirectUri
	}
	switch req.GrantType {
	case "authorization_code":
		form["code"] = req.Code
	case "refresh_token":
		form["refresh_token"] = req.RefreshToken
	default:
		ErrorStr(c, "Incorrect GrantType")
		return
	}
	res, err := RestyClient.R().
		SetFormData(form).
		Post("https://api.dropboxapi.com/oauth2/token")
	if err != nil {
		Error(c, err, http.StatusInternalServerError)
		return
	}
	c.Status(res.StatusCode())
	JsonBytes(c, res.Bytes())
	return
}
