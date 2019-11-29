/**
 * @Author: DollarKillerX
 * @Description: app.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午9:43 2019/11/29
 */
package service

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/dollarkillerx/easyutils/clog"
	"github.com/dollarkillerx/publicDns/datasource"
)

var dnsMu sync.Mutex
var DnsList []*datasource.DnsDataList

func UpdatePublicDnsListService() bool {
	infoNew := datasource.PublicDnsInfoNew()
	lists, e := infoNew.GetDataList()
	if e != nil {
		clog.PrintWa(e)
		return false
	}
	dnsMu.Lock()
	DnsList = lists
	dnsMu.Unlock()

	return true
}

func GetPublicDnsListService() ([]*datasource.DnsDataList, error) {
	dnsMu.Lock()
	lists := DnsList
	dnsMu.Unlock()
	if lists == nil {
		if err := UpdatePublicDnsListService(); !err {
			return nil, errors.New("dns error")
		}
	}
	dnsMu.Lock()
	defer dnsMu.Unlock()
	return DnsList, nil
}

func UpdateRegularly() {
	log.Println("定时任务 以启动")
	for {
		select {
		case <-time.After(time.Hour * 24):
			UpdatePublicDnsListService()
		}
	}
}
