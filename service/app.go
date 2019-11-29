/**
 * @Author: DollarKillerX
 * @Description: app.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午9:43 2019/11/29
 */
package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"sync"
	"time"

	"github.com/bogdanovich/dns_resolver"
	"github.com/dollarkillerx/erguotou/clog"
	"github.com/dollarkillerx/publicDns/datasource"
)

var dnsMu sync.Mutex
var DnsList []string

var dnsC sync.Mutex // 为了保证只能有一个协程来清洗dns

func UpdatePublicDnsListService() bool {
	infoNew := datasource.PublicDnsInfoNew()
	lists, e := infoNew.GetDataList()
	if e != nil {
		clog.PrintWa(e)
		return false
	}

	// 清洗dns
	dnsC.Lock()
	log.Println("进入清洗模块!!!")
	lis := cleanDnsListNew(lists).Run()
	log.Println("dns清洗完毕!!!")
	dnsC.Unlock()

	dnsMu.Lock()
	DnsList = lis
	dnsMu.Unlock()

	// 将可用dns写入文件
	bytes, e := json.Marshal(lis)
	if e != nil {
		return false
	}
	e = ioutil.WriteFile("dns.dic", bytes, 00666)
	if e != nil {
		return false
	}
	return true
}

func GetPublicDnsListService() ([]string, error) {
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

// 定时任务更新dnsList
func UpdateRegularly() {
	log.Println("定时任务 以启动")
	for {
		select {
		case <-time.After(time.Hour * 24):
			UpdatePublicDnsListService()
		}
	}
}

// dns数据清洗留下高可用的dns
type cleanDnsList struct {
	dnsList []*datasource.DnsDataList
}

func cleanDnsListNew(dnsList []*datasource.DnsDataList) *cleanDnsList {
	return &cleanDnsList{
		dnsList: dnsList,
	}
}

var okMu = sync.Mutex{}
var okDnsList = make([]string, 0)

// 调度器
func (c *cleanDnsList) Run() []string {
	dnsChan := make(chan string, len(c.dnsList))

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go c.initChan(dnsChan, wg)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go c.okDns(dnsChan, wg)
	}

	wg.Wait()
	return okDnsList
}

// 初始化chann
func (c *cleanDnsList) initChan(dnsChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, v := range c.dnsList {
		dnsChan <- v.Ip
	}
	close(dnsChan)
}

// 验证dns可用性
func (c *cleanDnsList) okDns(dnsChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
loop:
	for {
		select {
		case dns, ok := <-dnsChan:
			if ok {
				// 进行dns验证  尝试次数2
				for i := 0; i < 2; i++ {
					checkDns := c.checkDns(dns)
					if checkDns {
						//log.Println("成功: ", dns)
						okMu.Lock()
						okDnsList = append(okDnsList, dns)
						okMu.Unlock()
					} else {
						continue
					}
				}
			} else {
				break loop
			}
		}
	}
}

func (c *cleanDnsList) checkDns(dns string) bool {
	resolver := dns_resolver.New([]string{dns})
	ips, e := resolver.LookupHost("www.google.com")
	if e != nil || len(ips) == 0 {
		return false
	}
	return true
}
