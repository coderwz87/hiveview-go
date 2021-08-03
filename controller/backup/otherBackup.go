package backup

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hiveview"
	"hiveview/models"
	"hiveview/render"
	"hiveview/utils"
)

// @title 获取全部其他项目备份信息
// @version 1.0.0
// @ID get-all-other-backup-detail
// @description 获取全部其他项目备份信息
// @Success 200 {string} string "{"code":200,"msg":"success","data":[...]}"
// @Tags otherBackup
// @Security ApiKeyAuth
// @Router /allOtherBackupInfo/ [get]
func GetAllOtherBackupDetail(c *gin.Context) {
	data, err := models.GetAllOtherBackupDetail(hiveview.CONFIG.Db)
	if err != nil {
		render.DataError(c, fmt.Sprintf("查询异常:%s", err))
		return
	}
	render.JSON(c, data)
}

// @title 更新其他项目备份信息
// @version 1.0.0
// @ID update-other-backup-detail
// @description 更新其他项目备份信息
// @Success 200 {string} string "{"code":200,"msg":"已更新"}"
// @Param project_name formData  string true "项目名字"
// @Param remote_server formData string true "备份服务器地址"
// @Param remote_dir formData string true "备份目录"
// @Param backup_log formData string true "备份信息"
// @Tags otherBackup
// @Security ApiKeyAuth
// @Router /updateOtherBackupDetail/ [put]
func UpdateOtherBackupDetail(c *gin.Context) {
	ProjectName := c.PostForm("project_name")
	RemoteServer := c.PostForm("remote_server")
	RemoteDir := c.PostForm("remote_dir")
	BackupLog := c.PostForm("backup_log")

	newDetail := new(models.OtherBackupDetail)
	newDetail.ProjectName = ProjectName
	newDetail.RemoteDir = RemoteDir
	newDetail.RemoteServer = RemoteServer
	newDetail.BackupLog = BackupLog
	err := newDetail.UpdateOtherBackupDetail(hiveview.CONFIG.Db)
	if err != nil {
		utils.LogPrint("err", err)
		render.DataError(c, fmt.Sprintf("更新异常:%s", err))
		return
	}
	render.MSG(c, "已更新")
}
