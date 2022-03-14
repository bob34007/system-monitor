/**
 * @Author: guobob
 * @Description:
 * @File:  cpuinfo.go
 * @Version: 1.0.0
 * @Date: 2022/3/10 21:05
 */

package cpu

import (
	"encoding/json"
	//"runtime"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/system-monitor/utils"
)


type CPUS struct {
	*CPUInfo
	*CPUStats
}

type CPUInfo struct {
	Cores int
	Numa  int
	Mhz   float64
}

type CPUStats struct {
	//load status
	Last1Load  float64
	Last5Load  float64
	Last15Load float64
	//process status
	ProcessStatusCount map[string]int
	//
}

func NewCPUS() *CPUS{
	return &CPUS{
		NewCPUInfo(),
		NewCPUStats(),
	}
}

func NewCPUStats()*CPUStats{
	return &CPUStats{
		ProcessStatusCount: make(map[string]int),
	}
}
func NewCPUInfo() *CPUInfo {
	return &CPUInfo{}
}

func (ci *CPUInfo) SetInfo() error {
	info, err := cpu.Info()
	if err != nil {
		return err
	}

	ci.Cores = len(info)
	ci.Mhz = info[0].Mhz

	return err
}

func (cs *CPUStats) SetLoadAvg() error {
	loadInfo, err := load.Avg()
	if err != nil {
		return err
	}
	cs.Last1Load = loadInfo.Load1
	cs.Last5Load = loadInfo.Load5
	cs.Last15Load = loadInfo.Load15
	return nil
}

func (cs *CPUStats) SetProcessStatusCount()error{
	return utils.GetProcessStatusCount(cs.ProcessStatusCount)
}

func (c *CPUS)Print()string{
	strInfo ,_ := json.Marshal(c.CPUInfo)
	strStat ,_ :=json.Marshal(c.CPUStats)
	return string(strInfo)+"\n"+string(strStat)
}