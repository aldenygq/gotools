package slice

import "strings"

//校验切片中是否包含字符串
func CheckSlcHasStr(a string, b []string) bool {
	for _, v := range b {
		if strings.Contains(v, a) {
			return true
		}
	}
	
	return false
}

//去除数组中的空元素
func RemovEmpty(str []string) []string {
	var result []string = make([]string, 0)
	for i, _ := range str {
		if str[i] != "" {
			result = append(result, str[i])
		} else {
			continue
		}
	}
	
	return result
}
