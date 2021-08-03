package dashboard

import (
	"github.com/gin-gonic/gin"
	"hiveview"
	"hiveview/models"
	"hiveview/render"
	"hiveview/utils"
)

type ResData struct {
	BeijingCount int      `json:"beijing_count"`
	OtherCount   int      `json:"other_count"`
	LiveCount    int      `json:"live_count"`
	HlsCount     int      `json:"hls_count"`
	VirtualCount int      `json:"virtual_count"`
	PhysicsCount int      `json:"physics_count"`
	XAxis        []string `json:"xAxis"`
	Series       []int    `json:"series"`
}

func Dashboard(c *gin.Context) {
	beijingCount, otherCount := models.GetAssetIDCCount(hiveview.CONFIG.Db)
	liveCount, hlsCount := models.GetAssetUseCount(hiveview.CONFIG.Db)
	var data = new(ResData)
	data.BeijingCount = beijingCount
	data.OtherCount = otherCount
	data.LiveCount = liveCount
	data.HlsCount = hlsCount
	VirtualCount, PhysicsCount := models.GetVirtualAndPhysicsCount(hiveview.CONFIG.Db)
	data.VirtualCount = VirtualCount
	data.PhysicsCount = PhysicsCount

	render.JSON(c, data)
}

func AppUpdateLog(c *gin.Context) {
	var data = new(ResData)
	XAxis := utils.GetDate(7)
	data.XAxis = XAxis
	Series := models.GetAppUpdateCount(hiveview.CONFIG.Db, XAxis)
	data.Series = Series
	render.JSON(c, data)
}
