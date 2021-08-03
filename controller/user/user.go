package user

import (
	"github.com/gin-gonic/gin"
	"hiveview"
	"hiveview/models"
	"hiveview/render"
	"hiveview/utils"
)

//用户登录模块
func Login(c *gin.Context) {
	var form models.Users
	if err := c.ShouldBind(&form); err != nil {
		render.ParamError(c, err.Error())
		return
	}
	ifPassed := form.Verify(hiveview.CONFIG.Db)
	if !ifPassed {
		render.LoginError(c, "incorrect username or password")
		return
	}
	token, err := utils.GenerateToken(form)
	if err != nil {
		render.LoginError(c, "create token err")
		return
	}
	data := map[string]string{"token": token}

	render.JSON(c, data)
}
