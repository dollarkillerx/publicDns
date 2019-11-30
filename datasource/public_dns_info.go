/**
 * @Author: DollarKillerX
 * @Description: https://public-dns.info/
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午9:45 2019/11/29
 */
package datasource

import (
	"errors"
	"strconv"

	"github.com/dollarkillerx/csvtools"
	"github.com/dollarkillerx/easyutils/httplib"
	"github.com/dollarkillerx/publicDns/utils"
)

type PublicDnsInfo struct {
	url string
}

func PublicDnsInfoNew() *PublicDnsInfo {
	return &PublicDnsInfo{
		url: "https://public-dns.info/nameservers.csv",
	}
}

type PublicDnsInfoDataSourceCsv struct {
	IP          string `csv:"ip"`
	Name        string `csv:"name"`
	CountryId   string `csv:"country_id"`
	City        string `csv:"city"`
	Version     string `csv:"version"`
	Error       string `csv:"error"`
	Dnssec      string `csv:"dnssec"`
	Reliability string `csv:"reliability"`
	CheckedAt   string `csv:"checked_at"`
	CreatedAt   string `csv:"created_at"`
}

// 获取数据源
func (p *PublicDnsInfo) GetDataList() ([]*DnsDataList, error) {
	csvs, e := p.dowDataCsv()
	if e != nil {
		return nil, e
	}
	// 将本数据源数据结构 转变为通用数据结构
	csv := []*DnsDataList{}
	for _, v := range csvs {
		// 对数据进行过滤 把 > 0.8 稳定性的服务器 写出
		f, e := strconv.ParseFloat(v.Reliability, 32)
		if e != nil {
			return nil, e
		}
		if !utils.NotIpv6(v.IP) {
			continue
		}
		if f > 0.8 {
			csv = append(csv, &DnsDataList{Ip: v.IP, Country: v.City})
		}
	}
	return csv, nil
}

// 下载最新的dns list 返回 解析好的csv和 错误信息
func (p *PublicDnsInfo) dowDataCsv() ([]*PublicDnsInfoDataSourceCsv, error) {
	// 网络原因导致下载出错  进行尝试 最多三次
	for i := 0; i < 3; i++ {
		bytes, e := httplib.EuUserGet(p.url)
		if e != nil {
			continue
		}
		// 如果没有问题 就 下载文件到本地 并重新命名
		csv, e := p.bind(bytes)
		if e != nil {
			return nil, e
		}
		return csv, nil
	}
	return nil, errors.New("dow error")
}

func (p *PublicDnsInfo) bind(byt []byte) ([]*PublicDnsInfoDataSourceCsv, error) {
	result := []*PublicDnsInfoDataSourceCsv{}
	readByte, e := csvtools.ReadByte(byt)
	if e != nil {
		return nil, e
	}
	decode := readByte.Decode()
	for _, v := range decode {
		item := PublicDnsInfoDataSourceCsv{
			v[0],
			v[1],
			v[2],
			v[3],
			v[4],
			v[5],
			v[6],
			v[7],
			v[8],
			v[9],
		}
		result = append(result, &item)
	}
	result = result[1:]
	return result, nil
}
