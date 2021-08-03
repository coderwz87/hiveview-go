package models

import "gorm.io/gorm"
import "fmt"

type ApplicationUpdateLog struct {
	MD5        string `json:"md5" gorm:"type:varchar(128);"`
	AppName    string `json:"app_name" gorm:"type:varchar(128);"`
	Host       string `json:"host" gorm:"type:varchar(128);"`
	UpdateTime string `json:"update_time" gorm:"type:varchar(128);"`
	gorm.Model
}

func (u *ApplicationUpdateLog) TableName() string {
	return "AppUpdateLog"
}

func GetAllAppUpdateLog(db *gorm.DB) (resultList []ApplicationUpdateLog, err error) {

	r := db.Find(&resultList)
	return resultList, r.Error
}

func (u *ApplicationUpdateLog) CreateAppUpdateLog(db *gorm.DB) (err error) {
	result := db.Create(u)
	if result.Error != nil {
		return result.Error
	}
	return
}

func GetAppUpdateLogByFuzzyAppName(db *gorm.DB, key string) []ApplicationUpdateLog {
	var result []ApplicationUpdateLog
	db.Where("app_name like ?", fmt.Sprintf("%%%s%%", key)).Find(&result)
	return result
}

func GetAppUpdateCount(db *gorm.DB, dateList []string) []int {
	var result []int
	for _, v := range dateList {
		var tmp []ApplicationUpdateLog
		db.Where("update_time like ?", fmt.Sprintf("%%%s%%", v)).Find(&tmp)
		result = append(result, len(tmp))
	}
	return result
}
