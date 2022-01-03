package cron

import (
	"sync"
	"time"

	"github.com/MaiDian/XiaoMaiOS/pkg/c2s"
	"github.com/MaiDian/XiaoMaiOS/pkg/utils"
	"github.com/MaiDian/XiaoMaiOS/service"
)

//采集系统相关信息任务
type Cron struct {
	second int
	wg     sync.WaitGroup
}

func New(sc int) *Cron {
	return &Cron{second: sc}
}

func (c *Cron) Start() {
	go func() {
		time.Sleep(time.Second * time.Duration(c.second))

		c.wg.Add(1)
		go c.stat()
		c.wg.Wait()
	}()
}

func (c *Cron) stat() {
	defer c.wg.Done()
	stat := service.SysInfo.GetProcStat()

	service.Cache.SetTo2Run(c2s.KEY_STAT, stat, 10, func(i interface{}) {
		utils.Push(i)
	})
}
