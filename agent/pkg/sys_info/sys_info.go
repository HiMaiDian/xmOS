package sys_info

import (
	"strconv"
	"strings"

	"github.com/MaiDian/XiaoMaiOS/model"
	"github.com/MaiDian/XiaoMaiOS/pkg/c2s"
	"github.com/MaiDian/XiaoMaiOS/pkg/utils/command"
)

type SysInfo struct {
}

func New() *SysInfo {
	return &SysInfo{}
}

//CPU相关采集项
func (s *SysInfo) GetProcStat() model.Stat {
	r := command.Cat(c2s.CMD_PROC_STAT)
	var statSlice model.Stat
	var cnt = 0
	for _, line := range r {
		valueS := strings.Split(line, " ")
		if strings.Index(line, "cpu") == 0 {
			cnt++
			if strings.Index(line, "cpu ") == 0 {
				statSlice.User, _ = strconv.ParseInt(valueS[1], 10, 64)
				statSlice.Nice, _ = strconv.ParseInt(valueS[2], 10, 64)
				statSlice.System, _ = strconv.ParseInt(valueS[3], 10, 64)
				statSlice.Idle, _ = strconv.ParseInt(valueS[4], 10, 64)
				statSlice.IoWait, _ = strconv.ParseInt(valueS[5], 10, 64)
				statSlice.Irq, _ = strconv.ParseInt(valueS[6], 10, 64)
			}
			continue
		}
		if strings.Index(line, "intr") == 0 {
			statSlice.Intr = strings.Replace(line, "intr", "", 1)
			continue
		}
		if strings.Index(line, "ctxt") == 0 {
			statSlice.Ctxt, _ = strconv.ParseInt(valueS[1], 10, 64)
			continue
		}
		if strings.Index(line, "btime") == 0 {
			statSlice.Btime, _ = strconv.ParseInt(valueS[1], 10, 64)
			continue
		}
		if strings.Index(line, "processes") == 0 {
			statSlice.Processes, _ = strconv.ParseInt(valueS[1], 10, 64)
			continue
		}
		if strings.Index(line, "procs_running") == 0 {
			statSlice.ProcsRunning, _ = strconv.ParseInt(valueS[1], 10, 64)
			continue
		}
		if strings.Index(line, "procs_blocked") == 0 {
			statSlice.ProcsBlocked, _ = strconv.ParseInt(valueS[1], 10, 64)
			continue
		}
		if strings.Index(line, "softirq") == 0 {
			statSlice.SoftIrq = strings.Replace(line, "softirq", "", 1)
			continue
		}
	}

	statSlice.Cnt = cnt - 1

	return statSlice
}

//磁盘相关采集项
func (s *SysInfo) Disk() model.Disk {
	return model.Disk{}
}

//megacli工具输出
func (s *SysInfo) Megacli() model.Megacli {
	return model.Megacli{}
}

//SMART工具输出
func (s *SysInfo) Smart() model.Smart {
	return model.Smart{}
}

//分区读写监控
func (s *SysInfo) Partition() model.Partition {
	return model.Partition{}
}

//IO相关采集项
func (s *SysInfo) Io() model.Io {
	return model.Io{}
}

//机器负载相关采集项
func (s *SysInfo) Machine() model.Machine {
	return model.Machine{}
}

//内存相关采集项
func (s *SysInfo) Memory() model.Memory {
	return model.Memory{}
}

//网络相关采集项
func (s *SysInfo) Network() model.Network {
	return model.Network{}
}

//端口采集项
func (s *SysInfo) Port() model.Port {
	return model.Port{}
}

//机器内核配置
func (s *SysInfo) Kernel() model.Kernel {
	return model.Kernel{}
}

//ntp采集项
func (s *SysInfo) Ntp() model.Ntp {
	return model.Ntp{}
}

//进程监控
func (s *SysInfo) Processes() model.Proccess {
	return model.Proccess{}
}

//进程资源监控
func (s *SysInfo) ProcResource() model.ProcResource {
	return model.ProcResource{}
}

//ss命令输出
func (s *SysInfo) SS() model.SS {
	return model.SS{}
}
