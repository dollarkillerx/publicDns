/**
 * @Author: DollarKillerX
 * @Description: main.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午9:36 2019/11/29
 */
package main

import (
	"github.com/dollarkillerx/erguotou"
	"github.com/dollarkillerx/erguotou/clog"
	"github.com/dollarkillerx/publicDns/controller"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("请输入程序运行host 例如:  ./publicDns 0.0.0.0:8080")
	}

	app := erguotou.New()
	router(app)

	if err := app.Run(erguotou.SetHost(os.Args[1])); err != nil {
		clog.PrintWa(err)
		os.Exit(1)
	}
}

// 路由
func router(app *erguotou.Engine) {
	// 更新内部维护的dns列表
	app.Get("/update", controller.UpdateDnsList)
	// 返回给用户dns list
	app.Get("/getdnslist", controller.GetDnsList)
}
