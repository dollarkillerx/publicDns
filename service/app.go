/**
 * @Author: DollarKillerX
 * @Description: app.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午9:43 2019/11/29
 */
package service

import (
	"errors"

	"github.com/dollarkillerx/erguotou/clog"
	"github.com/dollarkillerx/publicDns/datasource"
)

var DnsList []*datasource.DnsDataList

func UpdatePublicDnsListService() bool {
	infoNew := datasource.PublicDnsInfoNew()
	lists, e := infoNew.GetDataList()
	if e != nil {
		clog.PrintWa(e)
		return false
	}
	DnsList = lists
	return true
}

func GetPublicDnsListService() ([]*datasource.DnsDataList, error) {
	if DnsList == nil {
		if err := UpdatePublicDnsListService(); !err {
			return nil, errors.New("dns error")
		}
	}
	return DnsList, nil
}
