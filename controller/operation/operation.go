package operation

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hiveview"
	"hiveview/models"
	"hiveview/render"
	"hiveview/utils"
	"strconv"
)

func GetOperationDetail(c *gin.Context) {
	id := c.Param("id")
	result := new(models.OperationDetail)
	ID, err := strconv.Atoi(id)
	if err != nil {
		render.DataError(c, err.Error())
		return
	}
	result.ID = uint(ID)
	err = result.GetOperationDetailByID(hiveview.CONFIG.Db)
	if err != nil {
		utils.LogPrint("err", err)
		render.DataError(c, err.Error())
		return
	}
	render.JSON(c, result)
}

func OperationApprove(c *gin.Context) {
	action := c.DefaultPostForm("action", "deny")
	operationID := c.DefaultPostForm("id", "")
	if len(operationID) == 0 {
		render.DataError(c, "操作信息错误")
		return
	}
	if action == "pass" {
		err := utils.AppOperation(operationID, action)
		if err != nil {
			utils.LogPrint("err", err)
			render.DataError(c, "操作失败")
			return
		}
		render.MSG(c, "已开始执行")
		return
	} else if action == "deny" {
		result := new(models.OperationDetail)
		ID, err := strconv.Atoi(operationID)
		if err != nil {
			return
		}
		result.ID = uint(ID)
		_ = result.GetOperationDetailByID(hiveview.CONFIG.Db)
		result.State = "已拒绝"
		_ = result.UpdateOperationStateByID(hiveview.CONFIG.Db)
		render.MSG(c, "已拒绝")
		return
	} else {
		render.DataError(c, "参数有误")
		return
	}

}

func GetAllOperationDetail(c *gin.Context) {
	data, err := models.GetAllOperationDetail(hiveview.CONFIG.Db)
	if err != nil {
		render.DataError(c, fmt.Sprintf("查询异常:%s", err))
		return
	}
	render.JSON(c, data)
}
