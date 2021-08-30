package hiveview

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// 结构体定义
type Settings struct {
	Settings     Config `yaml:"settings"`
	Db           *gorm.DB
	Gin          *gin.Engine
	Logger       *logrus.Logger
	AccessLogger *logrus.Logger
	Enforcer     *casbin.Enforcer
}

type Config struct {
	Database    DatabaseConfig    `yaml:"database"`
	Application ApplicationConfig `yaml:"application"`
	Ansible     AnsibleConfig     `yaml:"ansible"`
	Log         LogConfig         `yaml:"log"`
	Enforcer    EnforcerConfig    `yaml:"enforcer"`
	Zabbix      Zabbix            `yaml:"zabbix"`
}

type Zabbix struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Url      string `yaml:"url"`
}
type DatabaseConfig struct {
	Host     string `yaml:host`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
}

type ApplicationConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type AnsibleConfig struct {
	Factsdir  string `yaml:"factsdir"`
	Inventory string `yaml:"inventory"`
}

type LogConfig struct {
	LogPath       string `yaml:"logpath"`
	LogName       string `yaml:"logname"`
	AccessLogPath string `yaml:"accesslogpath"`
	AccessLogName string `yaml:"accesslogname"`
}

type EnforcerConfig struct {
	ConfPath string `yaml:"confpath"`
}
