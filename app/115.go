package app

import (
	"math/rand"

	"github.com/gin-gonic/gin"
	sdk "github.com/xhofe/115-sdk-go"
)

var (
	appID115 string
	letters  = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func authDeviceCode115(c *gin.Context) {
	client := sdk.Default()
	codeVerifier := randomString(44)
	appID := c.Query("app_id")
	if appID == "" {
		appID = appID115
	}
	resp, err := client.AuthDeviceCode(c, appID, codeVerifier)
	if err != nil {
		Error(c, err)
		return
	}
	c.JSON(200, gin.H{
		"code_verifier": codeVerifier,
		"resp":          resp,
	})
}

func getStatus115(c *gin.Context) {
	uid := c.Query("uid")
	time := c.Query("time")
	sign := c.Query("sign")
	if uid == "" || time == "" || sign == "" {
		ErrorStr(c, "invalid params")
		return
	}
	client := sdk.Default()
	resp, err := client.QrCodeStatus(c, uid, time, sign)
	if err != nil {
		Error(c, err)
		return
	}
	c.JSON(200, gin.H{
		"resp": resp,
	})
}

type GetToken115Req struct {
	Uid          string `json:"uid"`
	CodeVerifier string `json:"code_verifier"`
}

func getToken115(c *gin.Context) {
	var req GetToken115Req
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, err)
		return
	}
	if req.CodeVerifier == "" || req.Uid == "" {
		ErrorStr(c, "invalid params")
		return
	}
	client := sdk.Default()
	resp, err := client.CodeToToken(c, req.Uid, req.CodeVerifier)
	if err != nil {
		Error(c, err)
		return
	}
	c.JSON(200, gin.H{
		"resp": resp,
	})
}
