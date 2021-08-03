package render

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CODE_OK               = 200
	CODE_MSG              = 201
	CODE_ERR_PARAM        = 400
	CODE_ERR_LOGIN_FAILED = 401
	CODE_ERR_NO_LOGIN     = 402
	CODE_ERR_NO_DATA      = 403
)

func ParamError(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": CODE_ERR_PARAM,
		"msg":  message,
	})
}

func LoginError(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": CODE_ERR_LOGIN_FAILED,
		"msg":  message,
	})
}

func JSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": CODE_OK,
		"msg":  "success",
		"data": data,
	})
}

func MSG(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": CODE_MSG,
		"msg":  msg,
	})
}

func CustomerError(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
	})
}
func DataError(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": CODE_ERR_NO_DATA,
		"msg":  message,
	})
}

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": CODE_OK,
		"msg":  "success",
	})
}
