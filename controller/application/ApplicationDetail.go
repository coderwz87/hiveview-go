package application

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

// @title 获取全部应用信息
// @version 1.0.0
// @ID get-all-application-detail
// @description 获取全部应用信息
// @Success 200 {string} string "{"code":200,"msg":"success","data":[...]}"
// @Tags Application
// @Security ApiKeyAuth
// @Router /getAllAppDetail/ [get]
func GetAllApplicationDetail(c *gin.Context) {
	data, err := models.GetAllAppDetail(hiveview.CONFIG.Db)
	if err != nil {
		render.DataError(c, fmt.Sprintf("查询异常:%s", err))
		return
	}
	render.JSON(c, data)
}

// @title 删除单条信息
// @version 1.0.0
// @ID delete-application-detail
// @description 删除单条信息
// @Success 200 {string} string "{"code":200,"msg":"success"}"
// @Tags Application
// @Security ApiKeyAuth
// @Param id path int  true "应用id"
// @Router /appDetail/{id}/ [delete]
func DeleteApplicationDetail(c *gin.Context) {
	id := c.Param("id")
	result := new(models.AppDetail)
	ID, err := strconv.Atoi(id)
	if err != nil {
		render.DataError(c, err.Error())
		return
	}
	result.ID = uint(ID)
	err = result.DeleteAppDetailByID(hiveview.CONFIG.Db)
	if err != nil {
		render.DataError(c, fmt.Sprintf("删除异常:%s", err))
		return
	}
	render.Success(c)

}

// @title 添加应用信息
// @version 1.0.0
// @ID add-application-detail
// @description 添加应用信息
// @Success 200 {string} string "{"code":200,"msg":"创建成功"}"
// @Param app_name formData  string true "app名字"
// @Param host formData  string true "主机名"
// @Param dir formData  string true "目录"
// @Param type formData string true "类型"
// @Tags Application
// @Security ApiKeyAuth
// @Router /addAppDetail/ [post]
func AddApplicationDetail(c *gin.Context) {
	AppName := c.PostForm("app_name")
	Host := c.PostForm("host")
	Dir := c.PostForm("dir")
	Type := c.PostForm("type")
	if len(AppName) == 0 || len(Host) == 0 || len(Dir) == 0 || len(Type) == 0 {
		render.DataError(c, "提交数据有误")
		return
	}
	ifIP := utils.JudgeType(Host)
	if !ifIP {
		render.DataError(c, "服务器地址不是一个IP")
		return
	}
	ifAlive := utils.JudgeAlive(Host)

	if !ifAlive {
		render.DataError(c, "服务器不通")
		return
	}
	newAppDetail := new(models.AppDetail)
	newAppDetail.AppName = AppName
	newAppDetail.Host = Host
	newAppDetail.Dir = Dir
	newAppDetail.Type = Type
	err := newAppDetail.CreateAppDetail(hiveview.CONFIG.Db)
	if err != nil {
		render.DataError(c, "创建失败")
		return
	}
	render.MSG(c, "创建成功")
}

// @title 更新应用状态
// @version 1.0.0
// @ID update-application-detail
// @description 更新应用状态
// @Success 200 {string} string "{"code":200,"msg":"已更新"}"
// @Param app_name formData  string true "app名字"
// @Param host formData  string true "主机名"
// @Param dir formData  string true "目录"
// @Param type formData  string true "类型"
// @Param state formData  string true "状态"
// @Tags Application
// @Security ApiKeyAuth
// @Router /updateAppDetailState/ [put]
func UpdateApplicationState(c *gin.Context) {
	AppName := c.PostForm("app_name")
	Host := c.PostForm("host")
	Dir := c.PostForm("dir")
	Type := c.PostForm("type")
	State := c.PostForm("state")
	if len(AppName) == 0 || len(Host) == 0 || len(Dir) == 0 || len(Type) == 0 {
		render.DataError(c, "提交数据有误")
		return
	}
	ifIP := utils.JudgeType(Host)
	if !ifIP {
		render.DataError(c, "服务器地址不是一个IP")
		return
	}
	ifAlive := utils.JudgeAlive(Host)

	if !ifAlive {
		render.DataError(c, "服务器不通")
		return
	}
	newAppDetail := new(models.AppDetail)
	newAppDetail.AppName = AppName
	newAppDetail.Host = Host
	newAppDetail.Dir = Dir
	newAppDetail.Type = Type
	newAppDetail.State = State
	err := newAppDetail.UpdateAppDetailState(hiveview.CONFIG.Db)
	if err != nil {
		utils.LogPrint("err", err)
		return
	}

	render.MSG(c, "已更新")

}

// @title 批量删除应用信息
// @version 1.0.0
// @ID delete-all-application-detail
// @description 批量删除应用信息
// @Success 200 {string} string "{"code":200,"msg":"已全部删除"}"
// @Param detail_ids formData DeleteDetailIDs true "删除的id号"
// @Tags Application
// @Security ApiKeyAuth
// @Router /deleteAllAppDetail/ [delete]
func DeleteAllAppDetail(c *gin.Context) {
	var deleteIDs DeleteDetailIDs
	err := c.BindJSON(&deleteIDs)
	if err != nil {
		utils.LogPrint("err", err)
		render.DataError(c, fmt.Sprintf("%s", err))
		return
	}
	for _, v := range deleteIDs.DetailIDs {
		result := new(models.AppDetail)
		result.ID = v
		err = result.DeleteAppDetailByID(hiveview.CONFIG.Db)
		if err != nil {
			utils.LogPrint("err", err)
			render.DataError(c, fmt.Sprintf("%s", err))
			return
		}
	}
	render.MSG(c, "已全部删除")
}

// @title 操作应用
// @version 1.0.0
// @ID operation-application
// @description 操作应用
// @Success 200 {string} string "{"code":200,"msg":"已开始操作"}"
// @Param id formData string true "id"
// @Param action formData string  true "操作动作"
// @Tags Application
// @Security ApiKeyAuth
// @Router /operationApp/ [post]
func OperationApp(c *gin.Context) {
	appID := c.Query("appID")
	action := c.Query("action")
	claims, _ := c.Get("claims")
	username := claims.(*utils.CustomClaims).Username
	appDetail := new(models.AppDetail)
	ID, err := strconv.Atoi(appID)
	if err != nil {
		render.DataError(c, err.Error())
		return
	}
	appDetail.ID = uint(ID)
	err = appDetail.GetAppDetailByID(hiveview.CONFIG.Db)
	if err != nil {
		utils.LogPrint("err", err)
		render.DataError(c, err.Error())
		return
	}
	if username == "admin" {
		playbookName := "/etc/ansible/playbook/operation_app.yaml"
		data := make(map[string]string)
		data["hosts"] = appDetail.Host
		data["AppName"] = appDetail.AppName
		data["Dir"] = appDetail.Dir
		data["Type"] = appDetail.Type
		data["action"] = action
		go func() {
			err = utils.AnsiblePlaybook(playbookName, data)
			if err != nil {
				utils.LogPrint("err", err)
			}
		}()
		render.MSG(c, "已开始操作")
		return
	}
	var operationDetail = new(models.OperationDetail)
	operationDetail.Host = appDetail.Host
	operationDetail.AppName = appDetail.AppName
	operationDetail.Dir = appDetail.Dir
	operationDetail.Type = appDetail.Type
	operationDetail.Action = action
	operationDetail.State = "确认中"
	operationDetail.User = username
	err = operationDetail.CreateOperationDetail(hiveview.CONFIG.Db)
	if err != nil {
		utils.LogPrint("err", err)
		render.DataError(c, err.Error())
		return
	}
	dingMSG := fmt.Sprintf("http://211.103.138.124:18080/OperationDetail/%d/", operationDetail.ID)
	err = utils.SendMsgToDD(dingMSG)
	if err != nil {
		utils.LogPrint("err", err)
		render.DataError(c, err.Error())
		return
	}
	render.MSG(c, "已发送确认信息")
}

//func OperationApp(c *gin.Context) {
//	id := c.PostForm("id")
//	Action := c.PostForm("action")
//	result := new(models.AppDetail)
//	ID, err := strconv.Atoi(id)
//	if err != nil {
//		render.DataError(c, err.Error())
//		return
//	}
//	result.ID = uint(ID)
//	err = result.GetAppDetailByID(hiveview.CONFIG.Db)
//	playbookName := "/etc/ansible/playbook/operation_app.yaml"
//	data := make(map[string]string)
//	data["hosts"] = result.Host
//	data["AppName"] = result.AppName
//	data["Dir"] = result.Dir
//	data["Type"] = result.Type
//	data["action"] = Action
//	go func() {
//		err = utils.AnsiblePlaybook(playbookName, data)
//		if err != nil {
//			utils.LogPrint("err", err)
//		}
//	}()
//	render.MSG(c, "已开始操作")
//}

// @title 搜索应用信息
// @version 1.0.0
// @ID search-application-detail
// @description 搜索应用信息
// @Success 200 {string} string "{"code":200,"msg":"success","data":[...]}"
// @Tags Application
// @Security ApiKeyAuth
// @Router /searchAppDetail/ [get]
func SearchApplicationDetail(c *gin.Context) {
	searchKey := c.Query("search")
	resultList := models.GetAppDetailByFuzzySearchKey(hiveview.CONFIG.Db, searchKey)
	render.JSON(c, resultList)
}

// @title 获取单个app信息
// @version 1.0.0
// @ID get-application-detail
// @description 获取单个app信息
// @Success 200 {string} string "{"code":200,"msg":"success","data":{object}}"
// @Param id path  string true "id"
// @Tags Application
// @Security ApiKeyAuth
// @Router /AppDetail/{id}/ [get]
func GetApplicationDetail(c *gin.Context) {
	id := c.Param("id")
	result := new(models.AppDetail)
	ID, err := strconv.Atoi(id)
	if err != nil {
		render.DataError(c, err.Error())
		return
	}
	result.ID = uint(ID)
	err = result.GetAppDetailByID(hiveview.CONFIG.Db)
	if err != nil {
		utils.LogPrint("err", err)
		render.DataError(c, err.Error())
		return
	}
	render.JSON(c, result)
}

func UpdateApplicationBaseDetail(c *gin.Context) {
	id := c.Param("id")
	AppName := c.PostForm("app_name")
	Host := c.PostForm("host")
	Dir := c.PostForm("dir")
	Type := c.PostForm("type")
	if len(AppName) == 0 || len(Host) == 0 || len(Dir) == 0 || len(Type) == 0 {
		render.DataError(c, "提交数据有误")
		return
	}
	ifIP := utils.JudgeType(Host)
	if !ifIP {
		render.DataError(c, "服务器地址不是一个IP")
		return
	}
	ifAlive := utils.JudgeAlive(Host)

	if !ifAlive {
		render.DataError(c, "服务器不通")
		return
	}
	newAppDetail := new(models.AppDetail)
	newAppDetail.AppName = AppName
	newAppDetail.Host = Host
	newAppDetail.Dir = Dir
	newAppDetail.Type = Type
	Id, err := strconv.Atoi(id)
	if err != nil {
		render.DataError(c, err.Error())
		return
	}
	newAppDetail.ID = uint(Id)
	err = newAppDetail.UpdateAppBaseDetail(hiveview.CONFIG.Db)
	if err != nil {
		utils.LogPrint("err", err)
		render.DataError(c, fmt.Sprintf("更新异常:%s", err))
		return
	}
	render.MSG(c, "已更新")
}
