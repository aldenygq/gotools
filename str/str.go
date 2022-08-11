package str

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

//string转int64
func StrToInt64(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return num
	}
	
	return num
}

//interface转string
func InterfaceToString(data interface{}) (string, error) {
	d, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	
	return string(d), nil
}

//[]byte转string
func ByteToStr(b  []byte) string {
	return string(b)
}

//string转[]byte
func StrToByte(s string) []byte {
	return []byte(s)
}

//string转[]rune
func StrToRune(s string) []rune {
	return []rune(s)
}

//[]rune转string
func RuneToStr(r []rune) string {
	return string(r)
}

//float32转string
func Float32ToStr(num float32) string {
	return strconv.FormatFloat(float64(num), 'E', -1, 32)
}

//float64转string
func Float64ToStr(num float64) string {
	return strconv.FormatFloat(num, 'E', -1, 64)
}
//string转float64
func StrToFloat64(str string) float64 {
	s,err := strconv.ParseFloat(str, 64)
	if err != nil {
		return s
	}
	return s
}
//string转float32
func StrToFloat32(str string) float32 {
	s,err := strconv.ParseFloat(str, 32)
	if err != nil {
		return float32(s)
	}
	return float32(s)
}

//string转uint64
func StrToUnit64(str string) uint64 {
	value,err := strconv.ParseUint(str,10,64)
	if err != nil {
		return value
	}
	
	return value
}
//uint64转string
func Uint64ToString(num uint64) string {
	return  strconv.FormatUint(num,10)
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

//是否为连续自增整数
func IsContinuousIncrementNum(str string) bool {
	slc := strings.Split(str, "")
	if len(slc) <0 {
		return true
	}
	
	for i := 0; i < len(slc)-1; i++ {
		n, _ := strconv.Atoi(slc[i])
		new := n + 1
		m, _ := strconv.Atoi(slc[i+1])
		if new != m {
			return false
		}
	}
	return true
}
//是否为均等整数
func IsEqualNum(str string) bool {
	slc:= strings.Split(str, "")
	if len(slc) <0 {
		return true
	}
	for i := 0; i < len(slc)-1; i++ {
		s, _ := strconv.Atoi(slc[i])
		e, _ := strconv.Atoi(slc[i+1])
		if s != e {
			return false
		}
	}
	
	return true
}

//是否为连续递减整数
func IsDecreaseNum(str string) bool {
	slc := strings.Split(str, "")
	if len(slc) <0 {
		return true
	}
	for i := 0; i < len(slc)-1; i++ {
		n, _ := strconv.Atoi(slc[i])
		new := n - 1
		m, _ := strconv.Atoi(slc[i+1])
		if new != m {
			return false
		}
	}
	
	return true
}

//是否包含英文字符
func HasEnglishChar(str string) bool {
	dictionary := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	for _, v := range str {
		if strings.Contains(dictionary, string(v)) {
			return true
		}
	}
	return false
}

//是否包含中文字符
func HasChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}

//string转bool
func StrToBool(str string) bool {
	//todo :string to bool
	//接受 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False 等字符串；
	//其他形式的字符串会返回错误
	b, _ := strconv.ParseBool(str)
	return b
}

//bool转string
func BoolToStr(b bool) string {
	//todo :bool to string
	sBool := strconv.FormatBool(true) //方法1
	return sBool
}

