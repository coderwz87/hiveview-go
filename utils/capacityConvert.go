package utils

import (
	"log"
	"strconv"
	"strings"
)

//单位转换
func CapacityConvert(size string, except string) (float64, string) {
	rate_map := map[string]float64{
		"KB": 1024,
		"MB": 1024 * 1024,
		"GB": 1024 * 1024 * 1024,
		"TB": 1024 * 1024 * 1024 * 1024,
	}
	var std_size float64
	for k, v := range rate_map {
		if strings.HasSuffix(size, k) {
			sizeInt, err := strconv.ParseFloat(strings.Trim(strings.Trim(size, k), " "), 64)
			if err != nil {
				log.Fatal("CapacityConvert err", err.Error())
			}
			std_size = sizeInt * v
		}
	}
	if except == "auto" {
		for k, v := range rate_map {
			if (std_size/v >= 1.0 && std_size/v < 1024.0) || k == "TB" {
				except = k
				break
			}
		}
	}
	expect_size := float64(std_size / rate_map[except])
	return expect_size, except

}
