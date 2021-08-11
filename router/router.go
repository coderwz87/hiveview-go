package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"hiveview/controller/application"
	"hiveview/controller/asset"
	"hiveview/controller/backup"
	"hiveview/controller/commonTools"
	"hiveview/controller/dashboard"
	"hiveview/controller/operation"
	"hiveview/controller/service"
	"hiveview/controller/user"
	_ "hiveview/main/docs"
	"hiveview/middlewares"
)

func InitRouter(Gin *gin.Engine) {

	Gin.Use(middlewares.Cors())

	api := Gin.Group("/api", middlewares.LoggerToFile(), middlewares.JWTAuth(), middlewares.PermissionMiddleWare())
	api.POST("/login/", user.Login)
	api.POST("/ChangeUserPassword/", user.ChangeUserPassword)
	api.POST("/addUser/", user.AddUser)

	api.GET("/dashboard/", dashboard.Dashboard)
	api.GET("/dashboardUpdateLog/", dashboard.AppUpdateLog)

	api.POST("/addServer/", asset.AddServer)
	api.GET("/server/:id/", asset.GetServer)
	api.DELETE("/server/:id/", asset.DeleteServer)
	api.GET("/getAllServer/", asset.GetAllServer)
	api.PATCH("/server/:id/", asset.UpdateServer)
	api.DELETE("/deleteAllServer/", asset.DeleteAllServer)
	api.GET("/downloadAssetInfo/", asset.DownloadAssetInfo)
	api.GET("/searchServer/", asset.SearchServerInfo)
	api.GET("/getServerIdc/", asset.GetServerIdc)

	api.GET("/getAllAppDetail/", application.GetAllApplicationDetail)
	api.POST("/addAppDetail/", application.AddApplicationDetail)
	api.DELETE("/appDetail/:id/", application.DeleteApplicationDetail)
	api.DELETE("/deleteAllAppDetail/", application.DeleteAllAppDetail)
	api.PUT("/updateAppDetailState/", application.UpdateApplicationState)
	api.GET("/operationApp/", application.OperationApp)
	api.GET("/searchAppDetail/", application.SearchApplicationDetail)
	api.GET("/AppDetail/:id/", application.GetApplicationDetail)
	api.GET("/getAllAppName/", service.ResinProjectName)
	api.GET("/AppUpdateLog/", application.GetApplicationUpdateLog)
	api.POST("/AppUpdateLog/", application.CreateApplicationUpdateLog)
	api.GET("/searchAppUpdateLog/", application.SearchApplicationUpdateLog)
	api.PATCH("/AppDetail/:id/", application.UpdateApplicationBaseDetail)

	api.GET("/common/link/", commonTools.GetAllCommonLink)

	api.POST("/serviceInit/mysql/", service.MysqlInit)
	api.POST("/serviceInit/redis/", service.RedisInit)
	api.POST("/serviceInit/resin/", service.ResinInit)
	api.PUT("/serviceInit/batch/", service.BatchServiceInit)

	api.GET("/allMysqlBackupInfo/", backup.GetAllMysqlBackupDetail)
	api.POST("/addMysqlBackupDetail/", backup.AddMysqlBackupDetail)
	api.DELETE("/mysqlBackupDetail/:id/", backup.DeleteMysqlBackupDetail)
	api.DELETE("/deleteAllMysqlBackupDetail/", backup.DeleteAllMysqlBackupDetail)
	api.PUT("/updateMysqlBackupDetail/", backup.UpdateMysqlBackupDetail)
	api.GET("/searchMysqlBackupDetail/", backup.SearchMysqlBackupDetail)
	api.GET("/mysqlBackupDetail/:id/", backup.GetMysqlBackupDetail)
	api.PATCH("/mysqlBackupDetail/:id/", backup.UpdateMysqlBackupBaseDetail)

	api.GET("/allOtherBackupInfo/", backup.GetAllOtherBackupDetail)
	api.PUT("/updateOtherBackupDetail/", backup.UpdateOtherBackupDetail)
	api.GET("/operationDetail/:id/", operation.GetOperationDetail)
	api.POST("/OperationApprove/", operation.OperationApprove)
	api.GET("/GetAllOperationDetail/", operation.GetAllOperationDetail)
	url := ginSwagger.URL("http://127.0.0.1:9090/swagger/doc.json")
	Gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
