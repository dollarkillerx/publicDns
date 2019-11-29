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
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

//type sliTest
//
//func TestSlice(t *testing.T) {
//	sli := make([]string, 0)
//	sli1(&sli)
//	log.Println(sli)
//}
//func sli1(sli *[]string) {
//	sli = append(sli, "asdas")
//}

func TestSLic(t *testing.T) {
	c := []string{
		"sdasd",
		"dsadas",
	}
	marshal, e := json.Marshal(c)
	if e != nil {
		panic(e)
	}
	ioutil.WriteFile("ss.dic", []byte(marshal), 00666)
}
