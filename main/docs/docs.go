// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/addAppDetail/": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "添加应用信息",
                "tags": [
                    "Application"
                ],
                "operationId": "add-application-detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "app名字",
                        "name": "app_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "主机名",
                        "name": "host",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "目录",
                        "name": "dir",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "类型",
                        "name": "type",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"创建成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/addMysqlBackupDetail/": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "添加Mysql备份信息",
                "tags": [
                    "mysqlBackup"
                ],
                "operationId": "add-mysql-backup-detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "数据库实例名字",
                        "name": "db_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "实例端口号",
                        "name": "db_port",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "备份服务器地址",
                        "name": "remote_server",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "备份目录",
                        "name": "remote_dir",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"已创建\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/addServer/": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "添加资产信息",
                "tags": [
                    "asset"
                ],
                "operationId": "add-server",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ip",
                        "name": "ip",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "idc名字",
                        "name": "idc",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "备注",
                        "name": "comment",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "用途",
                        "name": "use",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "机柜号",
                        "name": "cabinet",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "U位",
                        "name": "uPosition",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"已开始创建\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/allMysqlBackupInfo/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取全部Mysql备份信息",
                "tags": [
                    "mysqlBackup"
                ],
                "operationId": "get-all-mysql-backup-detail",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"success\",\"data\":[...]}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/allOtherBackupInfo/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取全部其他项目备份信息",
                "tags": [
                    "otherBackup"
                ],
                "operationId": "get-all-other-backup-detail",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"success\",\"data\":[...]}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/appDetail/{id}/": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "删除单条信息",
                "tags": [
                    "Application"
                ],
                "operationId": "delete-application-detail",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "应用id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/common/link/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取全部常用连接信息",
                "tags": [
                    "common"
                ],
                "operationId": "get-all-common-link",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"success\",\"data\":[...]}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/deleteAllAppDetail/": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "批量删除应用信息",
                "tags": [
                    "Application"
                ],
                "operationId": "delete-all-application-detail",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "name": "detail_ids",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"已全部删除\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/deleteAllMysqlBackupDetail/": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "删除全部Mysql备份信息",
                "tags": [
                    "mysqlBackup"
                ],
                "operationId": "delete-all-mysql-backup-detail",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "name": "detail_ids",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"已全部删除\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/deleteAllServer/": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "批量删除资产信息",
                "tags": [
                    "asset"
                ],
                "operationId": "delete-all-server",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "name": "server_ids",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"已全部删除\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/getAllAppDetail/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取全部应用信息",
                "tags": [
                    "Application"
                ],
                "operationId": "get-all-application-detail",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"success\",\"data\":[...]}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/getAllServer/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取全部资产信息",
                "tags": [
                    "asset"
                ],
                "operationId": "get-all-server",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"success\",\"data\":[...]}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/mysqlBackupDetail/{id}/": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "删除单条Mysql备份信息",
                "tags": [
                    "mysqlBackup"
                ],
                "operationId": "delete-mysql-backup-detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/operationApp/": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "操作应用",
                "tags": [
                    "Application"
                ],
                "operationId": "operation-application",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "操作动作",
                        "name": "action",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"已开始操作\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/server/{id}/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取单台资产信息",
                "tags": [
                    "asset"
                ],
                "operationId": "get-server",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"success\",\"data\":{object}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "删除资产信息",
                "tags": [
                    "asset"
                ],
                "operationId": "delete-server",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "更新资产信息",
                "tags": [
                    "asset"
                ],
                "operationId": "update-server",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "idc名字",
                        "name": "idc",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "备注",
                        "name": "comment",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "用途",
                        "name": "use",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "机柜号",
                        "name": "cabinet",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "U位",
                        "name": "uPosition",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"更新成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/serviceInit/mysql/": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "mysql初始化",
                "tags": [
                    "service"
                ],
                "operationId": "mysql-init",
                "parameters": [
                    {
                        "type": "string",
                        "description": "初始化服务器",
                        "name": "ip",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "数据库版本",
                        "name": "version",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "实例端口号",
                        "name": "port",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "实例名",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"已开始初始化\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/serviceInit/redis/": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "redis初始化",
                "tags": [
                    "service"
                ],
                "operationId": "redis-init",
                "parameters": [
                    {
                        "type": "string",
                        "description": "初始化服务器",
                        "name": "ip",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "redis最大内存",
                        "name": "memsize",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "实例端口号",
                        "name": "port",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "实例名",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"已开始初始化\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/updateAppDetailState/": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "更新应用状态",
                "tags": [
                    "Application"
                ],
                "operationId": "update-application-detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "app名字",
                        "name": "app_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "主机名",
                        "name": "host",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "目录",
                        "name": "dir",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "类型",
                        "name": "type",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "状态",
                        "name": "state",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"已更新\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/updateMysqlBackupDetail/": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "更新mysql备份信息",
                "tags": [
                    "mysqlBackup"
                ],
                "operationId": "update-mysql-backup-detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "数据库实例名",
                        "name": "db_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "端口号",
                        "name": "db_port",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "备份服务器地址",
                        "name": "remote_server",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "备份目录",
                        "name": "remote_dir",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "备份信息",
                        "name": "backup_log",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"已更新\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/updateOtherBackupDetail/": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "更新其他项目备份信息",
                "tags": [
                    "otherBackup"
                ],
                "operationId": "update-other-backup-detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "项目名字",
                        "name": "project_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "备份服务器地址",
                        "name": "remote_server",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "备份目录",
                        "name": "remote_dir",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "备份信息",
                        "name": "backup_log",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"已更新\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "token",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0.0",
	Host:        "",
	BasePath:    "/api/",
	Schemes:     []string{},
	Title:       "运维平台",
	Description: "运维平台",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
