package main

import (
	"fmt"

	"github.com/MaiDian/XiaoMaiOS/pkg/cache"
	"github.com/MaiDian/XiaoMaiOS/pkg/config"
	"github.com/MaiDian/XiaoMaiOS/pkg/sys_info"
	"github.com/MaiDian/XiaoMaiOS/service"
)

func init() {
	service.Cache = cache.GetInstanse()
	service.SysInfo = sys_info.New()
}

func main() {
	// if runtime.GOOS == "windows" {
	// 	fmt.Println("不支持windows系统")
	// 	return
	// }

	//fmt.Printf("%#v\n", SysInfo.GetProcStat())
	// fmt.Println(sys_info.ExeCmdByShell("df -lh"))
	// fmt.Println(time.Now().Unix())
	// start := time.Now().Unix()
	// time.Sleep(time.Second)
	// fmt.Println(time.Now().Unix())

	// fmt.Println(time.Now().Unix() - start)
	// Cache.Set("name", "davie")
	// Cache.Set("age", 13)

	// fmt.Println(Cache.Get("name"))
	// fmt.Println(Cache.Get("age"))

	// Cache.Remove("name")
	// fmt.Println(Cache.Get("name"))
	// fmt.Println(Cache.Get("age"))

	// start := time.Now()
	// for i := 0; i < 10; i++ {
	// 	service.Cache.SetTo2Run(strconv.Itoa(i), i+1, 5, func(i interface{}) {
	// 		fmt.Println("timeout")
	// 	})
	// }
	// service.Cache.SetTo2Edit("3", 10, 5, func(i1, i2 interface{}) {
	// 	fmt.Println(i1, i2)
	// })
	// service.Cache.Set("3", 30)
	// time.Sleep(time.Second * 10)
	// fmt.Println(time.Now().Sub(start))
	// run()

	// for {

	// }
	// fmt.Println(conf.Server)
	config.Init()
	fmt.Println(config.Server)
	fmt.Println(config.Client)
}

// func run() {
// 	cron.New(5).Start()
// }
