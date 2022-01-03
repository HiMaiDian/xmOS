package model

//CPU相关采集项
type Stat struct {
	Name         string `json:"name"`          //cpu name
	User         int64  `json:"user"`          //用户态时间
	Nice         int64  `json:"nice"`          //用户态时间(低优先级，nice>0)
	System       int64  `json:"system"`        //内核态时间
	Idle         int64  `json:"idle"`          //空闲时间
	IoWait       int64  `json:"iowait"`        //I/O等待时间
	Irq          int64  `json:"irq"`           //硬中断
	SoftIrq      string `json:"softirq"`       //软中断
	Intr         string `json:"intr"`          //系统启动以来的所有interrupts的次数情况
	Ctxt         int64  `json:"ctxt"`          //系统上下文切换次数
	Btime        int64  `json:"btime"`         //启动时长(单位:秒)，从Epoch(即1970零时)开始到系统启动所经过的时长，每次启动会改变。
	Processes    int64  `json:"processes"`     //系统启动后所创建过的进程数量。当短时间该值特别大，系统可能出现异常
	ProcsRunning int64  `json:"procs_running"` //处于Runnable状态的进程个数
	ProcsBlocked int64  `json:"procs_blocked"` //处于Runnable状态的进程个数
	Cnt          int    `json:"cnt"`           //核心数量
}

//开机状态
type Uptime struct {
	RunningTime float32 `json:"running_time"` //开机累计时间
	FreeTime    float32 `json:"free_time"`    //开机CPU空间时间
}

//磁盘相关信息
type Disk struct {
}

//megacli工具输出
type Megacli struct {
}

//SMART工具输出
type Smart struct {
}

//分区读写监控
type Partition struct {
}

//IO相关采集项
type Io struct {
}

//机器负载相关采集项
type Machine struct {
}

//内存相关采集项
type Memory struct {
}

//网络相关采集项
type Network struct {
}

//端口采集项
type Port struct {
}

//机器内核配置
type Kernel struct {
}

//ntp采集项
type Ntp struct {
}

//进程监控
type Proccess struct {
}

//进程资源监控
type ProcResource struct {
}

//ss命令输出
type SS struct {
}
