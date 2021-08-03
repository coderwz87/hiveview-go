package utils

import (
	"fmt"
	"time"
)

//用ansible-playbook 检查mysql备份情况，并写库
func CheckMysqlBackup(DbName, DbPort, RemoteServer, RemoteDir string) {
	now := time.Now()
	DAY := fmt.Sprintf(now.Format("2006-01-02"))
	playbookName := "/etc/ansible/playbook/check_mysql_backup.yaml"
	data := make(map[string]string)
	data["hosts"] = RemoteServer
	data["DbName"] = DbName
	data["DbPort"] = DbPort
	data["RemoteDir"] = RemoteDir
	data["DAY"] = DAY
	err := AnsiblePlaybook(playbookName, data)
	if err != nil {
		LogPrint("err", err)
	}

}
