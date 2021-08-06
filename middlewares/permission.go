package middlewares

import (
	"github.com/gin-gonic/gin"
	"hiveview"
	"hiveview/render"
	"hiveview/utils"
)

func PermissionMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		path := c.Request.URL.Path
		if path == "/api/login/" {
			c.Next()
			return
		}
		claims, _ := c.Get("claims")
		role := claims.(*utils.CustomClaims).Username
		isPass, err := hiveview.CONFIG.Enforcer.Enforce(role, path, method)
		if err != nil {
			utils.LogPrint("err", err)
			render.DataError(c, "鉴权失败")
			return
		}
		if isPass {
			c.Next()
		} else {
			render.PermissionError(c)
			c.Abort()
			return
		}
	}

}
