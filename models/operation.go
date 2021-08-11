package models

import "gorm.io/gorm"

type OperationDetail struct {
	AppName string `json:"app_name" gorm:"type:varchar(128);"` //应用名
	Host    string `json:"host" gorm:"type:varchar(128);"`     //应用所在服务器
	Dir     string `json:"dir" gorm:"type:varchar(128);"`      //应用所在目录
	Type    string `json:"type" gorm:"type:varchar(128);"`     //应用容器类型
	Action  string `json:"action" gorm:"type:varchar(128);"`   //操作
	State   string `json:"state" gorm:"type:varchar(128);"`    //操作状态
	User    string `json:"user" gorm:"type:varchar(128);`      //操作人
	gorm.Model
}

func (u *OperationDetail) TableName() string {
	return "OperationDetail"
}

func (u *OperationDetail) CreateOperationDetail(db *gorm.DB) (err error) {
	result := db.Create(u)
	if result.Error != nil {
		return result.Error
	}
	return
}

func (u *OperationDetail) GetOperationDetailByID(db *gorm.DB) (err error) {
	result := db.Where("id=?", u.ID).First(u)
	if result.Error != nil {
		return result.Error
	}
	return
}

func (u *OperationDetail) UpdateOperationStateByID(db *gorm.DB) (err error) {
	result := db.Model(u).Where("id = ?", u.ID).Update("state", u.State)
	if result.Error != nil {
		return result.Error
	}
	return
}

func GetAllOperationDetail(db *gorm.DB) (resultList []OperationDetail, err error) {

	r := db.Find(&resultList)
	return resultList, r.Error
}
