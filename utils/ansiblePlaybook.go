package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/playbook"
	"hiveview"
	"io"
)

type AnsibleResult struct {
	Stats map[string]Stats `json:"stats"`
}

type Stats struct {
	Changed     int `json:"changed"`
	Failures    int `json:"failures"`
	Ignored     int `json:"ignored"`
	Ok          int `json:"ok"`
	Rescued     int `json:"rescued"`
	Skipped     int `json:"skipped"`
	Unreachable int `json:"unreachable"`
}

func AnsiblePlaybook(PlaybookName string, extraVars map[string]string) error {
	var err error
	var res = new(AnsibleResult)
	buff := new(bytes.Buffer)
	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
		//Connection: "local",
		User: "root",
	}

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory: hiveview.CONFIG.Settings.Ansible.Inventory,
	}
	for k, v := range extraVars {
		err = ansiblePlaybookOptions.AddExtraVar(k, v)
		if err != nil {
			LogPrint("err", err)
			return err
		}
	}

	Execute := execute.NewDefaultExecute(
		execute.WithWrite(io.Writer(buff)),
	)

	Playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:         []string{PlaybookName},
		Exec:              Execute,
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		StdoutCallback:    "json",
	}
	LogPrint("info", "begin")
	err = Playbook.Run(context.TODO())
	if err != nil {
		LogPrint("err", err)
		return err
	}
	LogPrint("info", "end")
	err = json.Unmarshal(buff.Bytes(), res)
	if err != nil {
		LogPrint("err", err)
		return err
	}

	if res.Stats[extraVars["hosts"]].Failures != 0 {

		err = fmt.Errorf("playbook执行失败")
	}
	return err
}
