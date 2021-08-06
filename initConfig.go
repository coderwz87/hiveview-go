package hiveview

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hiveview/models"
	"io/ioutil"
	"log"
	"path"
	"time"
)

var (
	CONFIG = new(Settings)
)

//根据配置文件初始化配置
func (CONFIG *Settings) InitConfig(fileName string) {

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(fmt.Sprintf("open config file fail: %s", err.Error()))
	}
	err = yaml.Unmarshal(data, CONFIG)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unmarshal config object fail: %s", err.Error()))
	}

}

// 初始化数据库连接
func (CONFIG *Settings) InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", CONFIG.Settings.Database.User, CONFIG.Settings.Database.Password, CONFIG.Settings.Database.Host, CONFIG.Settings.Database.Port, CONFIG.Settings.Database.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(fmt.Sprintf("OPEN DB fail: %s", err.Error()))
	}
	CONFIG.Db = db
}

// 数据库初始化
func (CONFIG *Settings) Migrate() {
	CONFIG.Db.AutoMigrate(&models.Users{}, &models.Assets{}, &models.CommonLink{}, &models.MysqlBackupDetail{}, &models.OtherBackupDetail{}, &models.AppDetail{}, &models.ApplicationUpdateLog{})
	//CONFIG.Db.AutoMigrate(&models.Assets{})

}

// gin初始化

func (CONFIG *Settings) InitGin() {
	CONFIG.Gin = gin.Default()
}

//logger初始化

func (CONFIG *Settings) InitLogger() {
	logFilePath := CONFIG.Settings.Log.LogPath
	logFileName := CONFIG.Settings.Log.LogName
	fileName := path.Join(logFilePath, logFileName)
	writer, _ := rotatelogs.New(
		fileName+".%Y%m%d",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(time.Duration(24*7)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)
	logger := logrus.New()
	logger.SetOutput(writer)
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.TextFormatter{})
	CONFIG.Logger = logger
}

func (CONFIG *Settings) InitAccessLogger() {
	logFilePath := CONFIG.Settings.Log.AccessLogPath
	logFileName := CONFIG.Settings.Log.AccessLogName
	fileName := path.Join(logFilePath, logFileName)
	writer, _ := rotatelogs.New(
		fileName+".%Y%m%d",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(time.Duration(24*7)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)
	logger := logrus.New()
	logger.SetOutput(writer)
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.TextFormatter{})
	CONFIG.AccessLogger = logger

}

//cron 初始化
func InitCron() *cron.Cron {
	c := cron.New()
	return c
}

//初始化管理员账号密码
func InitUser() {
	ifCreateAdminUser := models.IfInitAdminUser(CONFIG.Db)
	if ifCreateAdminUser {
		var admin = models.Users{
			Username: "admin",
			Password: "123456",
		}
		err := admin.CreateUser(CONFIG.Db)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

//初始化enforcer
func (CONFIG *Settings) InitEnforcer() error {
	var enforcerMysql = fmt.Sprintf("%s:%s@tcp(%s:%s)/", CONFIG.Settings.Database.User, CONFIG.Settings.Database.Password, CONFIG.Settings.Database.Host, CONFIG.Settings.Database.Port)
	link, _ := gormadapter.NewAdapter("mysql", enforcerMysql)
	var err error
	CONFIG.Enforcer, err = casbin.NewEnforcer(CONFIG.Settings.Enforcer.ConfPath, link)
	if err != nil {
		return err
	}

	CONFIG.Enforcer.EnableLog(true)
	CONFIG.Enforcer.LoadPolicy()

	return nil
}
