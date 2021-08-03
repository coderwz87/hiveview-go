package utils

import (
	"encoding/json"
	"fmt"
	"hiveview"
	"hiveview/models"
	"io/ioutil"
	"os/exec"
	"strings"
)

type Result struct {
	AnsibleFacts AssetInfo `json:"ansible_facts"`
}

type AssetInfo struct {
	SN        string          `json:"ansible_product_serial"`       //服务器序列号
	Mem       int             `json:"ansible_memtotal_mb"`          //服务器内存大小
	Vendor    string          `json:"ansible_system_vendor"`        //服务器供应商
	Model     string          `json:"ansible_product_name"`         //服务器型号
	CpuCount  int             `json:"ansible_processor_vcpus"`      //cpu核心数
	OS        string          `json:"ansible_distribution"`         //服务器系统版本
	OSVersion string          `json:"ansible_distribution_version"` //服务器系统版本号
	OSArch    string          `json:"ansible_architecture"`         //服务器系统位数
	Hostname  string          `json:"ansible_hostname"`             //服务器主机名
	CpuModel  []string        `json:"ansible_processor"`            //cpu型号
	DiskInfo  map[string]Disk `json:"ansible_devices"`              //硬盘信息
	Interface InterfaceInfo   `json:"ansible_default_ipv4"`         //网卡信息
	//IP        []string        `json:"ansible_all_ipv4_addresses"`   //服务器ip地址
}

type Disk struct {
	Removable string `json:"removable"`
	Size      string `json:"size"`
}

type InterfaceInfo struct {
	Address   string `json:"address"`
	Interface string `json:"interface"`
}

//获取资产信息
func (result *Result) GetAssetInfo(file string, ip string, info *models.Assets) (ServerInfo *models.Assets, err error) {
	ifExist := JudgeFileExist(file)
	if !ifExist {
		cmd := fmt.Sprintf("ansible -i %s %s -m setup --tree %s ", hiveview.CONFIG.Settings.Ansible.Inventory, ip, hiveview.CONFIG.Settings.Ansible.Factsdir)
		CMD := exec.Command("/bin/bash", "-c", cmd)
		_, err = CMD.Output()
		if err != nil {
			LogPrint("err", err)
			return
		}

	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		LogPrint("error", "do not find this asset info")
		return
	}
	err = json.Unmarshal(data, result)
	if err != nil {
		LogPrint("error", "unmarshal data err ")
		return
	}
	info.SN = result.AnsibleFacts.SN
	info.Mem = result.AnsibleFacts.Mem
	info.Vendor = result.AnsibleFacts.Vendor
	info.SeverModel = result.AnsibleFacts.Model
	info.CpuCount = result.AnsibleFacts.CpuCount
	info.OS = result.AnsibleFacts.OS
	info.OSVersion = result.AnsibleFacts.OSVersion
	info.OSArch = result.AnsibleFacts.OSArch
	info.Hostname = result.AnsibleFacts.Hostname
	info.CpuModel = result.GetAssetCpuModel()
	info.DiskTotal = result.GetDiskTotal()

	if result.AnsibleFacts.Model == "KVM" {
		info.ServerType = "虚拟机"
		info.InterfaceSpeed = ""
	} else {
		info.ServerType = "物理机"
		InterfaceSpeed := result.AnsibleFacts.Interface.GetInterfaceSpeed(ip)
		if len(InterfaceSpeed) == 0 {
			info.InterfaceSpeed = "网卡速率有误"
		} else {
			info.InterfaceSpeed = InterfaceSpeed
		}

	}
	return info, nil
}

//获取网卡网速信息
func (interfaceInfo InterfaceInfo) GetInterfaceSpeed(ip string) (speed string) {
	if interfaceInfo.Address != ip {
		LogPrint("err", "网卡信息有误")
		return
	}
	cmd := fmt.Sprintf("ansible -i %s %s -m shell -a 'ethtool %s' |grep Speed|awk -F: '{print $2}' ", hiveview.CONFIG.Settings.Ansible.Inventory, ip, interfaceInfo.Interface)
	CMD := exec.Command("bash", "-c", cmd)
	out, err := CMD.Output()
	if err != nil {
		LogPrint("err", err)
		return
	}
	speed = strings.Trim(strings.Trim(string(out), " "), "\n")
	return speed
}

//获取cpu型号信息
func (result *Result) GetAssetCpuModel() (CpuModel string) {
	for _, v := range result.AnsibleFacts.CpuModel {
		if strings.HasPrefix(v, "Intel") || strings.HasSuffix(v, "GHz") {
			CpuModel = v
			break
		}
	}
	return
}

//获取硬盘总大小
func (result *Result) GetDiskTotal() (DiskTotal string) {
	var DiskList []Disk
	for k, v := range result.AnsibleFacts.DiskInfo {
		if strings.HasPrefix(k, "loop") || strings.HasPrefix(k, "ram") || strings.HasPrefix(k, "dm") {
			continue
		}
		DiskList = append(DiskList, v)
	}
	var TotalSize float64
	for _, v := range DiskList {
		size, _ := CapacityConvert(v.Size, "MB")
		TotalSize += size
	}
	DiskTotal = fmt.Sprintf("%f MB", TotalSize)
	DiskTotalSize, DiskTotalUnit := CapacityConvert(DiskTotal, "auto")
	DiskTotal = fmt.Sprintf("%.2f %s", DiskTotalSize, DiskTotalUnit)
	return
}
