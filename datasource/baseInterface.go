/**
 * @Author: DollarKillerX
 * @Description: baseInterface.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午10:17 2019/11/29
 */
package datasource

// 未来可能多数据源  我们这里为了应对这个问题 采取了插件式的设计思路

type DnsDataSource interface {
	GetDataList() ([]DnsDataList, error)
}

type DnsDataList struct {
	Ip      string
	Country string // 国家代码
}
