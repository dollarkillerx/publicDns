# publicDns
获取可用全球公共dns，可用于域名爆破数据源

### 依赖
- Erguotou Web Framework github.com/dollarkillerx/erguotou
- easyutils github.com/dollarkillerx/easyutils
- gocsv github.com/gocarina/gocsv

### 如何使用
``` 
./publicDns 0.0.0.0:8080
```
router
``` 
/update        # 更新dns列表
/getdnslist    # 获取dns list
```

### 更新日志
- init project (完善基本功能)
- 加入 定时任务  (每三天更新一次dnsList)