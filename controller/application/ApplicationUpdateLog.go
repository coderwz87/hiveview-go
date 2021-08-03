package application

import (
	"github.com/gin-gonic/gin"
	"hiveview"
	"hiveview/models"
	"hiveview/render"
	"hiveview/utils"
	"time"
)

func GetApplicationUpdateLog(c *gin.Context) {
	AllAppUpdateLog, err := models.GetAllAppUpdateLog(hiveview.CONFIG.Db)
	if err != nil {
		utils.LogPrint("err", err)
		render.DataError(c, "查询失败")
		return
	}
	resData := utils.CombineAppUpdateLog(AllAppUpdateLog)
	render.JSON(c, resData)
}

func CreateApplicationUpdateLog(c *gin.Context) {
	MD5 := c.DefaultPostForm("md5", "")
	AppName := c.DefaultPostForm("app_name", "")
	Host := c.DefaultPostForm("host", "")
	now := time.Now()
	var ApplicationUpdateData = models.ApplicationUpdateLog{
		MD5:        MD5,
		AppName:    AppName,
		Host:       Host,
		UpdateTime: now.Format("2006-01-02 15:04"),
	}
	err := ApplicationUpdateData.CreateAppUpdateLog(hiveview.CONFIG.Db)
	if err != nil {
		utils.LogPrint("err", err)
		render.DataError(c, "创建失败")
		return
	}
	render.JSON(c, "已创建")
}

func SearchApplicationUpdateLog(c *gin.Context) {
	searchKey := c.Query("search")
	resultList := models.GetAppUpdateLogByFuzzyAppName(hiveview.CONFIG.Db, searchKey)
	resData := utils.CombineAppUpdateLog(resultList)
	render.JSON(c, resData)
}
