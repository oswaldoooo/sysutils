package stat

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	_ "unsafe"

	"github.com/oswaldoooo/sysutils/types"
)

func LoadAvg() types.LoadAvgInfo {
	content, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return types.LoadAvgInfo{}
	}
	next := bytes.IndexByte(content, ' ')
	content = content[:next]
	content = bytes.Replace(content, []byte("."), []byte{}, 1)
	avg, _ := strconv.ParseUint(string(content), 10, 64)
	return types.LoadAvgInfo{Avg: avg}
}

func CpuPressure() uint64 {
	content, err := os.ReadFile("/proc/pressure/cpu")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 0
	}
	return parsePressure(content)
}
func parsePressure(content []byte) uint64 {
	next := bytes.IndexByte(content, ' ')
	content = content[next+1:]
	next = bytes.IndexByte(content, ' ')
	content = content[:next]
	next = bytes.IndexByte(content, '=')
	content = content[next+1:]
	content = bytes.Replace(content, []byte{'.'}, []byte{}, 1)
	val, _ := strconv.ParseUint(string(content), 10, 64)
	return val
}
func PortCheck(netype ...string) (ans []types.PortInfo) {
	for i := range netype {
		raw, err := os.ReadFile("/proc/net/" + netype[i])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		content := string(raw)
		contentLines := strings.Split(content, "\n")
		for e := range contentLines[1:] {
			if len(contentLines[e]) == 0 {
				continue
			}
			line := contentLines[e][:81]
			fmt.Printf("[%s] [%s] [%s]\n", line[6:19], line[6:14], line[15:19])
		}
	}
	return ans
}
func MemoryInfo() (ans types.MemoryInfo) {
	content, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	index := bytes.IndexRune(content, '\n')
	line1 := content[:index]
	content = content[index+1:]
	index = bytes.IndexRune(content, '\n')
	// line2 := content[:index]
	content = content[index+1:]
	index = bytes.IndexRune(content, '\n')
	line3 := content[:index]
	// content = content[index+1:]
	exactLine := func(content []byte) uint64 {
		contentLen := len(content)
		var (
			status    uint8
			lastIndex int
		)
		for i := contentLen - 1; i >= 0; i-- {
			if status == 0 && content[i] >= '0' && content[i] <= '9' {
				status = 1
				lastIndex = i
			} else if status == 1 && !(content[i] >= '0' && content[i] <= '9') {
				content = content[i+1 : lastIndex+1]
				val, err := strconv.ParseUint(string(content), 10, 64)
				if err != nil {
					fmt.Fprintln(os.Stderr, "parse meminfo failed ", err)
					return 0
				}
				return val
			}
		}
		return 0
	}
	ans.Total = exactLine(line1)
	ans.Available = exactLine(line3)
	return
}

func IOPressure() uint64 {
	content, err := os.ReadFile("/proc/pressure/io")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 0
	}
	return parsePressure(content)
}
