package utils_test

import (
	"strconv"
	"testing"

	"github.com/oswaldoooo/sysutils/utils"
)

func TestHex(t *testing.T) {
	t.Log(utils.Hex2IP("92C8A8C0"))
	val, _ := strconv.ParseUint("92", 16, 8)
	t.Log(val)
}
