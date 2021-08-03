package service

import (
	"github.com/gin-gonic/gin"
	"hiveview/render"
	"hiveview/utils"
)

// @title mysql初始化
// @version 1.0.0
// @ID mysql-init
// @description mysql初始化
// @Success 200 {string} string "{"code":200,"msg":"已开始初始化"}"
// @Param ip formData  string true "初始化服务器"
// @Param version formData  string true "数据库版本"
// @Param port formData string true "实例端口号"
// @Param name formData string true "实例名"
// @Tags service
// @Security ApiKeyAuth
// @Router /serviceInit/mysql/ [post]
func MysqlInit(c *gin.Context) {
	Ip := c.DefaultPostForm("ip", "")
	MysqlVersion := c.DefaultPostForm("version", "5.6.23")
	Port := c.DefaultPostForm("port", "3306")
	Name := c.DefaultPostForm("name", "default")

	if len(Ip) == 0 {
		render.DataError(c, "not input ip")
		return
	}

	if !utils.JudgeType(Ip) {
		render.DataError(c, "input is not ip")
		return
	}
	if !utils.JudgeInMysql(Ip) {
		render.DataError(c, "请先添加资产")
		return
	}
	go func() {
		err := utils.ServiceInit("mysql", Port, MysqlVersion, Name, Ip, "")
		if err != nil {
			utils.LogPrint("err", err)
			return
		}
	}()

	render.MSG(c, "已开始初始化")
}
