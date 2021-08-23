package utils

import (
	"bytes"
	"context"
	"github.com/apenella/go-ansible/pkg/adhoc"
	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/options"
	"hiveview"
	"io"
	"io/ioutil"
)

func AnsibleAdhoc(moduleName, args, targetHost, Logfile string) error {
	buff := new(bytes.Buffer)
	ansibleAdhocConnectionOptions := &options.AnsibleConnectionOptions{
		User: "root",
	}
	ansibleAdhocOptions := &adhoc.AnsibleAdhocOptions{
		Inventory:  hiveview.CONFIG.Settings.Ansible.Inventory,
		ModuleName: moduleName,
		Args:       args,
	}
	Execute := execute.NewDefaultExecute(
		execute.WithWrite(io.Writer(buff)),
	)
	Adhoc := &adhoc.AnsibleAdhocCmd{
		Pattern:           targetHost,
		Exec:              Execute,
		ConnectionOptions: ansibleAdhocConnectionOptions,
		Options:           ansibleAdhocOptions,
		StdoutCallback:    "json",
	}
	err := Adhoc.Run(context.TODO())
	if err != nil {
		LogPrint("err", err)
		return err
	}
	err = ioutil.WriteFile(Logfile, []byte(buff.String()), 0755)
	if err != nil {
		LogPrint("err", err)
		return err
	}
	return nil
}
