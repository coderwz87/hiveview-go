package backup

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hiveview"
	"hiveview/models"
	"hiveview/render"
	"hiveview/utils"
	"strconv"
)

type DeleteDetailIDs struct {
	DetailIDs []uint `json:"detail_ids"`
}

// @title 获取全部Mysql备份信息
// @version 1.0.0
// @ID get-all-mysql-backup-detail
// @description 获取全部Mysql备份信息
// @Success 200 {string} string "{"code":200,"msg":"success","data":[...]}"
// @Tags mysqlBackup
// @Security ApiKeyAuth
// @Router /allMysqlBackupInfo/ [get]
func GetAllMysqlBackupDetail(c *gin.Context) {
	data, err := models.GetAllMysqlBackupDetail(hiveview.CONFIG.Db)
	if err != nil {
		render.DataError(c, fmt.Sprintf("查询异常:%s", err))
		return
	}
	render.JSON(c, data)
}

// @title 添加Mysql备份信息
// @version 1.0.0
// @ID add-mysql-backup-detail
// @description 添加Mysql备份信息
// @Success 200 {string} string "{"code":200,"msg":"已创建"}"
// @Param db_name formData  string true "数据库实例名字"
// @Param db_port formData  string true "实例端口号"
// @Param remote_server formData string true "备份服务器地址"
// @Param remote_dir formData string true "备份目录"
// @Tags mysqlBackup
// @Security ApiKeyAuth
// @Router /addMysqlBackupDetail/ [post]
func AddMysqlBackupDetail(c *gin.Context) {
	DbName := c.PostForm("db_name")
	DbPort := c.PostForm("db_port")
	RemoteServer := c.PostForm("remote_server")
	RemoteDir := c.PostForm("remote_dir")
	if len(DbName) == 0 || len(DbPort) == 0 || len(RemoteServer) == 0 || len(RemoteDir) == 0 {
		render.DataError(c, fmt.Sprintf("参数有误"))
		return
	}
	if !utils.JudgeType(RemoteServer) {
		render.DataError(c, fmt.Sprintf("异地服务器不是一个ip地址"))
		return
	}
	if !utils.JudgeAlive(RemoteServer) {
		render.DataError(c, fmt.Sprintf("异地服务器不通"))
		return
	}
	newDetail := new(models.MysqlBackupDetail)
	newDetail.DbName = DbName
	newDetail.DbPort = DbPort
	newDetail.RemoteDir = RemoteDir
	newDetail.RemoteServer = RemoteServer
	err := newDetail.CreateMysqlBackupDetail(hiveview.CONFIG.Db)
	if err != nil {
		render.DataError(c, fmt.Sprintf("创建失败：%s", err))
		return
	}
	go utils.CheckMysqlBackup(DbName, DbPort, RemoteServer, RemoteDir)
	render.MSG(c, "已创建")
}

// @title 删除单条Mysql备份信息
// @version 1.0.0
// @ID delete-mysql-backup-detail
// @description 删除单条Mysql备份信息
// @Success 200 {string} string "{"code":200,"msg":"success"}"
// @Param id path  string true "id"
// @Tags mysqlBackup
// @Security ApiKeyAuth
// @Router /mysqlBackupDetail/{id}/ [delete]
func DeleteMysqlBackupDetail(c *gin.Context) {
	id := c.Param("id")
	result := new(models.MysqlBackupDetail)
	ID, err := strconv.Atoi(id)
	if err != nil {
		render.DataError(c, err.Error())
		return
	}
	result.ID = uint(ID)
	err = result.DeleteMysqlBackupDetailByID(hiveview.CONFIG.Db)
	if err != nil {
		render.DataError(c, fmt.Sprintf("删除异常:%s", err))
		return
	}
	render.Success(c)

}

// @title 删除全部Mysql备份信息
// @version 1.0.0
// @ID delete-all-mysql-backup-detail
// @description 删除全部Mysql备份信息
// @Success 200 {string} string "{"code":200,"msg":"已全部删除"}"
// @Param id formData  DeleteDetailIDs true "id"s
// @Tags mysqlBackup
// @Security ApiKeyAuth
// @Router /deleteAllMysqlBackupDetail/ [delete]
func DeleteAllMysqlBackupDetail(c *gin.Context) {
	var deleteIDs DeleteDetailIDs
	err := c.BindJSON(&deleteIDs)
	if err != nil {
		utils.LogPrint("err", err)
		render.DataError(c, fmt.Sprintf("%s", err))
		return
	}
	for _, v := range deleteIDs.DetailIDs {
		result := new(models.MysqlBackupDetail)
		result.ID = v
		err = result.DeleteMysqlBackupDetailByID(hiveview.CONFIG.Db)
		if err != nil {
			utils.LogPrint("err", err)
			render.DataError(c, fmt.Sprintf("%s", err))
			return
		}
	}
	render.MSG(c, "已全部删除")
}

// @title 更新mysql备份信息
// @version 1.0.0
// @ID update-mysql-backup-detail
// @description 更新mysql备份信息
// @Success 200 {string} string "{"code":200,"msg":"已更新"}"
// @Param db_name formData  string true "数据库实例名"
// @Param db_port formData  string true "端口号"
// @Param remote_server formData string true "备份服务器地址"
// @Param remote_dir formData string true "备份目录"
// @Param backup_log formData string true "备份信息"
// @Tags mysqlBackup
// @Security ApiKeyAuth
// @Router /updateMysqlBackupDetail/ [put]
func UpdateMysqlBackupDetail(c *gin.Context) {
	DbName := c.PostForm("db_name")
	DbPort := c.PostForm("db_port")
	RemoteServer := c.PostForm("remote_server")
	RemoteDir := c.PostForm("remote_dir")
	BackupLog := c.PostForm("backup_log")

	newDetail := new(models.MysqlBackupDetail)
	newDetail.DbName = DbName
	newDetail.DbPort = DbPort
	newDetail.RemoteDir = RemoteDir
	newDetail.RemoteServer = RemoteServer
	newDetail.BackupLog = BackupLog
	err := newDetail.UpdateMysqlBackupDetail(hiveview.CONFIG.Db)
	if err != nil {
		utils.LogPrint("err", err)
		render.DataError(c, fmt.Sprintf("更新异常:%s", err))
		return
	}
	render.MSG(c, "已更新")
}

// @title 搜索mysql备份信息
// @version 1.0.0
// @ID search-mysql-backup-detail
// @description 搜索mysql备份信息
// @Success 200 {string} string "{"code":200,"msg":"success","data":[...]}"
// @Tags mysqlBackup
// @Security ApiKeyAuth
func SearchMysqlBackupDetail(c *gin.Context) {
	searchKey := c.Query("search")
	resultList := models.GetMysqlBackupDetailByFuzzySearchKey(hiveview.CONFIG.Db, searchKey)
	render.JSON(c, resultList)

}

// @title 获取单个mysql备份信息
// @version 1.0.0
// @ID get-mysql-backup-detail
// @description 获取单台资产信息
// @Success 200 {string} string "{"code":200,"msg":"success","data":{object}}"
// @Param id path  string true "id"
// @Tags mysqlBackup
// @Security ApiKeyAuth
// @Router /mysqlBackupDetail/{id}/ [get]
func GetMysqlBackupDetail(c *gin.Context) {
	id := c.Param("id")
	result := new(models.MysqlBackupDetail)
	ID, err := strconv.Atoi(id)
	if err != nil {
		render.DataError(c, err.Error())
		return
	}
	result.ID = uint(ID)
	err = result.GetMysqlBackupDetailByID(hiveview.CONFIG.Db)

	if err != nil {
		utils.LogPrint("err", err)
		render.DataError(c, err.Error())
		return
	}
	render.JSON(c, result)
}

// @title 更新mysql基础备份信息
// @version 1.0.0
// @ID get-mysql-backup-base-detail
// @description 更新mysql基础备份信息
// @Success 200 {string} string "{"code":200,"msg":"已更新"}"
// @Param id path  string true "id"
// @Tags mysqlBackup
// @Param db_name formData  string true "数据库实例名"
// @Param db_port formData  string true "端口号"
// @Param remote_server formData string true "备份服务器地址"
// @Param remote_dir formData string true "备份目录"
// @Security ApiKeyAuth
// @Router /mysqlBackupDetail/{id}/ [patch]
func UpdateMysqlBackupBaseDetail(c *gin.Context) {
	id := c.Param("id")
	DbName := c.PostForm("db_name")
	DbPort := c.PostForm("db_port")
	RemoteServer := c.PostForm("remote_server")
	RemoteDir := c.PostForm("remote_dir")
	newDetail := new(models.MysqlBackupDetail)
	newDetail.DbName = DbName
	newDetail.DbPort = DbPort
	newDetail.RemoteDir = RemoteDir
	newDetail.RemoteServer = RemoteServer
	Id, err := strconv.Atoi(id)
	if err != nil {
		render.DataError(c, err.Error())
		return
	}
	newDetail.ID = uint(Id)
	err = newDetail.UpdateMysqlBackupBaseDetailByID(hiveview.CONFIG.Db)
	if err != nil {
		utils.LogPrint("err", err)
		render.DataError(c, fmt.Sprintf("更新异常:%s", err))
		return
	}
	render.MSG(c, "已更新")
}
