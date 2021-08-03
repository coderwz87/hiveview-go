
edit config/settings.yaml

go build

mkdir config dir and touch settings.yaml like :




settings:

  database:
  
    host:
      0.0.0.0
    port:
      3306
    user:
      test
    password:
      test
    dbname:
      test

  application:
  
    host:
      0.0.0.0
    port:
      9090

  ansible:
  
    factsdir:
      /tmp/fact
    inventory:
      /root/getAnsibleHost.py
  log:
  
    logpath:
      D:\goproject\src\hiveview\log\
    logname:
      hiveview.log
    accesslogpath:
      D:\goproject\src\hiveview\log\
    accesslogname:
      hiveview_access.log