/**
 * @Author: DollarKillerX
 * @Description: csv解析
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午11:45 2019/11/29
 */
package test

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/dollarkillerx/publicDns/service"
	"io"
	"io/ioutil"
	"log"
	"testing"
)

func TestDecodeCsv(t *testing.T) {
	byt, e := ioutil.ReadFile("nameservers.csv")
	if e != nil {
		panic(e)
	}
	reader := csv.NewReader(bytes.NewReader(byt))
	for {
		record, e := reader.Read()
		if e == io.EOF {
			break
		}
		if e != nil {
			continue
		}
		fmt.Println(record)
	}
}

func TestDnsList(t *testing.T) {
	lists, e := service.GetPublicDnsListService()
	if e != nil {
		panic(e)
	}
	log.Println(lists)
}