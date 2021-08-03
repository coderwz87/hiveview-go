package commonTools

import (
	"github.com/gin-gonic/gin"
	"hiveview"
	"hiveview/models"
	"hiveview/render"
	"hiveview/utils"
)

// @title 获取全部常用连接信息
// @version 1.0.0
// @ID get-all-common-link
// @description 获取全部常用连接信息
// @Success 200 {string} string "{"code":200,"msg":"success","data":[...]}"
// @Tags common
// @Security ApiKeyAuth
// @Router /common/link/ [get]
func GetAllCommonLink(c *gin.Context) {
	result, err := models.GetAllLink(hiveview.CONFIG.Db)
	if err != nil {
		utils.LogPrint("err", err)
		render.DataError(c, "获取失败")
		return
	}
	render.JSON(c, result)
}
