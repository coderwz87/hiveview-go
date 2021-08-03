package models

import (
	"fmt"
	"gorm.io/gorm"
)

type AppDetail struct {
	AppName string `json:"app_name" gorm:"type:varchar(128);"` //应用名
	Host    string `json:"host" gorm:"type:varchar(128);"`     //应用所在服务器
	Dir     string `json:"dir" gorm:"type:varchar(128);"`      //应用所在目录
	Type    string `json:"type" gorm:"type:varchar(128);"`     //应用容器类型
	State   string `json:"state" gorm:"type:varchar(128);"`    //应用状态
	gorm.Model
}

func (u *AppDetail) TableName() string {
	return "AppDetail"
}

func GetAllAppDetail(db *gorm.DB) (resultList []AppDetail, err error) {

	r := db.Find(&resultList)
	return resultList, r.Error
}

func GetAllAppName(db *gorm.DB) []string {
	var resultList []AppDetail
	var result []string
	db.Distinct("app_name").Find(&resultList)
	for _, v := range resultList {
		result = append(result, v.AppName)
	}
	return result
}

func (u *AppDetail) DeleteAppDetailByID(db *gorm.DB) (err error) {
	err = u.GetAppDetailByID(db)
	if err != nil {
		err = fmt.Errorf("不存在此记录")
		return
	}
	result := db.Delete(u, u.ID)
	if result.Error != nil {
		return result.Error
	}
	return
}

func (u *AppDetail) GetAppDetailByID(db *gorm.DB) (err error) {
	result := db.Where("id=?", u.ID).First(u)
	if result.Error != nil {
		return result.Error
	}
	return
}

func (u *AppDetail) CreateAppDetail(db *gorm.DB) (err error) {
	result := db.Create(u)
	if result.Error != nil {
		return result.Error
	}
	return
}

func (u *AppDetail) UpdateAppDetailState(db *gorm.DB) (err error) {
	result := db.Model(u).Where("app_name = ? and host = ? and dir = ? and type = ?", u.AppName, u.Host, u.Dir, u.Type).Update("state", u.State)
	if result.Error != nil {
		return result.Error
	}
	return
}

func GetAppDetailByFuzzySearchKey(db *gorm.DB, key string) []AppDetail {
	var result []AppDetail
	db.Where("app_name like ? or host like ?", fmt.Sprintf("%%%s%%", key), fmt.Sprintf("%%%s%%", key)).Find(&result)
	return result
}

func (u *AppDetail) UpdateAppBaseDetail(db *gorm.DB) (err error) {
	result := db.Model(u).Updates(map[string]interface{}{"host": u.Host, "dir": u.Dir, "type": u.Type})
	if result.Error != nil {
		return result.Error
	}
	return
}
