package base

import (
	"crypto/rand"
	"fmt"
	"webgo/pkg/errors"
	"reflect"
	"strconv"
	"strings"
)

// 用于判断一个interface是否为空
func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

func ParseInterfaceToString(v interface{}) string {
	if IsNil(v) {
		return ""
	}
	switch v.(type) {
	case string:
		return v.(string)
	case float64:
		num := v.(float64)
		tmp := strings.Split(fmt.Sprintf("%.6f", v.(float64)), ".")
		n, err := strconv.Atoi(tmp[1])
		errors.Assert(err)
		if n > 0 {
			return strconv.FormatFloat(num, 'f', 6, 64)
		}
		return strconv.FormatFloat(num, 'f', 0, 64)
	case bool:
		ok := v.(bool)
		if ok {
			return "true"
		}
		return "false"

	default:
		return ""
	}
}

func UUID(L int) string {
	const charMap = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	buf := make([]byte, L)
	_, err := rand.Read(buf)
	errors.Assert(err)
	for i := 0; i < L; i++ {
		ch := buf[i]
		buf[i] = charMap[int(ch)%62]
	}
	return string(buf)
}
