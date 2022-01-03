package cache

import (
	"math/rand"
	"sync"
	"time"
)

var instanse *Cache

type Cache struct {
	m  map[string]item
	mu sync.Mutex
}

type item struct {
	value      interface{}                                   //keyvalue
	timeout    int64                                         //超时时间
	timeoutRun func(value interface{})                       //超时后执行
	editRun    func(value interface{}, oldValue interface{}) //编辑后执行
	oTime      int64                                         //操作时间
}

func GetInstanse() *Cache {
	if instanse == nil {
		c := &Cache{}
		c.m = make(map[string]item)

		go c.cron()

		instanse = c
	}

	return instanse
}

func (c *Cache) IsExist(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.m[key]
	return ok
}

func (c *Cache) set(key string, value interface{}, to int64, to_run func(value interface{}), edit_run func(value interface{}, oldValue interface{})) bool {
	oldValue, b := c.Get(key)
	to = to + int64(rand.Intn(50))
	i := item{
		value:      value,
		timeout:    to,
		timeoutRun: to_run,
		editRun:    edit_run,
		oTime:      time.Now().Unix(),
	}
	c.mu.Lock()
	c.m[key] = i
	c.mu.Unlock()

	if b && value != oldValue {
		go edit_run(value, oldValue)
	}
	return true
}

func (c *Cache) Set(key string, value interface{}) error {
	c.set(key, value, 60, func(interface{}) {}, func(interface{}, interface{}) {})
	return nil
}

func (c *Cache) SetTo(key string, value interface{}, to int64) error {
	c.set(key, value, to, func(interface{}) {}, func(interface{}, interface{}) {})
	return nil
}

func (c *Cache) SetTo2Run(key string, value interface{}, to int64, to_run func(interface{})) error {
	c.set(key, value, to, to_run, func(interface{}, interface{}) {})
	return nil
}

func (c *Cache) SetTo2Edit(key string, value interface{}, to int64, edit_run func(interface{}, interface{})) error {
	c.set(key, value, to, func(value interface{}) {}, edit_run)
	return nil
}

func (c *Cache) SetTo2Run2Edit(key string, value interface{}, to int64, to_run func(interface{}), edit_run func(interface{}, interface{})) error {
	c.set(key, value, to, to_run, edit_run)
	return nil
}

func (c *Cache) Get(key string) (interface{}, bool) {
	if !c.IsExist(key) {
		return nil, false
	}

	c.mu.Lock()
	item := c.m[key]
	c.mu.Unlock()

	if time.Now().Unix()-item.oTime >= item.timeout {
		c.Remove(key)
		return nil, false
	}
	return item.value, true
}

func (c *Cache) Remove(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.m, key)
	return true
}

//定时任务
func (c *Cache) cron() {
	for {
		c.mu.Lock()
		if len(c.m) > 0 {
			for key, item := range c.m {
				if int64(time.Now().Unix()-item.oTime) >= item.timeout {
					go item.timeoutRun(item)
					c.Remove(key)
				}
			}
		}
		c.mu.Unlock()
		time.Sleep(time.Microsecond)
	}
}
