package utils

import (
	"hiveview"
	"hiveview/models"
	"strconv"
)

func AppOperation(id, action string) (err error) {
	result := new(models.OperationDetail)
	ID, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	result.ID = uint(ID)
	err = result.GetOperationDetailByID(hiveview.CONFIG.Db)
	playbookName := "/etc/ansible/playbook/operation_app.yaml"
	data := make(map[string]string)
	data["hosts"] = result.Host
	data["AppName"] = result.AppName
	data["Dir"] = result.Dir
	data["Type"] = result.Type
	data["action"] = action
	go func() {
		err = AnsiblePlaybook(playbookName, data)
		if err != nil {
			LogPrint("err", err)
			return
		}
		result.State = "已执行"
		_ = result.UpdateOperationStateByID(hiveview.CONFIG.Db)
	}()
	return
}
