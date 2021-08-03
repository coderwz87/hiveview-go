package utils

import (
	"bytes"
	"encoding/json"
	"hiveview/models"
	"io/ioutil"
	"net/http"
)

const (
	USERNAME = "Admin"
	PASSWORD = "hiveview@2020"
	URL      = "http://124.205.11.222/api_jsonrpc.php"
)

type TokenRes struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`
	ID      int    `json:"id"`
}

type HostIdRes struct {
	Jsonrpc string              `json:"jsonrpc"`
	Result  []map[string]string `json:"result"`
	ID      int                 `json:"id"`
}

func ZabbixGetToken() (token string, err error) {
	var loginData = make(map[string]interface{})
	loginData["jsonrpc"] = "2.0"
	loginData["method"] = "user.login"
	loginData["params"] = map[string]string{
		"user":     USERNAME,
		"password": PASSWORD,
	}
	loginData["id"] = 1
	bytesData, err := json.Marshal(loginData)
	if err != nil {
		LogPrint("err", err)
		return
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", URL, reader)
	if err != nil {
		LogPrint("err", err)
		return
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		LogPrint("err", err)
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		LogPrint("err", err)
		return
	}

	var result = new(TokenRes)
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		LogPrint("err", err)
		return
	}
	return result.Result, nil
}

func ZabbixGetHostId(token string, hostname string) (id string, err error) {
	var hostIdData = make(map[string]interface{})
	hostIdData["jsonrpc"] = "2.0"
	hostIdData["method"] = "host.get"
	hostIdData["params"] = map[string]interface{}{
		"output": "host",
		"filter": map[string]string{"host": hostname},
	}
	hostIdData["auth"] = token
	hostIdData["id"] = 1

	bytesData, err := json.Marshal(hostIdData)
	if err != nil {
		LogPrint("err", err)
		return
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", URL, reader)
	if err != nil {
		LogPrint("err", err)
		return
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		LogPrint("err", err)
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		LogPrint("err", err)
		return
	}
	var result = new(HostIdRes)
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		LogPrint("err", err)
		return
	}
	return result.Result[0]["hostid"], err
}

func ZabbixDeleteHost(asset *models.Assets) (err error) {
	token, err := ZabbixGetToken()
	if err != nil {
		LogPrint("err", err)
		return
	}
	hostId, err := ZabbixGetHostId(token, asset.Hostname)
	var deleteData = map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "host.delete",
		"params":  []string{hostId},
		"auth":    token,
		"id":      1,
	}
	bytesData, err := json.Marshal(deleteData)
	if err != nil {
		LogPrint("err", err)
		return
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", URL, reader)
	if err != nil {
		LogPrint("err", err)
		return
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	_, err = client.Do(request)
	if err != nil {
		LogPrint("err", err)
		return
	}
	return err
}
