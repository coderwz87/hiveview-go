package utils

import (
	"fmt"
	"time"
)

//用ansible-playbook 检查其他项目备份情况，并写库
func CheckOtherBackupDetail(ProjectName, RemoteServer, RemoteDir string) {
	now := time.Now()
	DAY := fmt.Sprintf(now.Format("2006-01-02"))
	playbook := "/etc/ansible/playbook/check_other_backup.yaml"
	data := make(map[string]string)
	data["hosts"] = RemoteServer
	data["ProjectName"] = ProjectName
	data["RemoteDir"] = RemoteDir
	data["DAY"] = DAY
	err := AnsiblePlaybook(playbook, data)
	//cmd := fmt.Sprintf("ansible-playbook -i %s /etc/ansible/playbook/check_other_backup.yaml --extra-vars 'hosts=%s' --extra-vars 'ProjectName=%s'  --extra-vars 'RemoteDir=%s' --extra-vars 'DAY=%s'", hiveview.CONFIG.Settings.Ansible.Inventory, RemoteServer, ProjectName, RemoteDir, DAY)
	//CMD := exec.Command("bash", "-c", cmd)
	//_, err := CMD.Output()
	if err != nil {
		LogPrint("err", err)
	}
}
