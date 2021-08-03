package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hiveview"
	"hiveview/models"
	"hiveview/render"
	"hiveview/utils"
	"os"
	"os/exec"
	"strings"
	"time"
)

func ResinInit(c *gin.Context) {
	Ip := c.DefaultPostForm("ip", "")
	projectName := c.DefaultPostForm("project_name", "")
	go func() {
		result := models.GetAppDetailByFuzzySearchKey(hiveview.CONFIG.Db, projectName)
		projectPath := result[0].Dir
		projectHost := result[0].Host
		tmp := strings.Split(projectPath, "/")
		date := time.Now().Format("2006-01-02-15-04-05")
		//var copyData = map[string]string{
		//	"hosts":      projectHost,
		//	"date":       date,
		//	"projectDir": projectPath,
		//	"localPath":  tmp[len(tmp)-2],
		//}
		os.Mkdir(fmt.Sprintf("/etc/ansible/playbook_packages/resin/%s", date), 0777)
		cmd := fmt.Sprintf("ansible-playbook -i %s /etc/ansible/playbook/copy_resin_dir.yaml --extra-vars 'hosts=%s' --extra-vars 'date=%s' --extra-vars 'projectDir=%s'  --extra-vars 'localPath=%s'  ", hiveview.CONFIG.Settings.Ansible.Inventory, projectHost, date, projectPath, tmp[len(tmp)-2])
		CMD := exec.Command("bash", "-c", cmd)
		_, err := CMD.Output()
		if err != nil {
			utils.LogPrint("err", err)
			return
		}
		cmd2 := fmt.Sprintf("ansible-playbook -i %s /etc/ansible/playbook/deploy_resin.yaml --extra-vars 'hosts=%s' --extra-vars 'date=%s' --extra-vars 'projectDir=%s'  ", hiveview.CONFIG.Settings.Ansible.Inventory, projectHost, date, tmp[len(tmp)-2])
		CMD2 := exec.Command("bash", "-c", cmd2)
		_, err = CMD2.Output()
		if err != nil {
			utils.LogPrint("err", err)
			return
		}
		//err = utils.AnsiblePlaybook("/etc/ansible/playbook/copy_resin_dir.yaml", copyData)
		//if err != nil {
		//	utils.LogPrint("err", err)
		//	return
		//}

		//var deployData = map[string]string{
		//	"hosts":      Ip,
		//	"date":       date,
		//	"projectDir": tmp[len(tmp)-2],
		//}
		//utils.LogPrint("info", deployData)
		//err = utils.AnsiblePlaybook("/etc/ansible/playbook/deploy_resin.yaml", deployData)
		//if err != nil {
		//	utils.LogPrint("err", err)
		//	return
		//}
		var app = new(models.AppDetail)
		app.Host = Ip
		app.Dir = projectPath
		app.AppName = projectName
		err = app.CreateAppDetail(hiveview.CONFIG.Db)
		if err != nil {
			utils.LogPrint("err", err)
			return
		}
	}()
	render.JSON(c, "已开始创建")

}

func ResinProjectName(c *gin.Context) {
	result := models.GetAllAppName(hiveview.CONFIG.Db)
	render.JSON(c, result)
}
