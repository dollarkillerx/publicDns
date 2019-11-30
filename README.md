# publicDns
获取可用全球公共dns，可用于域名爆破数据源

### 依赖
- Erguotou Web Framework github.com/dollarkillerx/erguotou
- easyutils  github.com/dollarkillerx/easyutils
- csvtools   github.com/dollarkillerx/csvtools

自研全家桶 哈

### 如何使用
``` 
./publicDns 0.0.0.0:8080
```
router
``` 
/update              # 更新dns列表
/getdnslist          # 获取dns list
/getdnslist/:filter  # 按国家过滤 (fd返回一些发达国家dns,其他国家输入CN,US,JP)
```

### 更新日志
- init project (完善基本功能)
- 加入 定时任务  (每三天更新一次dnsList)

### 分支 与 发布
- master  带有web支持的
- support 作为调用的(没有web模块) 默认发布的是这个版本 `go get github.com/dollarkillerx/publicDns`

### 测试发现好多dns都不能用 本地还要做个dns的数据清洗