package models

import "gorm.io/gorm"

type BatchLog struct {
	Host    string `json:"host" gorm:"type:varchar(128);`
	Type    string `json:"type" gorm:"type:varchar(128);`
	LogFile string `json:"log_file" gorm:"type:varchar(128);`
	Detail  string `json:"detail" gorm:"type:varchar(256);`
	gorm.Model
}

func (u *BatchLog) TableName() string {
	return "BatchLog"
}

func (u *BatchLog) CreateBatchLog(db *gorm.DB) (err error) {
	result := db.Create(u)
	if result.Error != nil {
		return result.Error
	}
	return
}

func GetAllBatchLog(db *gorm.DB) (resultList []BatchLog, err error) {

	r := db.Find(&resultList)
	return resultList, r.Error
}

func (u *BatchLog) GetBatchLogByID(db *gorm.DB) (err error) {
	result := db.Where("id=?", u.ID).First(u)
	if result.Error != nil {
		return result.Error
	}
	return
}
