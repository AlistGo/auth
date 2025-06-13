package app

import (
	"github.com/gin-gonic/gin"
	"resty.dev/v3"
)

func Error(c *gin.Context, err error, codes ...int) {
	ErrorStr(c, err.Error(), codes...)
}

func ErrorStr(c *gin.Context, str string, codes ...int) {
	code := 200
	if len(codes) > 0 {
		code = codes[0]
	}
	c.JSON(code, gin.H{
		"error": str,
	})
}

func JsonBytes(c *gin.Context, body []byte) {
	c.Header("Content-Type", gin.MIMEJSON)
	c.Writer.Write(body)
}

func ErrorJson(c *gin.Context, data interface{}, code ...int) {
	code = append(code, 500)
	c.JSON(code[0], data)
}

func RealIP(c *gin.Context) string {
	// Cloudflare
	if ip := c.GetHeader("Cf-Connecting-Ip"); ip != "" {
		return ip
	}
	return c.ClientIP()
}

var RestyClient = resty.New()
