//API 参数检查，O开头表示非必传参数，X开头表示必传参数
package base

import (
	"fmt"
	"mime/multipart"
	"reflect"
	"strconv"
	"strings"
)

const (
	ARGS_OK             = 0
	ARGS_EMPTY          = 1
	ARGS_NOT_FOUND      = 2
	ARGS_INVALID_FORMAT = 3
)

type ApiArgs struct {
	Name  string
	raw   map[string]string
	ReqId string
	File  *multipart.FileHeader
}

func NewArgs(apiName string) *ApiArgs {
	return &ApiArgs{
		Name: apiName,
		raw:  make(map[string]string),
	}
}

func (args *ApiArgs) SetRaw(key, val string) {
	args.raw[key] = val
}

func (args *ApiArgs) GetRAW(key string) (string, int) {
	v, ok := args.raw[key]
	if ok {
		return v, ARGS_OK
	}
	return v, ARGS_NOT_FOUND
}

func (args *ApiArgs) ExistsKey(key string) bool {
	_, ok := args.raw[key]
	return ok
}

func ParseArgsStrToUint(key, v string) uint64 {
	num, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		Err(PARAMS_ERROR, key)
	}
	return num
}

//isCheck 为true表示检查required选项
func (args *ApiArgs) InputRequest(ptr interface{}, isCheck bool) error {
	v := reflect.ValueOf(ptr).Elem()
	for i, l := 0, v.NumField(); i < l; i++ {
		fieldTag := v.Type().Field(i).Tag
		fieldValue := v.Field(i)
		if !fieldValue.CanSet() {
			continue
		}
		key := fieldTag.Get("key")
		defaultValue := fieldTag.Get("default")
		required := fieldTag.Get("required")
		value, ok := args.raw[key]
		if !ok {
			if isCheck && required != "" {
				return fmt.Errorf("Missing PARAMS %s", key)
			}
			value = defaultValue
		}
		if len(value) == 0 {
			value = defaultValue
		}
		value = strings.TrimSpace(value)
		if len(value) == 0 {
			continue
		}
		if fieldValue.Kind() == reflect.Slice {
			vv := strings.Split(value, ",")
			for _, v := range vv {
				elem := reflect.New(fieldValue.Type().Elem()).Elem()
				if err := populate(elem, v); err != nil {
					return fmt.Errorf("%s format err", key)
				}
				fieldValue.Set(reflect.Append(fieldValue, elem))
			}
		} else {
			if err := populate(fieldValue, value); err != nil {
				return fmt.Errorf("%s format err", key)
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)
	case reflect.Uint32:
		i, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return err
		}
		v.SetUint(i)
	case reflect.Uint64:
		i, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetUint(i)
	case reflect.Int64:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)
	case reflect.Int32:
		i, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return err
		}
		v.SetInt(i)
	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)
	case reflect.Float64:
		i, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		v.SetFloat(i)
	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)
	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}
