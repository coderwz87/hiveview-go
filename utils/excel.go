package utils

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"reflect"
)

//生成EXCEL文件流
func WriteDataToExcel(xlsx *excelize.File, data interface{}) {
	firstCharacter := 65
	s := reflect.ValueOf(data)
	for i := 0; i < s.Len(); i++ {
		elem := s.Index(i).Interface()
		elemType := reflect.TypeOf(elem)
		elemValue := reflect.ValueOf(elem)
		for j := 0; j < elemType.NumField(); j++ {
			field := elemType.Field(j)
			tag := field.Tag.Get("xlsx")
			name := tag
			column := string(firstCharacter + j)
			if tag == "" {
				continue
			}
			// 设置表头
			if i == 0 {
				xlsx.SetCellValue("Sheet1", fmt.Sprintf("%s%d", column, i+1), name)
			}
			// 设置内容
			xlsx.SetCellValue("Sheet1", fmt.Sprintf("%s%d", column, i+2), elemValue.Field(j).Interface())
		}
	}
}
