/**
 * @Author: DollarKillerX
 * @Description: app.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午9:36 2019/11/29
 */
package controller

import (
	"github.com/dollarkillerx/erguotou"
	"github.com/dollarkillerx/publicDns/service"
)

// 更新内部维护的dns列表
func UpdateDnsList(ctx *erguotou.Context) {
	if bool := service.UpdatePublicDnsListService(); bool {
		ctx.Json(200, "ok")
		return
	}
	ctx.Json(500, "server error")
}

// 返回给用户dns list
func GetDnsList(ctx *erguotou.Context) {
	lists, e := service.GetPublicDnsListService()
	if e != nil {
		ctx.Json(500, "server error")
		return
	}
	ctx.Json(200, lists)
}

// 获取dns list 通过国家名称过滤
func GetDnsListFilterByCountry(ctx *erguotou.Context) {
	s, b := ctx.PathValueString("filter")
	if !b || s == "" || s == "fd" {
		dnsService := service.GetDevelopedCountryPublicDnsService()
		ctx.Json(200, dnsService)
		return
	}
	dnsService := service.GetFilterCountryPublicDnsService(s)
	ctx.Json(200, dnsService)

}
