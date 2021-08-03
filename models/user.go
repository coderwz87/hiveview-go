package models

import (
	"crypto/md5"
	"fmt"
	"gorm.io/gorm"
)

type Users struct {
	Username string `json:"username" gorm:"type:varchar(100);comment:用户名" form:"username" binding:"required"`
	Password string `json:"password" gorm:"type:varchar(128);comment:密码MD5" form:"password" binding:"required"`
	gorm.Model
}

func (u *Users) TableName() string {
	return "Users"
}

func (u *Users) CreateUser(db *gorm.DB) (err error) {
	u.Password = u.PasswordMD5()
	result := db.Create(u)
	if result.Error != nil {
		return result.Error
	}
	return
}

func (u *Users) Verify(db *gorm.DB) bool {
	psswordMD5 := u.PasswordMD5()
	db.Where("username=?", u.Username).First(u)
	return psswordMD5 == u.Password
}

func (u *Users) PasswordMD5() string {
	password := []byte(u.Password)
	has := md5.Sum(password)
	return fmt.Sprintf("%x", has)
}
