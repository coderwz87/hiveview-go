package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Assets struct {
	SN             string `json:"sn" gorm:"type:varchar(128);" xlsx:"序列号"`             //服务器序列号
	IP             string `json:"ip" gorm:"type:varchar(256);" xlsx:"IP地址"`            //服务器ip地址
	Mem            int    `json:"mem" gorm:"type:varchar(128);" xlsx:"内存大小"`           //服务器内存大小
	Vendor         string `json:"vendor" gorm:"type:varchar(128);" xlsx:"服务器品牌"`       //服务器供应商
	SeverModel     string `json:"sever_model" gorm:"type:varchar(128);" xlsx:"服务器型号"`  //服务器型号
	CpuCount       int    `json:"cpu_count" gorm:"type:varchar(128);" xlsx:"cpu核心数"`   //cpu核心数
	OS             string `json:"os" gorm:"type:varchar(128);" xlsx:"系统版本"`            //服务器系统版本
	OSVersion      string `json:"os_version" gorm:"type:varchar(128);" xlsx:"系统版本号"`   //服务器系统版本号
	OSArch         string `json:"os_arch" gorm:"type:varchar(128);" xlsx:"系统位数"`       //服务器系统位数
	Hostname       string `json:"hostname" gorm:"type:varchar(128);" xlsx:"主机名"`       //服务器主机名
	CpuModel       string `json:"cpu_model" gorm:"type:varchar(128);" xlsx:"cpu型号"`    //cpu型号
	DiskTotal      string `json:"disk_total" gorm:"type:varchar(128);" xlsx:"硬盘空间"`    //硬盘空间大小
	Comment        string `json:"comment" gorm:"type:varchar(128);" xlsx:"备注"`         //备注
	IDC            string `json:"idc" gorm:"type:varchar(128);column:idc" xlsx:"IDC"`  //归属IDC
	Use            string `json:"use" gorm:"type:varchar(256);" xlsx:"用途"`             //服务器用途
	Cabinet        string `json:"cabinet" gorm:"type:varchar(128)" xlsx:"机柜"`          //服务器机柜
	UPosition      string `json:"uposition" gorm:"type:varchar(128)" xlsx:"U位"`        //服务器U位
	ServerType     string `json:"server_type" gorm:"type:varchar(56)" xlsx:"服务器类型"`    //服务器类型,物理机或者虚拟机
	InterfaceSpeed string `json:"interface_speed" gorm:"type:varchar(56)" xlsx:"网卡速率"` //网卡速率
	gorm.Model
}

func (u *Assets) TableName() string {
	return "Assets"
}

func (u *Assets) CreateAsset(db *gorm.DB) (err error) {
	result := db.Create(u)
	if result.Error != nil {
		return result.Error
	}
	return
}

func (u *Assets) UpdateAssetInfo(db *gorm.DB) (err error) {
	result := db.Save(u)
	if result.Error != nil {
		return result.Error
	}
	return
}

func (u *Assets) GetAssetByID(db *gorm.DB) (err error) {
	result := db.Where("id=?", u.ID).First(u)
	if result.Error != nil {
		return result.Error
	}
	return
}

func (u *Assets) GetAssetByIP(db *gorm.DB) bool {
	result := db.Where("ip=?", u.IP).First(u)
	//if result.Error != nil {
	//	return result.Error
	//}
	if result.RowsAffected != 0 {
		return true
	}
	return false
}

func (u *Assets) DeleteAssetByID(db *gorm.DB) (err error) {
	err = u.GetAssetByID(db)
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

func (u *Assets) UpdateAssetByID(db *gorm.DB) (err error) {
	result := db.Model(u).Updates(&Assets{Use: u.Use, IDC: u.IDC, Cabinet: u.Cabinet, UPosition: u.UPosition, Comment: u.Comment})
	if result.Error != nil {
		return result.Error
	}
	return
}

func GetAllAsset(db *gorm.DB) (resultList []Assets, err error) {

	r := db.Find(&resultList)
	return resultList, r.Error
}

func GetAssetByFuzzyIP(db *gorm.DB, key string) []Assets {
	var result []Assets
	db.Where("ip like ? or hostname like ?", fmt.Sprintf("%%%s%%", key), fmt.Sprintf("%%%s%%", key)).Find(&result)
	return result
}

func GetAssetIdc(db *gorm.DB) []string {
	var resultList []Assets
	var result []string
	db.Distinct("idc").Find(&resultList)
	for _, v := range resultList {
		result = append(result, v.IDC)
	}
	return result
}

func GetAssetIDCCount(db *gorm.DB) (int, int) {
	var beijingResultList []Assets
	var otherResultList []Assets
	db.Where("idc = '北京'").Find(&beijingResultList)
	db.Where("idc != '北京'").Find(&otherResultList)
	return len(beijingResultList), len(otherResultList)

}

func GetAssetUseCount(db *gorm.DB) (int, int) {
	var liveResultList []Assets
	var hlsResultList []Assets
	db.Where("`use` = '直播' ").Find(&liveResultList)
	db.Where("`use` = '点播'").Find(&hlsResultList)
	return len(liveResultList), len(hlsResultList)
}

func GetBeiJingAsset(db *gorm.DB) ([]Assets, error) {
	var beijingResultList []Assets
	r := db.Where("idc = '北京'").Find(&beijingResultList)
	return beijingResultList, r.Error
}

func GetVirtualAndPhysicsCount(db *gorm.DB) (int, int) {
	var VirtualResultList []Assets
	var PhysicsResultList []Assets
	db.Where("sever_model = 'KVM'").Find(&VirtualResultList)
	db.Where("sever_model != 'KVM'").Find(&PhysicsResultList)
	return len(VirtualResultList), len(PhysicsResultList)
}
