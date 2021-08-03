package utils

import "hiveview/models"

type AppUpdateLogRes struct {
	ID         int                           `json:"ID"`
	UpdateTime string                        `json:"update_time"`
	AppName    string                        `json:"app_name"`
	Children   []models.ApplicationUpdateLog `json:"children"`
}

//type AppUpdateLogResChild struct {
//	Host       string `json:"host"`
//	UpdateTime string `json:"update_time"`
//	AppName    string `json:"app_name"`
//}

func CombineAppUpdateLog(list []models.ApplicationUpdateLog) (resData []AppUpdateLogRes) {
	var tmpAppUpdateLogResChildList = make(map[string][]models.ApplicationUpdateLog)
	for _, v := range list {
		tmpAppUpdateLogResChildList[v.MD5] = append(tmpAppUpdateLogResChildList[v.MD5], v)
	}
	var id = -1
	for _, v := range tmpAppUpdateLogResChildList {
		var tmpAppUpdateLogRes AppUpdateLogRes
		tmpAppUpdateLogRes.AppName = v[0].AppName
		tmpAppUpdateLogRes.UpdateTime = v[0].UpdateTime
		tmpAppUpdateLogRes.Children = v
		tmpAppUpdateLogRes.ID = id
		resData = append(resData, tmpAppUpdateLogRes)
		id -= 1
	}
	return
}
