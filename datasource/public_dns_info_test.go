/**
 * @Author: DollarKillerX
 * @Description: public_dns_info_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午10:00 2019/11/29
 */
package datasource

import (
	"log"
	"testing"
)

func TestPublicDnsInfoNew(t *testing.T) {
	infoNew := PublicDnsInfoNew()
	lists, e := infoNew.GetDataList()
	if e != nil {
		panic(e)
	}
	for _, v := range lists {
		log.Println(v.Ip)
	}
	log.Println(len(lists))
}
