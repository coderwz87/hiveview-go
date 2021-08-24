package batch

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hiveview"
	"hiveview/models"
	"hiveview/render"
	"hiveview/utils"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func UploadFile(c *gin.Context) {
	f, err := c.FormFile("filename")
	if err != nil {
		render.DataError(c, "文件上传错误")
		utils.LogPrint("err", err)
		return
	}
	savePath := fmt.Sprintf("/etc/ansible/push_file/%s", f.Filename)
	err = c.SaveUploadedFile(f, savePath)
	if err != nil {
		render.DataError(c, "保存文件异常")
		utils.LogPrint("err", err)
		return
	}
	render.MSG(c, "已上传")
}

func FileList(c *gin.Context) {
	files, err := ioutil.ReadDir("/etc/ansible/push_file/")
	if err != nil {
		utils.LogPrint("err", err)
		render.DataError(c, "获取文件失败")
		return
	}
	var result []string
	for _, file := range files {
		result = append(result, file.Name())
	}

	render.JSON(c, result)
}

func GroupList(c *gin.Context) {
	group := models.GetAssetComment(hiveview.CONFIG.Db)
	render.JSON(c, group)
}

func GroupIpList(c *gin.Context) {
	groupName := c.PostForm("group")
	result := models.GetAssetByComment(hiveview.CONFIG.Db, groupName)
	render.JSON(c, result)
}

func FilePush(c *gin.Context) {
	ip := c.PostForm("ip")
	targetDir := c.PostForm("target_dir")
	fileName := c.PostForm("filename")
	groupName := c.PostForm("group")
	Type := c.PostForm("type")
	var logRecord = new(models.BatchLog)
	var IP string
	if Type == "single" {
		IP = ip
		logRecord.Host = IP
	} else if Type == "multi" {
		ipList := models.GetAssetByComment(hiveview.CONFIG.Db, groupName)
		IP = strings.Join(ipList, ",")
		logRecord.Host = groupName
	}
	arg := fmt.Sprintf("src=/etc/ansible/push_file/%s dest=%s", fileName, targetDir)
	logPath := fmt.Sprintf("/tmp/ansibleBatchLog/%d.log", time.Now().Unix())

	logRecord.Type = "push file"
	logRecord.LogFile = logPath
	logRecord.Detail = fmt.Sprintf("push %s to %s", fileName, targetDir)
	err := logRecord.CreateBatchLog(hiveview.CONFIG.Db)
	if err != nil {
		utils.LogPrint("err", err)
	}
	go func() {
		err = utils.AnsibleAdhoc("copy", arg, IP, logPath)
		if err != nil {
			utils.LogPrint("err", err)
			return
		}
	}()
	render.MSG(c, "已开始推送")
}

func GetAllBatchLog(c *gin.Context) {
	data, err := models.GetAllBatchLog(hiveview.CONFIG.Db)
	if err != nil {
		render.DataError(c, fmt.Sprintf("查询异常:%s", err))
		return
	}
	render.JSON(c, data)
}

func GetBatchLogDetail(c *gin.Context) {
	id := c.Param("id")
	result := new(models.BatchLog)
	ID, err := strconv.Atoi(id)
	if err != nil {
		render.DataError(c, err.Error())
		return
	}
	result.ID = uint(ID)
	err = result.GetBatchLogByID(hiveview.CONFIG.Db)
	if err != nil {
		utils.LogPrint("err", err)
		render.DataError(c, err.Error())
		return
	}

	f, err := os.Open(result.LogFile)

	defer f.Close()
	if err != nil {
		utils.LogPrint("err", err)
		render.DataError(c, "读取文件失败")
		return
	}
	bytes, _ := ioutil.ReadAll(f)
	context := string(bytes)

	render.JSON(c, context)
}

func CommandExec(c *gin.Context) {
	ip := c.PostForm("ip")
	Type := c.PostForm("type")
	groupName := c.PostForm("group")
	Command := c.PostForm("command")
	var logRecord = new(models.BatchLog)
	var IP string
	if Type == "single" {
		IP = ip
		logRecord.Host = IP
	} else if Type == "multi" {
		ipList := models.GetAssetByComment(hiveview.CONFIG.Db, groupName)
		IP = strings.Join(ipList, ",")
		logRecord.Host = groupName
	}
	logPath := fmt.Sprintf("/tmp/ansibleBatchLog/%d.log", time.Now().Unix())
	logRecord.Type = "exec shell"
	logRecord.LogFile = logPath
	logRecord.Detail = fmt.Sprintf("exec %s", Command)

	go func() {
		err := utils.AnsibleAdhoc("shell", Command, IP, logPath)
		if err != nil {
			utils.LogPrint("err", err)
			return
		}
		err = logRecord.CreateBatchLog(hiveview.CONFIG.Db)
		if err != nil {
			utils.LogPrint("err", err)
		}

	}()
	render.MSG(c, "已开始执行")
}
