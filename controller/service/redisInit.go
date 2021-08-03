package service

import (
	"github.com/gin-gonic/gin"
	"hiveview/render"
	"hiveview/utils"
)

// @title redis初始化
// @version 1.0.0
// @ID redis-init
// @description redis初始化
// @Success 200 {string} string "{"code":200,"msg":"已开始初始化"}"
// @Param ip formData  string true "初始化服务器"
// @Param memsize formData  string true "redis最大内存"
// @Param port formData string true "实例端口号"
// @Param name formData string true "实例名"
// @Tags service
// @Security ApiKeyAuth
// @Router /serviceInit/redis/ [post]
func RedisInit(c *gin.Context) {
	Ip := c.DefaultPostForm("ip", "")

	Port := c.DefaultPostForm("port", "3306")
	Name := c.DefaultPostForm("name", "default")
	MemSize := c.DefaultPostForm("memsize", "1G")

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
		err := utils.ServiceInit("redis", Port, "", Name, Ip, MemSize)
		if err != nil {
			utils.LogPrint("err", err)
			return
		}
	}()

	render.MSG(c, "已开始初始化")
}
