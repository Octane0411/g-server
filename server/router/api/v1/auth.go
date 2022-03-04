package v1

import (
	"g-server/common/cache"
	"g-server/common/errcode"
	"g-server/common/response"
	"g-server/utils/captcha"
	"g-server/utils/email"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

}

func Register(c *gin.Context) {
	emailAddr := c.Query("email")
	if ok := email.ValidEmail(emailAddr); !ok {
		r := response.NewResponse(c)
		r.ToErrorResponse(errcode.InvalidParams)
	}
	captchaStr := captcha.GetCaptcha()
	email.SendEmail(emailAddr, "", "感谢您注册账号", "你的验证码是："+captchaStr, "游戏名字")
	cache.Cache.Set(emailAddr, []byte(captchaStr))
	r := response.NewResponse(c)
	r.ToResponse(nil)
}

func ConfirmEmail(c *gin.Context) {
	emailAddr := c.Query("email")
	cCaptcha := c.Query("captcha")
	if entry, err := cache.Cache.Get(emailAddr); err != nil {
		r := response.NewResponse(c)
		r.ToErrorResponse(errcode.InvalidParams)
	} else if cCaptcha != string(entry) {
		r := response.NewResponse(c)
		r.ToErrorResponse(errcode.InvalidCaptcha)
	} else {

		r := response.NewResponse(c)
		r.ToResponse(nil)
	}
}
