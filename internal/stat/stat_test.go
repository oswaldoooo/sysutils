package stat_test

import (
	"testing"

	"github.com/oswaldoooo/sysutils/internal/stat"
)

func TestPressure(t *testing.T) {
	t.Log(stat.CpuPressure())
	t.Log(stat.IOPressure())
}
func TestLoadAvg(t *testing.T) {
	t.Log(stat.LoadAvg())
}
func TestPortCheck(t *testing.T) {
	t.Log(stat.PortCheck("tcp"))
}
func TestMem(t *testing.T) {
	t.Log(stat.MemoryInfo())
}
