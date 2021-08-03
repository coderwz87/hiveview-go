package utils

//ansible-playbook完成应用初始化
func ServiceInit(serviceName, port, version, name, host, memSize string) (err error) {
	var playbookName string
	data := make(map[string]string)
	data["hosts"] = host
	data["project_name"] = name
	switch serviceName {
	case "mysql":
		playbookName = "/etc/ansible/playbook/mysql_install.yaml"
		data["mysql_version"] = version
		data["project_port"] = port
		//CMD = fmt.Sprintf("ansible-playbook -i %s /etc/ansible/playbook/mysql_install.yaml --extra-vars 'hosts=%s' --extra-vars 'mysql_version=%s' --extra-vars 'project_name=%s' --extra-vars 'project_port=%s' ", hiveview.CONFIG.Settings.Ansible.Inventory, host, version, name, port)
	case "redis":
		playbookName = "/etc/ansible/playbook/redis_install.yaml"
		data["project_mem"] = memSize
		//CMD = fmt.Sprintf("ansible-playbook -i %s /etc/ansible/playbook/redis_install.yaml --extra-vars 'hosts=%s' --extra-vars 'project_name=%s' --extra-vars 'project_port=%s' --extra-vars 'project_mem=%s' ", hiveview.CONFIG.Settings.Ansible.Inventory, host, name, port, memSize)
	}
	err = AnsiblePlaybook(playbookName, data)
	//cmd := exec.Command("/bin/bash", "-c", CMD)
	//_, err = cmd.Output()
	return err
}
