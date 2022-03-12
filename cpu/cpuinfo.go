/**
 * @Author: guobob
 * @Description:
 * @File:  cpuinfo.go
 * @Version: 1.0.0
 * @Date: 2022/3/10 21:05
 */

package cpu

import (
	//"runtime"
	"github.com/shirou/gopsutil/cpu"

)
type CPUInfo struct {
	CPUCores int
	CPUNuma  int
	CPUMhz float64
	CPULast1MinUse float32
	CPULast5MinUse float32
	CPULast15MinUse float32
}

func NewCPUInfo() *CPUInfo{
	return &CPUInfo{

	}
}



func (c *CPUInfo )SetInfo() error {
	info ,err := cpu.Info()
	if err != nil {
		return err
	}
	c.CPUCores = len(info)
	c.CPUMhz=info[0].Mhz


	return err
}