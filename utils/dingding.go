package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DingMsg struct {
	Msgtype string            `json:"msgtype"`
	Text    map[string]string `json:"text"`
	At      DingAt            `json:"at"`
}

type DingAt struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

//发送钉钉报警
func SendMsgToDD(msg string) (err error) {
	URL := "https://oapi.dingtalk.com/robot/send?access_token=8c095e756df4b68c5409be46c69bf013e502cef4dce194ff1d3eb222355fb425"
	data := new(DingMsg)
	data.Msgtype = "text"
	data.Text = make(map[string]string)
	data.Text["content"] = msg
	data.At.AtMobiles = []string{"test"}
	data.At.IsAtAll = false
	bytesData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", URL, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	_, err = client.Do(request)
	return
}
