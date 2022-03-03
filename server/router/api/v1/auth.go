package v1

import (
	"g-server/utils/email"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

}

func Register(c *gin.Context) {
	//captchaStr := captcha.GetCaptcha()
	emailAddr := c.Query("email")
	if ok := email.ValidEmail(emailAddr); !ok {

	}
}

func ConfirmEmail(c *gin.Context) {

}
