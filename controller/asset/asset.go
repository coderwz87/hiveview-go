package asset

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hiveview"
	"hiveview/models"
	"hiveview/render"
	"hiveview/utils"
	"os"
	"strconv"
)

type DeleteServerIDs struct {
	ServerIDs []uint `json:"server_ids"`
}

// @title 添加资产信息
// @version 1.0.0
// @ID add-server
// @description 添加资产信息
// @Success 200 {string} string "{"code":200,"msg":"已开始创建"}"
// @Param ip formData  string true "ip"
// @Param idc formData  string false "idc名字"
// @Param comment formData  string false "备注"
// @Param use formData string false "用途"
// @Param cabinet formData string false "机柜号"
// @Param uPosition formData string false "U位"
// @Tags asset
// @Security ApiKeyAuth
// @Router /addServer/ [post]
func AddServer(c *gin.Context) {
	idc := c.DefaultPostForm("idc", "北京")
	comment := c.DefaultPostForm("comment", "")
	ip := c.DefaultPostForm("ip", "")
	use := c.DefaultPostForm("use", "其他")
	cabinet := c.DefaultPostForm("cabinet", "")
	uPosition := c.DefaultPostForm("u_position", "")
	if ip == "" {
		render.DataError(c, "no ip input")
		return
	}
	ifIP := utils.JudgeType(ip)
	if !ifIP {
		render.DataError(c, "input is not ip")
		return
	}

	ifAlive := utils.JudgeAlive(ip)

	if !ifAlive {
		render.DataError(c, "ip is not alive")
		return
	}

	ifExist := new(models.Assets)
	ifExist.IP = ip

	if ifExist.GetAssetByIP(hiveview.CONFIG.Db) {
		render.DataError(c, fmt.Sprintf("资产已存在"))
		return
	}
	ifExist.IDC = idc
	ifExist.Comment = comment
	ifExist.Cabinet = cabinet
	ifExist.UPosition = uPosition
	ifExist.Use = use
	err := ifExist.CreateAsset(hiveview.CONFIG.Db)
	if err != nil {
		render.DataError(c, fmt.Sprintf("资产创建失败"))
	}
	go func() {
		result := new(utils.Result)

		data, err := result.GetAssetInfo(fmt.Sprintf("%s%s", hiveview.CONFIG.Settings.Ansible.Factsdir, ip), ip, ifExist)
		if err != nil {
			utils.LogPrint("err", err)
			return
		}
		err = data.UpdateAssetInfo(hiveview.CONFIG.Db)
		if err != nil {
			utils.LogPrint("err", err)
			return
		}
	}()

	render.MSG(c, "已开始创建")
}

// @title 获取单台资产信息
// @version 1.0.0
// @ID get-server
// @description 获取单台资产信息
// @Success 200 {string} string "{"code":200,"msg":"success","data":{object}}"
// @Param id path  string true "id"
// @Tags asset
// @Security ApiKeyAuth
// @Router /server/{id}/ [get]
func GetServer(c *gin.Context) {
	id := c.Param("id")
	result := new(models.Assets)
	ID, err := strconv.Atoi(id)
	if err != nil {
		render.DataError(c, err.Error())
		return
	}
	result.ID = uint(ID)
	err = result.GetAssetByID(hiveview.CONFIG.Db)
	if err != nil {
		render.DataError(c, "查询失败")
		return
	}

	render.JSON(c, result)

}

// @title 更新资产信息
// @version 1.0.0
// @ID update-server
// @description 更新资产信息
// @Success 200 {string} string "{"code":200,"msg":"更新成功"}"
// @Param id path  string true "id"
// @Param idc formData  string false "idc名字"
// @Param comment formData  string false "备注"
// @Param use formData string false "用途"
// @Param cabinet formData string false "机柜号"
// @Param uPosition formData string false "U位"
// @Tags asset
// @Security ApiKeyAuth
// @Router /server/{id}/ [patch]
func UpdateServer(c *gin.Context) {
	id := c.Param("id")
	idc := c.DefaultPostForm("idc", "")
	cabinet := c.DefaultPostForm("cabinet", "")
	uPosition := c.DefaultPostForm("u_position", "")
	use := c.DefaultPostForm("use", "其他")
	comment := c.DefaultPostForm("comment", "")
	Id, err := strconv.Atoi(id)
	if err != nil {
		render.DataError(c, err.Error())
		return
	}
	UpdateData := &models.Assets{
		IDC:       idc,
		Cabinet:   cabinet,
		UPosition: uPosition,
		Use:       use,
		Comment:   comment,
	}
	UpdateData.ID = uint(Id)
	utils.LogPrint("info", UpdateData)
	err = UpdateData.UpdateAssetByID(hiveview.CONFIG.Db)
	if err != nil {
		render.DataError(c, "更新失败")
		return
	}
	render.MSG(c, "更新成功")

}

// @title 删除资产信息
// @version 1.0.0
// @ID delete-server
// @description 删除资产信息
// @Success 200 {string} string "{"code":200,"msg":"success"}"
// @Param id path  string true "id"
// @Tags asset
// @Security ApiKeyAuth
// @Router /server/{id}/ [delete]
func DeleteServer(c *gin.Context) {
	id := c.Param("id")
	result := new(models.Assets)
	ID, err := strconv.Atoi(id)
	if err != nil {
		render.DataError(c, err.Error())
		return
	}
	result.ID = uint(ID)
	err = result.GetAssetByID(hiveview.CONFIG.Db)
	if err != nil {
		render.DataError(c, fmt.Sprintf("删除异常:%s", err))
		return
	}
	err = utils.DeleteInfoFromConsul(*result, "node-exporter")
	if err != nil {
		render.DataError(c, fmt.Sprintf("删除prometheus异常:%s", err))
		return
	}
	//err = utils.ZabbixDeleteHost(result)
	//if err != nil {
	//	render.DataError(c, fmt.Sprintf("删除zabbix异常:%s", err))
	//	return
	//}
	os.Remove(fmt.Sprintf("%s%s", hiveview.CONFIG.Settings.Ansible.Factsdir, result.IP))
	err = result.DeleteAssetByID(hiveview.CONFIG.Db)
	if err != nil {
		render.DataError(c, fmt.Sprintf("删除异常:%s", err))
		return
	}

	render.Success(c)

}

// @title 批量删除资产信息
// @version 1.0.0
// @ID delete-all-server
// @description 批量删除资产信息
// @Success 200 {string} string "{"code":200,"msg":"已全部删除"}"
// @Param id formData  DeleteServerIDs true "ids"
// @Tags asset
// @Security ApiKeyAuth
// @Router /deleteAllServer/ [delete]
func DeleteAllServer(c *gin.Context) {
	var deleteIDs DeleteServerIDs
	err := c.BindJSON(&deleteIDs)
	if err != nil {
		utils.LogPrint("info", err)
		render.DataError(c, fmt.Sprintf("%s", err))
		return
	}
	for _, v := range deleteIDs.ServerIDs {
		result := new(models.Assets)
		result.ID = v
		err = result.GetAssetByID(hiveview.CONFIG.Db)
		if err != nil {
			render.DataError(c, fmt.Sprintf("删除异常:%s", err))
			return
		}
		os.Remove(fmt.Sprintf("%s%s", hiveview.CONFIG.Settings.Ansible.Factsdir, result.IP))
		err = utils.DeleteInfoFromConsul(*result, "node-exporter")
		if err != nil {
			render.DataError(c, fmt.Sprintf("删除异常:%s", err))
			return
		}
		//err = utils.ZabbixDeleteHost(result)
		//if err != nil {
		//	render.DataError(c, fmt.Sprintf("删除zabbix异常:%s", err))
		//	return
		//}
		result.DeleteAssetByID(hiveview.CONFIG.Db)
	}
	render.MSG(c, "已全部删除")
}

// @title 获取全部资产信息
// @version 1.0.0
// @ID get-all-server
// @description 获取全部资产信息
// @Success 200 {string} string "{"code":200,"msg":"success","data":[...]}"
// @Tags asset
// @Security ApiKeyAuth
// @Router /getAllServer/ [get]
func GetAllServer(c *gin.Context) {
	data, err := models.GetAllAsset(hiveview.CONFIG.Db)
	if err != nil {
		render.DataError(c, fmt.Sprintf("查询异常:%s", err))
		return
	}
	render.JSON(c, data)
}

// @title 搜索资产信息
// @version 1.0.0
// @ID search-server-info
// @description 搜索资产信息
// @Success 200 {string} string "{"code":200,"msg":"success","data":[...]}"
// @Tags asset
// @Security ApiKeyAuth
func SearchServerInfo(c *gin.Context) {
	ip := c.Query("search")
	resultList := models.GetAssetByFuzzyIP(hiveview.CONFIG.Db, ip)
	render.JSON(c, resultList)
}

// @title 获取资产IDC信息
// @version 1.0.0
// @ID get-server-idc
// @description 获取资产IDC信息
// @Success 200 {string} string "{"code":200,"msg":"success","data":[...]}"
// @Tags asset
// @Security ApiKeyAuth
func GetServerIdc(c *gin.Context) {
	idc := models.GetAssetIdc(hiveview.CONFIG.Db)
	render.JSON(c, idc)
}
