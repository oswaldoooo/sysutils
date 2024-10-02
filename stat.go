package sysutils

import (
	_ "unsafe"

	"github.com/oswaldoooo/sysutils/internal/stat"
	"github.com/oswaldoooo/sysutils/types"
)

func LoadAvg() types.LoadAvgInfo {
	return stat.LoadAvg()
}
func CpuPressure() uint64 {
	return stat.CpuPressure()
}
func CheckPort() []types.PortInfo {
	return stat.PortCheck("tcp", "udp")
}
func Memory() types.MemoryInfo {
	return stat.MemoryInfo()
}
func IOPressure() uint64 {
	return stat.IOPressure()
}
