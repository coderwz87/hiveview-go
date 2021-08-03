package models

import (
	"fmt"
	"gorm.io/gorm"
)

type MysqlBackupDetail struct {
	DbName       string `json:"db_name" gorm:"type:varchar(128);"`
	DbPort       string `json:"db_port" gorm:"type:varchar(128);"`
	RemoteServer string `json:"remote_server" gorm:"type:varchar(128);"`
	RemoteDir    string `json:"remote_dir" gorm:"type:varchar(128);"`
	BackupLog    string `json:"backup_log" gorm:"type:varchar(256);"`
	gorm.Model
}

func (u *MysqlBackupDetail) TableName() string {
	return "MysqlBackupDetail"
}

func (u *MysqlBackupDetail) CreateMysqlBackupDetail(db *gorm.DB) (err error) {
	result := db.Create(u)
	if result.Error != nil {
		return result.Error
	}
	return
}

func GetAllMysqlBackupDetail(db *gorm.DB) (resultList []MysqlBackupDetail, err error) {

	r := db.Find(&resultList)
	return resultList, r.Error
}

func (u *MysqlBackupDetail) GetMysqlBackupDetailByID(db *gorm.DB) (err error) {
	result := db.Where("id=?", u.ID).First(u)
	if result.Error != nil {
		return result.Error
	}
	return
}

func (u *MysqlBackupDetail) DeleteMysqlBackupDetailByID(db *gorm.DB) (err error) {
	err = u.GetMysqlBackupDetailByID(db)
	if err != nil {
		err = fmt.Errorf("不存在此资产记录")
		return
	}
	result := db.Delete(u, u.ID)
	if result.Error != nil {
		return result.Error
	}
	return
}

func (u *MysqlBackupDetail) UpdateMysqlBackupDetail(db *gorm.DB) (err error) {
	result := db.Model(u).Where("db_name = ? and db_port = ? and remote_server = ? and remote_dir = ?", u.DbName, u.DbPort, u.RemoteServer, u.RemoteDir).Update("backup_log", u.BackupLog)
	if result.Error != nil {
		return result.Error
	}
	return
}

func GetMysqlBackupDetailByFuzzySearchKey(db *gorm.DB, key string) []MysqlBackupDetail {
	var result []MysqlBackupDetail
	db.Where("db_name like ? or db_port like ?", fmt.Sprintf("%%%s%%", key), fmt.Sprintf("%%%s%%", key)).Find(&result)
	return result
}

func (u *MysqlBackupDetail) UpdateMysqlBackupBaseDetailByID(db *gorm.DB) (err error) {
	result := db.Model(u).Updates(map[string]interface{}{"RemoteServer": u.RemoteServer, "remote_dir": u.RemoteDir})
	if result.Error != nil {
		return result.Error
	}
	return
}
