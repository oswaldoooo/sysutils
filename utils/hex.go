package utils

import "strconv"

func Hex2IP(s string) (ans string) {
	for i := 0; i < len(s); i += 2 {
		var (
			char string
			val  uint8
		)
		if '0' <= s[i+1] && s[i+1] <= '9' {
			val = s[i+1]
		} else if s[i+1] >= 'a' && s[i+1] <= 'f' {
			val = s[i+1] - 'a' + 10
		} else if s[i+1] >= 'A' && s[i+1] <= 'F' {
			val = s[i+1] - 'A' + 10
		}
		if '0' <= s[i] && s[i] <= '9' {
			val += s[i] * 16
		} else if s[i] >= 'a' && s[i] <= 'f' {
			val += (s[i] - 'a' + 10) * 16
		} else if s[i] >= 'A' && s[i] <= 'F' {
			val += (s[i] - 'A' + 10) * 16
		}
		char = strconv.FormatUint(uint64(val), 10)
		ans += char + "."
	}
	ans = ans[:len(ans)-1]
	return
}
func Bytes2Hex(src []byte) (result string) {
	for _, b := range src {
		result += strconv.FormatUint(uint64(b/16), 16) + strconv.FormatUint(uint64(b%16), 16)
	}
	return
}
