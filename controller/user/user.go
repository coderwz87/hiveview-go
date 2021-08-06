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

func ChangeUserPassword(c *gin.Context) {
	password := c.PostForm("password")
	claims, _ := c.Get("claims")
	username := claims.(*utils.CustomClaims).Username
	var user = new(models.Users)
	user.Username = username
	user.Password = password
	ifSuccess := user.UpdatePassword(hiveview.CONFIG.Db)
	if ifSuccess {
		render.MSG(c, "已修改")
	} else {
		render.DataError(c, "修改失败")
	}

}

func AddUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	var user = new(models.Users)
	user.Username = username
	user.Password = password
	err := user.CreateUser(hiveview.CONFIG.Db)
	if err != nil {
		utils.LogPrint("err", err)
		return
	}
	hiveview.CONFIG.Enforcer.AddPolicy(username, "/api/*", "GET")
	hiveview.CONFIG.Enforcer.AddPolicy(username, "/api/ChangeUserPassword/", "POST")
	hiveview.CONFIG.Enforcer.SavePolicy()
	render.MSG(c, "已创建")

}
