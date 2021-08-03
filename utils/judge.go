package utils

import (
	"github.com/go-ping/ping"
	"hiveview"
	"hiveview/models"
	"net"
	"os"
	"time"
)

//判断是否为IP
func JudgeType(ip string) bool {
	address := net.ParseIP(ip)
	if address == nil {
		return false
	} else {
		return true
	}
}

//判断IP是否存活
func JudgeAlive(ip string) bool {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		return false
	}
	pinger.Count = 3
	pinger.Timeout = time.Duration(1000 * time.Millisecond)
	pinger.SetPrivileged(true)
	pinger.Run()
	stats := pinger.Statistics()
	return stats.PacketsRecv >= 1
}

//判断文件是否存在
func JudgeFileExist(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

//判断资产表中是否存在此ip地址
func JudgeInMysql(ip string) bool {
	item := new(models.Assets)
	item.IP = ip
	return item.GetAssetByIP(hiveview.CONFIG.Db)
}
