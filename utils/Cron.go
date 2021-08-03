package utils

import (
	"fmt"
	"github.com/robfig/cron"
	"hiveview"
	"hiveview/models"
	"regexp"
	"sync"
	"time"
)

//每天检测mysql备份情况，并更新到数据库
func MysqlBackupCronCheck() {
	data, err := models.GetAllMysqlBackupDetail(hiveview.CONFIG.Db)
	if err != nil {
		LogPrint("err", err)
		return
	}
	for _, v := range data {
		CheckMysqlBackup(v.DbName, v.DbPort, v.RemoteServer, v.RemoteDir)
	}
}

//核实mysql备份情况，异常情况发送钉钉告警
func MysqlBackupDetailCheck() {
	data, err := models.GetAllMysqlBackupDetail(hiveview.CONFIG.Db)
	if err != nil {
		LogPrint("err", err)
		return
	}
	re := regexp.MustCompile("not")

	for _, v := range data {
		match := re.FindString(v.BackupLog)
		if len(match) != 0 {
			msg := fmt.Sprintf("数据库备份有误：%s_%s,请检查", v.DbName, v.DbPort)
			err := SendMsgToDD(msg)
			if err != nil {
				LogPrint("err", err)
			}
			time.Sleep(10 * time.Second)
		}
	}
}

//每天检测其他项目备份情况，并更新到数据库
func OtherBackupCronCheck() {
	data, err := models.GetAllOtherBackupDetail(hiveview.CONFIG.Db)
	if err != nil {
		LogPrint("err", err)
		return
	}
	for _, v := range data {
		CheckOtherBackupDetail(v.ProjectName, v.RemoteServer, v.RemoteDir)
	}
}

//核实其他项目备份情况，异常情况发送钉钉告警
func OtherBackupDetailCheck() {
	data, err := models.GetAllOtherBackupDetail(hiveview.CONFIG.Db)
	if err != nil {
		LogPrint("err", err)
		return
	}
	re := regexp.MustCompile("not")

	for _, v := range data {
		match := re.FindString(v.BackupLog)
		LogPrint("info", match)
		if len(match) != 0 {
			msg := fmt.Sprintf("项目备份有误：%s,请检查", v.ProjectName)
			err := SendMsgToDD(msg)
			if err != nil {
				LogPrint("err", err)
			}
			time.Sleep(10 * time.Second)
		}
	}
}

//检查app运行状态
func CheckAppState() {
	data, err := models.GetAllAppDetail(hiveview.CONFIG.Db)
	if err != nil {
		LogPrint("err", err)
		return
	}
	ch := make(chan bool, 10)
	var wg = sync.WaitGroup{}
	playbook := "/etc/ansible/playbook/check_app_status.yaml"
	for _, v := range data {
		ch <- true
		wg.Add(1)
		go func(v models.AppDetail) {
			tmp := make(map[string]string)
			tmp["hosts"] = v.Host
			tmp["AppName"] = v.AppName
			tmp["Dir"] = v.Dir
			err := AnsiblePlaybook(playbook, tmp)
			//cmd := fmt.Sprintf("ansible-playbook -i %s /etc/ansible/playbook/check_app_status.yaml --extra-vars 'hosts=%s' --extra-vars 'AppName=%s'  --extra-vars 'Dir=%s' --extra-vars 'Type=resin' ", hiveview.CONFIG.Settings.Ansible.Inventory, v.Host, v.AppName, v.Dir)
			//CMD := exec.Command("bash", "-c", cmd)
			//_, err := CMD.Output()
			if err != nil {
				LogPrint("err", err)
			}
			<-ch
			wg.Done()
		}(v)
	}
	wg.Wait()
}

//更新tag信息到consul
func UpdateNodeExporterToConsul() {
	resultList, err := models.GetBeiJingAsset(hiveview.CONFIG.Db)
	if err != nil {
		LogPrint("err", err)
		return
	}
	for _, v := range resultList {
		err := PutInfoToConsul(&v, "node-exporter")
		if err != nil {
			LogPrint("err", err)
		}
	}
}

//计划任务注册
func InitCronJob(c *cron.Cron) {
	c.AddFunc("0 0 8 * * *", MysqlBackupCronCheck)
	c.AddFunc("0 15 9 * * *", MysqlBackupDetailCheck)
	c.AddFunc("0 0 8 * * *", OtherBackupCronCheck)
	c.AddFunc("0 15 9 * * *", OtherBackupDetailCheck)
	c.AddFunc("0 0 */1 * * *", CheckAppState)
	c.AddFunc("0 5 */1 * * *", UpdateNodeExporterToConsul)
	c.Start()
}
