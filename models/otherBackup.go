package models

import "gorm.io/gorm"

type OtherBackupDetail struct {
	ProjectName  string `json:"project_name" gorm:"type:varchar(128);"`
	RemoteServer string `json:"remote_server" gorm:"type:varchar(128);"`
	RemoteDir    string `json:"remote_dir" gorm:"type:varchar(128);"`
	BackupLog    string `json:"backup_log" gorm:"type:varchar(256);"`
	gorm.Model
}

func (u *OtherBackupDetail) TableName() string {
	return "other_backup_detail"
}

func (u *OtherBackupDetail) CreateOtherBackupDetail(db *gorm.DB) (err error) {
	result := db.Create(u)
	if result.Error != nil {
		return result.Error
	}
	return
}

func GetAllOtherBackupDetail(db *gorm.DB) (resultList []OtherBackupDetail, err error) {

	r := db.Find(&resultList)
	return resultList, r.Error
}

func (u *OtherBackupDetail) UpdateOtherBackupDetail(db *gorm.DB) (err error) {
	result := db.Model(u).Where("project_name = ? and remote_server = ? and remote_dir = ?", u.ProjectName, u.RemoteServer, u.RemoteDir).Update("backup_log", u.BackupLog)
	if result.Error != nil {
		return result.Error
	}
	return
}
