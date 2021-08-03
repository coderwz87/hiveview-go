package middlewares

import (
	"github.com/gin-gonic/gin"
	"hiveview/render"
	"hiveview/utils"
	"strings"
)

// jwt认证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		path := strings.TrimSpace(c.Request.URL.Path)
		if path == "/api/login/" {
			c.Next()
			return
		}
		if token == "" {
			respondWithError(c, render.CODE_ERR_NO_LOGIN, "not login")
			return
		}
		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			respondWithError(c, render.CODE_ERR_NO_LOGIN, err.Error())
			return
		}
		c.Set("claims", claims)
	}
}

func respondWithError(c *gin.Context, code int, message string) {
	render.CustomerError(c, code, message)
	c.Abort()
}
