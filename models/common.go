package models

import "gorm.io/gorm"

type CommonLink struct {
	Label string `json:"label" gorm:"type:varchar(256);"`
	Link  string `json:"link" gorm:"type:varchar(256);"`
	gorm.Model
}

func (u *CommonLink) TableName() string {
	return "CommonLink"
}

func GetAllLink(db *gorm.DB) (resultList []CommonLink, err error) {

	r := db.Find(&resultList)
	return resultList, r.Error
}

func (u *CommonLink) CreateLink(db *gorm.DB) (err error) {
	result := db.Create(u)
	if result.Error != nil {
		return result.Error
	}
	return
}
