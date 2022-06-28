package str

import (
	"regexp"
	"strconv"
)

func StrToInt64(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return num
	}
	
	return num
}

func InterfaceToString(data interface{}) (string, error) {
	d, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	
	return string(d), nil
}

//检测字符串是否包含汉字,包含为true，不包含为false
func CheckString(str string) bool {
	reg := regexp.MustCompile(`[\p{Han}]+`)
	result := reg.FindAllString(str, -1)
	if len(result) >= 1 {
		return true
	}
	return false
}

//检测字符串是否包含阿拉伯数字,包含为true，不包含为false
func CheckInt(str string) bool {
	reg := regexp.MustCompile(`[0-9]+`)
	result := reg.FindAllString(str, -1)
	if len(result) != 0 {
		return true
	}
	return false
}

