package asset

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
	"hiveview"
	"hiveview/models"
	"hiveview/utils"
)

//下载资产excel
func DownloadAssetInfo(c *gin.Context) {
	xlsx := excelize.NewFile()
	//设置单元格宽度
	xlsx.SetColWidth("Sheet1", "A", "J", 25)
	xlsx.SetColWidth("Sheet1", "K", "K", 40)
	xlsx.SetColWidth("Sheet1", "L", "S", 25)
	//xlsx.SetCellValue("Sheet1", "A1", "test")
	result, err := models.GetAllAsset(hiveview.CONFIG.Db)
	if err != nil {
		utils.LogPrint("err", err)
		return
	}

	centerStyle, err := xlsx.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		//Border: border,
	})
	if err != nil {
		utils.LogPrint("err", err)
	}
	err = xlsx.SetColStyle("Sheet1", "A:S", centerStyle)
	if err != nil {
		utils.LogPrint("err", err)
	}
	utils.WriteDataToExcel(xlsx, result)

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"Asset.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")

	_ = xlsx.Write(c.Writer)
}
