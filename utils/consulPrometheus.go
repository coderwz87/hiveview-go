package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hiveview/models"
	"net/http"
	"strings"
)

const (
	RegisterUrl   = "http://124.205.11.219:8500/v1/agent/service/register?replace-existing-checks=1"
	DeRegisterUrl = "http://124.205.11.219:8500/v1/agent/service/deregister/"
)

type RequestStruct struct {
	ID                string            `json:"ID"`
	Name              string            `json:"Name"`
	Tags              []string          `json:"Tags"`
	Address           string            `json:"Address"`
	Port              int               `json:"Port"`
	Meta              map[string]string `json:"Meta"`
	EnableTagOverride bool              `json:"EnableTagOverride"`
	Check             map[string]string `json:"Check"`
	Weights           map[string]int    `json:"Weights"`
}

func PutInfoToConsul(AssetInfo *models.Assets, exporterName string) error {
	var RequestData RequestStruct
	RequestData.ID = fmt.Sprintf("%s-%s", exporterName, AssetInfo.IP)
	RequestData.Name = exporterName
	RequestData.Tags = []string{exporterName}
	RequestData.Address = AssetInfo.IP
	RequestData.Port = 9100
	MetaData := make(map[string]string)
	MetaData["idc"] = AssetInfo.IDC
	MetaData["use"] = AssetInfo.Use
	MetaData["hostname"] = AssetInfo.Hostname
	if AssetInfo.Comment != "" {
		MetaData["comment"] = AssetInfo.Comment
	}
	RequestData.Meta = MetaData
	RequestData.EnableTagOverride = false
	CheckData := make(map[string]string)
	CheckData["HTTP"] = fmt.Sprintf("http://%s:9100/metrics", AssetInfo.IP)
	CheckData["Interval"] = "10s"
	RequestData.Check = CheckData
	WeightsData := make(map[string]int)
	WeightsData["Passing"] = 10
	WeightsData["Warning"] = 1
	RequestData.Weights = WeightsData
	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(RequestData)
	req, err := http.NewRequest("PUT", RegisterUrl, requestBody)
	if err != nil {
		return err
	}
	client := &http.Client{}

	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func DeleteInfoFromConsul(AssetInfo models.Assets, exporterName string) error {
	DeleteURL := fmt.Sprintf("%s%s-%s", DeRegisterUrl, exporterName, AssetInfo.IP)
	req, _ := http.NewRequest("PUT", DeleteURL, strings.NewReader(""))
	http.DefaultClient.Do(req)
	return nil
}
