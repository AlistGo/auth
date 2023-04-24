package alist

import (
	"api.nn.ci/apps/common"
	"api.nn.ci/utils"
	"github.com/gin-gonic/gin"
)

func Qr(c *gin.Context) {
	url := "https://passport.aliyundrive.com/newlogin/qrcode/generate.do"
	res, err := utils.RestyClient.R().SetQueryParams(map[string]string{
		"appName":     "aliyun_drive",
		"fromSite":    "52",
		"appEntrance": "web",
		"_csrf_token": "",
		"umidToken":   "",
		"isMobile":    "false",
		"lang":        "zh_CN",
		"returnUrl":   "",
		"hsiz":        "",
		"bizParams":   "",
		"_bx-v":       "2.0.31",
	}).Get(url)
	if err != nil {
		common.Error(c, err)
		return
	}
	c.Header("Content-Type", "application/json")
	c.Writer.Write(res.Body())
}

type CK struct {
	T  string `json:"t"`
	Ck string `json:"ck"`
}

func Ck(c *gin.Context) {
	var ck CK
	err := c.ShouldBind(&ck)
	if err != nil {
		common.Error(c, err)
		return
	}
	res, err := utils.RestyClient.R().SetFormData(map[string]string{
		"t":            ck.T,
		"ck":           ck.Ck,
		"ua":           "140#ApzoT1O+zzPDRQo245+u33Sc2qq3vOsx37btKgtYp/+IQTwpilmRWL7UklKiXOSxEgmrlBPw4oOKU3hqzzngEzCNOa+xzWz8ijlulFzx2DD3VthqzFHcHbzum51xxD2iVP//lbzx2dfHKCUI1wba7XElyb98FLkGcBq9NLTwSgAzL+yICWq/l5WrYJ+B3qlPPFJg+BxuOJkpm+kszeUq29TiOuclegGQGrpKbFQOPCQE+u94nT7aL8G9Aq84NbL7nhfeFD9BpnzRPrEJrbCbpA3Kk7IsEW3gDIgSC4pQVKuM1VwwGaIuNdotnVtfuCceOFxWedDGMKHlr9NLAu9JKzRJBASFHRNdObSUeSklxZdXIHnupibAkG9mTwAEtajstVuX75Y7icOS5KhgQFP7iNuqEEeARX3DiMkI0pDw0Ybj5Q5JrXCz9AL6CTW3t0Zw5lE68UmECpi1eMwuY46BXykk4ET7+Jm7a+RVUTnWP5vfFV0omNauBNpsVggw/MYMxy4czMfMRiQwglJGBIVw+Mr18S+BAvJzqaUXg+HDUphISFsirUND0/u3zg+FM06Zc6rsVmxE2eSffE3cpgfVYoN/Hf24yFJCOVnVlIEagQF2CPxBQIDL+Q9E/f1l3lfQktqrC0GgxdPNv5ifjzp9IDb3t4h75O2daoJDnKcYhDfbKFvpqUgwkUCzzYspDRPv4XXAhsNq6KQZr3nP1AKdSjEL4XQSAGh4HCE1zHrvKPz93BYl68ZHZig9975vH+/fQlgzMRQE3NRaPBSh1a2If53LnMFj6f1g5OH1ZEPIZBq+K6RSGs6RJJ8NRKibX8weXQEXwVar9UeBKxIwGPW4Nysitb9/Le2NYpEf0oKIrGB/T0AEyieR1BNv8M8pNDIJ9M/lPDyoN4kB5sxD0E+=",
		"appName":      "aliyun_drive",
		"appEntrance":  "web",
		"_csrf_token":  "uJPMkz6XudG40RXo6xCuW5",
		"umidToken":    "6795f5c4caafbf6e9623941b8a3056b3e318c1fd",
		"isMobile":     "false",
		"lang":         "zh_CN",
		"returnUrl":    "",
		"hsiz":         "10918f04a35e8c83cf032e462eb88647",
		"fromSite":     "52",
		"bizParams":    "",
		"navlanguage":  "zh-CN",
		"navUserAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36",
		"navPlatform":  "Win32",
	}).
		SetHeaders(map[string]string{
			"origin":  "https://passport.aliyundrive.com",
			"referer": "https://passport.aliyundrive.com/mini_login.htm?&appName=aliyun_drive",
		}).
		Post("https://passport.aliyundrive.com/newlogin/qrcode/query.do?appName=aliyun_drive&fromSite=52&_bx-v=2.0.31")
	//data := utils.Json.Get(res.Body(), "content", "data")
	//loginResult := data.Get("loginResult").ToString()
	//bizExt := data.Get("bizExt").ToString()
	c.Header("Content-Type", "application/json")
	c.Writer.Write(res.Body())
}
