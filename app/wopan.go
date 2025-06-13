package app

import (
	"github.com/gin-gonic/gin"
	"github.com/xhofe/wopan-sdk-go"
)

func wopanLogin(c *gin.Context) {
	req := struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		Error(c, err)
		return
	}
	w := wopan.Default()
	res, err := w.PcWebLogin(req.Phone, req.Password)
	if err != nil {
		Error(c, err)
		return
	}
	c.JSON(200, res)
}

func wopanVerifyCode(c *gin.Context) {
	req := struct {
		Phone      string `json:"phone"`
		Password   string `json:"password"`
		VerifyCode string `json:"verify_code"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		Error(c, err)
		return
	}
	w := wopan.Default()
	res, err := w.PcLoginVerifyCode(req.Phone, req.Password, req.VerifyCode)
	if err != nil {
		Error(c, err)
		return
	}
	c.JSON(200, res)
}
