package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hiveview"
	"hiveview/models"
	"hiveview/render"
	"hiveview/utils"
)

type BatchServiceRequest struct {
	ServiceName string   `json:"service_name"`
	InitIps     []string `json:"init_ips"`
}

// @title 批量服务初始化初始化
// @version 1.0.0
// @ID batch-service-init
// @description 批量服务初始化初始化
// @Success 200 {string} string "{"code":200,"msg":"已开始初始化"}"
// @Param service_name formData  string true "初始化服务名"
// @Param init_ips formData  []string true "初始化服务器"
// @Tags service
// @Security ApiKeyAuth
// @Router /serviceInit/batch/ [put]
func BatchServiceInit(c *gin.Context) {
	var data BatchServiceRequest
	err := c.BindJSON(&data)
	if err != nil {
		utils.LogPrint("err", err)
		render.DataError(c, fmt.Sprintf("%s", err))
		return
	}
	go serviceInit(data.ServiceName, data.InitIps)
	render.MSG(c, "已开始初始化")
}

func serviceInit(serviceName string, ips []string) {
	if serviceName == "node_exporter" {
		for _, v := range ips {
			NodeExporterInit(v)
		}
	}

}

func NodeExporterInit(ip string) {
	var data = map[string]string{
		"hosts": ip,
	}
	err := utils.AnsiblePlaybook("/etc/ansible/playbook/node_exporter_install.yaml", data)
	if err != nil {
		utils.LogPrint("err", err)
	}
	var ipInfo = new(models.Assets)
	ipInfo.IP = ip
	ipInfo.GetAssetByIP(hiveview.CONFIG.Db)
	err = utils.PutInfoToConsul(ipInfo, "node-exporter")
	if err != nil {
		utils.LogPrint("err", err)
	}
}
