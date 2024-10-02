package examples_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/oswaldoooo/sysutils"
	"github.com/oswaldoooo/sysutils/utils"
)

func GetStat() {
	t := time.NewTicker(time.Second * 5)
	format := "%-10.2f%-10.2f%-10.2f"
	fmt.Printf("%-10s%-10s%-10s"+"\n", "Cpu", "IO", "Memory")
	for range t.C {
		mem := sysutils.Memory()
		fmt.Printf(format+"\n", utils.Pressure2Float(sysutils.CpuPressure()), utils.Pressure2Float(sysutils.IOPressure()), float64(mem.Available)/float64(mem.Total))
	}
}
func TestGetStat(t *testing.T) {
	GetStat()
}
